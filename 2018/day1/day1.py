def parse_numbers(number: str):
    """Parses a number with a symbol"""
    mod = number[0]
    num = int(number[1:])

    if mod == "+":
        return num
    elif mod == "-":
        return 0 - num


def assignment1(input_file):
    i = 0
    with open(input_file) as f:
        for line in f:
            i += parse_numbers(line.strip())
    return (i)


if __name__ == '__main__':
    print(assignment1("input.txt"))