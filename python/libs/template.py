import sys
S = input()


def next_line():
    return input()


def next_line_value():
    return int(input())


def next_line_values():
    s = input()
    valstr = s.split(" ")
    return [int(v) for v in valstr]


def digit_total():
    pass


l1 = next_line_values()
print(l1)
l2 = next_line_values()
print(l2)
l3 = next_line_values()
print(l3)
matrix = list()
matrix.extend(l1, l2, l3)

n = next_line_value()
values = dict()
for i in range(n):
    values[next_line_value()] = True

print(matrix)
print(values)