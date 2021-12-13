def read_file(filename):
    """
    Read input file and save the lines.
    :param filename: input file
    :return: list of read lines (start and end)
    """
    lines = []
    with open(filename, 'r', encoding='UTF-8') as file:
        for line in file:
            start, end = line.strip().split('->')
            start = start.strip()
            end = end.strip()
            lines.append([start, end])

    return lines


def number_range(a, b):
    """
    Generate consecutive values list between two numbers.
    """
    if (a == b):
        return [a]
    else:
        mx = max(a, b)
        mn = min(a, b)
        result = []
        # inclusive upper limit. If not needed, delete '+1' in the line below
        while (mn < mx + 1):
            # We go from min to max
            result.append(mn)
            mn += 1
        if a > b:
            # Reverse the list if the first number is bigger than the second
            result = reversed(result)
        return result


def mark_hv_vents(lines):
    """
    Mark horizontal and vertical vents.
    :param lines: lines of coordinates
    :return: dictionary of vents
    """
    vents = {}
    for line in lines:
        start, end = line
        x1, y1 = start.split(",")
        x2, y2 = end.split(",")

        x1 = int(x1)
        x2 = int(x2)
        y1 = int(y1)
        y2 = int(y2)

        if x1 == x2 or y1 == y2:
            for x in number_range(x1, x2):
                for y in number_range(y1, y2):
                    key = f'{x},{y}'
                    if key not in vents:
                        vents[key] = 1
                    else:
                        vents[key] += 1

    return vents


def count_overlaps(vents):
    """
    Count overlaping vents.
    :param vents: dictionary of vents
    :return: number of overlaping vents
    """
    counter = 0
    for n_vents in vents.values():
        if n_vents > 1:
            counter += 1

    return counter


def mark_vents(lines):
    """
    Mark all vents.
    :param lines: lines of coordinates
    :return: dictionary of vents
    """
    vents = {}
    for line in lines:
        start, end = line
        x1, y1 = start.split(",")
        x2, y2 = end.split(",")

        x1 = int(x1)
        x2 = int(x2)
        y1 = int(y1)
        y2 = int(y2)

        if x1 == x2 or y1 == y2:
            for x in number_range(x1, x2):
                for y in number_range(y1, y2):
                    key = f'{x},{y}'
                    if key not in vents:
                        vents[key] = 1
                    else:
                        vents[key] += 1
        elif abs(y1 - y2) == abs(x1 - x2):
            for x, y in zip(number_range(x1, x2), number_range(y1, y2)):
                key = f'{x},{y}'
                if key not in vents:
                    vents[key] = 1
                else:
                    vents[key] += 1

    return vents


if __name__ == '__main__':
    lines = read_file('input.txt')
    hv_vents = mark_hv_vents(lines)
    hv_counter = count_overlaps(hv_vents)
    vents = mark_vents(lines)
    counter = count_overlaps(vents)
    print(hv_counter)
    print(counter)
