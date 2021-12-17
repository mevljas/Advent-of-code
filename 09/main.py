from functools import reduce


def read_file(filename):
    """
    Read input file and save the locations into a list.
    :param filename: input file
    :return: list of locations
    """
    locations = []
    with open(filename, 'r', encoding='UTF-8') as file:
        for line in file:
            locations.append([[int(location), False, False] for location in line.strip()])
    return locations


def mark_low_points(points):
    """
    Mark lowest points in their area.
    :param points: list of points
    :return: marked list of points
    """
    for i in range(0, len(points)):
        for j in range(0, len(points[i])):
            point, _, _ = points[i][j]
            # Check if there are lower point around
            # up
            if i != 0 and points[i - 1][j][0] <= point:
                continue
            # down
            if i != len(points) - 1 and points[i + 1][j][0] <= point:
                continue
            # left
            if j != 0 and points[i][j - 1][0] <= point:
                continue
            # right
            if j != len(points[i]) - 1 and points[i][j + 1][0] <= point:
                continue
            # this is the lowest point
            points[i][j][1] = True
    return points


def sum_risk_level(marked_locations):
    """
    Sum risk level of all points combined
    :param points: marked points
    :return: sum of risk level
    """
    counter = 0
    for row in marked_locations:
        for point, is_lowest, _ in row:
            if is_lowest:
                counter += 1 + point

    return counter


def sum_largest_basins(marked_locations):
    """
    Multiply together the sizes of the three largest basins
    :param marked_locations: list of marked locations
    :return: sum of the sizes of the three largest basins
    """
    basins = []
    for i in range(0, len(marked_locations)):
        for j in range(0, len(marked_locations[i])):
            point, _, visited = marked_locations[i][j]
            if not visited and point != 9:
                marked_locations, b = find_basin(marked_locations, i, j)
                if b:
                    basins.append(b)
    basins.sort(key=len, reverse=True)
    result = 1
    for basin in basins[:3]:
        result *= len(basin)
    return result


def find_basin(marked_locations, i, j):
    """
    Recursive function for finding points that belong to a basin
    :param marked_locations: list of marked locations
    :param i: list row index
    :param j: list column index
    :return: modified marked locations and found basins
    """
    basins = []
    if i == len(marked_locations) \
            or j == len(marked_locations[i]) \
            or marked_locations[i][j][0] == 9 \
            or marked_locations[i][j][2] \
            or i == -1 \
            or j == -1:
        return [marked_locations, basins]

    marked_locations[i][j][2] = True
    basins.append(marked_locations[i][j][0])

    # top
    marked_locations, b = find_basin(marked_locations, i - 1, j)
    basins.extend(b)

    # down
    marked_locations, b = find_basin(marked_locations, i + 1, j)
    basins.extend(b)

    # left
    marked_locations, b = find_basin(marked_locations, i, j - 1)
    basins.extend(b)

    # right
    marked_locations, b = find_basin(marked_locations, i, j + 1)
    basins.extend(b)

    return [marked_locations, basins]


if __name__ == '__main__':
    locations = read_file('input.txt')
    # print(locations)
    marked_locations = mark_low_points(locations)
    # print(marked_locations)
    risk_level_sum = sum_risk_level(marked_locations)
    # print(risk_level_sum)
    basin_sum = sum_largest_basins(marked_locations)
    print(basin_sum)
