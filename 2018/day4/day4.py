import re
from collections import defaultdict
from datetime import datetime
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
            current_guard = int(metadata.split()[1][1:])
        else:
            guards[current_guard].append(dict(datetime=datetime, metadata=metadata))
    return guards


def calculate_sleep(guard: List[Dict[datetime, AnyStr]]):
    asleep = 0
    time_last_asleep = datetime(1, 1, 1)
    minute_count = defaultdict(int)
    for iteration in guard:
        if iteration['metadata'].lower() == 'falls asleep':
            time_last_asleep = iteration['datetime']
        elif iteration['metadata'].lower() == 'wakes up':
            time_last_awake = iteration['datetime']
            asleep += ((time_last_awake - time_last_asleep).seconds // 60)
            for minute in range(time_last_asleep.minute, time_last_awake.minute):
                minute_count[minute] += 1

    return asleep, [key for key, value in minute_count.items() if value == max(minute_count.values())][0]


if __name__ == '__main__':
    inputs = sorted(parse_input_file("input.txt"))
    guard_list = dict(parse_guard_metadata(*inputs))
    sleepy_guard, sleep_minute, sleep_time = 0, 0, 0
    for guard, guard_data in guard_list.items():
        sleeptime, minute = calculate_sleep(guard_data)
        if sleeptime > sleep_time:
            sleep_time = sleeptime
            sleep_minute = minute
            sleepy_guard = guard

    print(sleepy_guard * sleep_minute)
