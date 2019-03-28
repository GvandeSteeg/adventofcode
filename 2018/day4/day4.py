import re
from collections import defaultdict
from datetime import datetime
from typing import *

from iso8601 import parse_date

from day1 import parse_input_file


def parse_guard_metadata(*inputs: str) -> Dict[int, List[Dict[str, Union[datetime, str]]]]:
    """
    Parses guard metadata into a list of dicts

    [1518-11-01 00:00] Guard #10 begins shift
    [1518-11-01 00:05] falls asleep
    [1518-11-01 00:25] wakes up

    becomes {10: [
                    {'datetime': 1518-11-01 00:05, 'metadata': 'falls asleep'},
                    {'datetime': 1518-11-01 00:25, 'metadata': 'wakes up'}
                 ],
                 11: [...]
                 }

    :param inputs: Iterable of strings with guard metadata
    """
    guards = defaultdict(list)
    current_guard = ''
    for item in inputs:
        datetime, metadata = re.search(r"\[([\d -:]+)] (.+)", item).groups()
        datetime = parse_date(datetime)
        if "guard" in metadata.lower():
            current_guard = int(metadata.split()[1][1:])
        else:
            guards[current_guard].append(dict(datetime=datetime, metadata=metadata))
    return dict(guards)


def calculate_sleep(guard: List[Dict[datetime, str]]):
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


def assignment1(guards):
    sleepy_guard, sleep_minute, sleep_time = 0, 0, 0
    for guard, guard_data in guards.items():
        sleeptime, minute = calculate_sleep(guard_data)
        if sleeptime > sleep_time:
            sleep_time = sleeptime
            sleep_minute = minute
            sleepy_guard = guard

    return (sleepy_guard * sleep_minute)


def assignment2():
    pass


if __name__ == '__main__':
    inputs = sorted(parse_input_file("input.txt"))
    guard_list = parse_guard_metadata(*inputs)
    print("Assignment 1: {}".format(assignment1(guard_list)))
    print("Assignment 2: {}".format(assignment2()))
