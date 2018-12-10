from day1 import parse_input_file


def parse_claim(*claim: iter) -> dict:
    """
    Retrieves start points and length of the claim

    Args:
        *claim (): iterable containing parsable claim information

    Returns:
        A dict containing claim information

    """
    parsed_inputs = {}

    for input_string in claim:
        identifier, _, skip, length = input_string.strip().split()
        identifier = int(identifier[1:])
        skip = tuple(map(int, skip[:-1].split(',')))
        length = tuple(map(int, length.split('x')))

        parsed_inputs[identifier] = {'skip': skip, 'length': length}

    return parsed_inputs


def find_start(skip: tuple) -> tuple:
    """Returns top-left starting point"""
    return tuple(x + 1 for x in skip)


def retrieve_coords(barcodes: dict):
    """Adds coordinates parsed from skip and length metadata to claim"""
    for barcode in barcodes.values():
        barcode['coords'] = []
        start = find_start(barcode['skip'])
        for i in range(start[0], start[0] + barcode['length'][0]):
            for j in range(start[1], start[1] + barcode['length'][1]):
                barcode['coords'].append((i, j))
        barcode['coords'] = sorted(barcode['coords'])

    return barcodes


def find_overlap(barcode1, barcode2):
    """Finds overlap between any two claims"""
    return set(barcode1['coords']) & set(barcode2['coords'])


def assignment1(parsed_dict):
    coords_dict = retrieve_coords(parsed_dict)

    overlaps = []
    ids = set()
    for i in range(1, len(coords_dict)):
        for j in range(i + 1, len(coords_dict) + 1):
            overlap = sorted(find_overlap(coords_dict[i], coords_dict[j]))
            if overlap:
                ids.add(i)
                ids.add(j)
            overlaps.extend(overlap)

    return set(overlaps), ids


def assignment2(barcode_ids, overlapping_ids):
    return barcode_ids - overlapping_ids


if __name__ == '__main__':
    barcode_dict = parse_claim(*parse_input_file("input.txt"))
    assignment_1 = assignment1(barcode_dict)
    assignment_2 = assignment2(set(barcode_dict.keys()), assignment_1[1])[0]
    print("Assignment 1: {}".format(len(assignment_1[0])))
    print("Assignment 2: {}".format(assignment_2))
