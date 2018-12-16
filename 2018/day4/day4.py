import re
from collections import defaultdict

from day1 import parse_input_file
from iso8601 import parse_date


def parse_guard_metadata(*inputs):
    guards = defaultdict(list)
    current_guard = ''
    for item in inputs:
        datetime, metadata = re.search(r"\[([\d -:]+)] (.+)", item).groups()
        datetime = parse_date(datetime)
        if "guard" in metadata.lower():
            current_guard = metadata.split()[1]
        else:
            guards[current_guard].append(dict(datetime=datetime, metadata=metadata))
    return guards


if __name__ == '__main__':
    inputs = sorted(parse_input_file("input.txt"))
    print(dict(parse_guard_metadata(*inputs[:5])))
