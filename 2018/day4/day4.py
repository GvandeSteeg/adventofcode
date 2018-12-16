import re
from collections import defaultdict
from datetime import datetime, timedelta
from typing import *

from iso8601 import parse_date

from day1 import parse_input_file


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


def calculate_sleep(guard: List[Dict[datetime, AnyStr]]):
    asleep = []
    time_last_asleep = None
    time_most_asleep = None
    for iteration in guard:
        if iteration['metadata'].lower() == 'falls asleep':
            time_last_asleep = iteration['datetime']
        elif iteration['metadata'].lower() == 'wakes up':
            time_last_awake = iteration['datetime'] - timedelta(minutes=1)
            sleep = ((time_last_awake - time_last_asleep).seconds // 60)
            asleep.append(sleep)
            if sleep >= max(asleep):
                time_most_asleep = time_last_asleep

    return sum(asleep), time_most_asleep


if __name__ == '__main__':
    inputs = sorted(parse_input_file("input.txt"))
    guard_list = dict(parse_guard_metadata(*inputs))
    for guard, guard_data in guard_list.items():
        print(guard, calculate_sleep(guard_data))
