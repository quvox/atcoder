from argparse import ArgumentParser
import os
import sys
import subprocess
import time


def _parser():
    usage = 'python {} [-d directory] [-t testcase file] [-a anser file] [--help]'.format(__file__)
    argparser = ArgumentParser(usage=usage)
    argparser.add_argument('-d', '--directory', type=str, help='root directory for the test case files')
    argparser.add_argument('-t', '--testcase', type=str, help='filename of testcase')
    argparser.add_argument('-a', '--answer', type=str, help='filename of test anser')
    args = argparser.parse_args()
    return args


def run_code(directory, testfile, answerfile):
    os.chdir(directory)
    cmd1 = "cat %s" % testfile
    cmd2 = "make run"
    try:
        start_time = int(time.time())
        res = subprocess.Popen(cmd1.split(" "), stdout=subprocess.PIPE)
        code_result = subprocess.check_output(cmd2.split(" "), stdin=res.stdout).decode().rstrip()
        process_time = int(time.time()) - start_time
    except subprocess.CalledProcessError as e:
        print("\n#### Exit with error: code=%d" % (e.returncode))
        return

    timeover = {True: " [TLE]", False: ""}
    print("** test case and answer: %s, %s" % (testfile, answerfile))
    print("  elapsed_time = %d ms%s" % (process_time, timeover[process_time >= 2000]))
    print("  [%s] code_result: %s" % (compare_result(answerfile, code_result), code_result))


def compare_result(filepath, result_string):
    with open(filepath, "r") as f:
        answer_string = f.read()
    return answer_string.rstrip() == result_string


if __name__ == '__main__':
    arg = _parser()
    if not os.path.exists(arg.directory):
        sys.stderr.write("# no such directory")
        sys.exit(1)
    if not os.path.exists(os.path.join(arg.directory, arg.testcase)):
        sys.stderr.write("# no such test file")
        sys.exit(1)
    if not os.path.exists(os.path.join(arg.directory, arg.answer)):
        sys.stderr.write("# no such test answer file")
        sys.exit(1)
    run_code(arg.directory, arg.testcase, arg.answer)

