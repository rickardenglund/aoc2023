input = """1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
"""

input_2 = """two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
"""


def main():
    print("p1", solve(input, read_row))
    print("p2", solve(input_2, read_row_2))


def solve(in_data, get_calibration_values):
    lines = in_data.split("\n")

    cvs = [
        int(get_calibration_values(x))
        for x in lines
        if x != ""
    ]

    return sum(cvs)


def read_row(l):
    res = ""
    for c in l:
        if c >= '0' and c <= '9':
            res += c

    assert len(res) > 0, "failed on line, " + l
    return res[0] + res[-1]


def read_row_2(l):
    res = ""
    for i in range(len(l)):
        c = l[i]
        if c >= '0' and c <= '9':
            res += c
        else:
            res += get_number(l[i:])

    return res[0] + res[-1]


def get_number(r):
    for k in numbers:
        if r.startswith(k):
            return numbers[k]
    return ""


numbers = {
    "one": "1",
    "two": "2",
    "three": "3",
    "four": "4",
    "five": "5",
    "six": "6",
    "seven": "7",
    "eight": "8",
    "nine": "9",
    "zero": "0"
}

main()
