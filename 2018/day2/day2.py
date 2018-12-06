from collections import Counter

from day1 import parse_input_file


def count_twos_and_threes(input_string):
    count = Counter(input_string).values()
    return any(i == 2 for i in count), any(i == 3 for i in count)


def assignment1(*input_list):
    twos = 0
    threes = 0
    for i in input_list:
        two, three = count_twos_and_threes(i)
        twos += 1 if two else 0
        threes += 1 if three else 0

    return twos * threes


def find_hamm1_strings(*input_list):
    for i in range(len(input_list)):
        for j in range(i, len(input_list)):
            hamm = 0
            for x, y in zip(input_list[i], input_list[j]):
                if x != y:
                    hamm += 1
            if hamm == 1:
                return input_list[i], input_list[j]


def common_letters(string_a, string_b):
    return ''.join([leta for leta, letb in zip(string_a, string_b) if leta == letb])


def assignment2(*input_list):
    return common_letters(*find_hamm1_strings(*input_list))


if __name__ == '__main__':
    box_IDs = parse_input_file("input.txt")
    print("Assignment 1: {}".format(assignment1(*box_IDs)))
    print("Assignment 2: {}".format(assignment2(*box_IDs)))
