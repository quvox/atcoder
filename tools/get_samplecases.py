from argparse import ArgumentParser
import sys
import os
import re
import requests
from bs4 import BeautifulSoup


BASE_URL = "https://atcoder.jp/contests/"


def _parser():
    usage = 'python {} [-u username] [-p password] [-c contest] [-q question] [-d directory] [--help]'.format(__file__)
    argparser = ArgumentParser(usage=usage)
    argparser.add_argument('-d', '--directory', type=str, default='samples', help='root directory for the test case files')
    argparser.add_argument('-u', '--user', type=str, help='username to login')
    argparser.add_argument('-p', '--password', type=str, help='password to login')
    argparser.add_argument('-c', '--contest', type=str, help='contest name (e.g. abc123)')
    argparser.add_argument('-q', '--question', type=str, help='question name (e.g. a)')
    args = argparser.parse_args()
    return args


def login(username, password):
    session = requests.Session()
    res = session.get("https://atcoder.jp/login")
    soup = BeautifulSoup(res.text, 'html.parser')
    csrf_input = soup.find("input")
    csrf_token = csrf_input["value"]
    resp = session.post("https://atcoder.jp/login",
                 data={'username': username, "password": password, "csrf_token": csrf_token})
    return session


if __name__ == '__main__':
    arg = _parser()
    if not os.path.exists(arg.directory):
        os.makedirs(arg.directory)
    os.makedirs(os.path.join(arg.directory, "in"), exist_ok=True)
    os.makedirs(os.path.join(arg.directory, "out"), exist_ok=True)

    url = BASE_URL+arg.contest+"/tasks/"+arg.question
    session = login(arg.user, arg.password)

    html = session.get(url).text
    soup = BeautifulSoup(html, 'html.parser')
    samples = soup.find_all("div", class_="part")
    if len(samples) == 0:
        print("XXXXX no such contest or question!!", arg.contest, arg.question)
        sys.exit(1)
    for x in samples:
        title_text = x.find("h3").text
        if title_text.find("Sample ") == -1: continue
        m = re.match(r"Sample (.*?) (\d+)", title_text)
        if m is None: continue
        tp, num = m.groups()
        filename = "sample_%s" % num
        content = x.find("pre").text
        if tp == "Input":
            path = os.path.join(arg.directory, "in", filename)
        else:
            path = os.path.join(arg.directory, "out", filename)
        with open(path, "w") as f:
            f.write(content)
