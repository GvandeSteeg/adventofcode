def parse_input_file(input_file):
    with open(input_file) as f:
        return f.read().split('\n')


def parse_numbers(number: str):
    """Parses a number with a symbol"""
    mod = number[0]
    num = int(number[1:])

    if mod == "+":
        return num
    elif mod == "-":
        return 0 - num


def assignment1(num_list):
    i = 0
    for line in num_list:
        i += parse_numbers(line.strip())
    return i


def assignment2(num_list):
    i = 0
    test_set = set()
    for line in num_list:
        i += parse_numbers(line.strip())
        if i not in test_set:
            test_set.add(i)
        else:
            return i
    return 0


if __name__ == '__main__':
    numbers = parse_input_file("input.txt")
    print("Assignment 1: {}".format(assignment1(numbers)))
    print("Assignment 2: {}".format(assignment2(numbers)))
