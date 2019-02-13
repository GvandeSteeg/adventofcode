def parse_input_file(input_file):
    with open(input_file) as f:
        return [line.strip() for line in f]


def parse_numbers(number: str):
    """Parses a number with a symbol"""
    return int(number)


def assignment1(num_list):
    i = 0
    for line in num_list:
        i += parse_numbers(line.strip())
    return i


def assignment2(num_list: list):
    test_set = set()
    frequency = 0

    while True:
        for line in num_list:
            frequency += parse_numbers(line.strip())
            if frequency not in test_set:
                test_set.add(frequency)
            else:
                return frequency


if __name__ == '__main__':
    numbers = parse_input_file("input.txt")
    print("Assignment 1: {}".format(assignment1(numbers)))
    print("Assignment 2: {}".format(assignment2(numbers)))
