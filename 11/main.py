def read_file(filename):
    """
    Read input file and save the lines into a list.
    :param filename: input file
    :return: grid of octopuses
    """
    grid = []
    with open(filename, 'r', encoding='UTF-8') as file:
        for line in file:
            grid.append([int(s) for s in list(line.strip())])
    return grid


def model_light(grid, steps):
    """
    Simulate steps
    :param grid: grid of octopuses
    :param steps: number of steps to simulate
    :return: number of flashes
    """
    flashes = 0
    for step in range(0, steps):
        previous_flashes = -1
        for i in range(0, len(grid)):
            for j in range(0, len(grid[i])):
                if grid[i][j]:
                    grid[i][j] += 1
                else:
                    grid[i][j] = 1

        while previous_flashes != flashes:
            previous_flashes = flashes
            for i in range(0, len(grid)):
                for j in range(0, len(grid[i])):
                    if grid[i][j] and grid[i][j] > 9:
                        flashes += 1
                        grid[i][j] = None
                        # Up
                        if i > 0 and grid[i - 1][j]:
                            grid[i - 1][j] += 1
                        # Down
                        if i < len(grid) - 1 and grid[i + 1][j]:
                            grid[i + 1][j] += 1
                        #
                        if j > 0 and grid[i][j - 1]:
                            grid[i][j - 1] += 1
                        if j < len(grid[i]) - 1 and grid[i][j + 1]:
                            grid[i][j + 1] += 1
                        if i > 0 and j > 0 and grid[i - 1][j - 1]:
                            grid[i - 1][j - 1] += 1
                        if i > 0 and j < len(grid[i]) - 1 and grid[i - 1][j + 1]:
                            grid[i - 1][j + 1] += 1
                        if i < len(grid) - 1 and j < len(grid[i]) - 1 and grid[i + 1][j + 1]:
                            grid[i + 1][j + 1] += 1
                        if i < len(grid) - 1 and j > 0 and grid[i + 1][j - 1]:
                            grid[i + 1][j - 1] += 1
    return flashes

def all_flashing_steps(grid):
    """
    Find the exact moments when the octopuses will all flash simultaneously
    :param grid: grid of octopuses
    :return: number of steps required for the exact moments when the octopuses will all flash simultaneously
    """
    steps = 0
    flashes = 0
    n_of_octopuses = len(grid) * len(grid[0])
    current_flashes = 0
    flashed = 0
    while(flashed != n_of_octopuses):
        steps += 1
        previous_flashes = -1
        for i in range(0, len(grid)):
            for j in range(0, len(grid[i])):
                if grid[i][j]:
                    grid[i][j] += 1
                else:
                    grid[i][j] = 1
        flashed = 0
        while previous_flashes != flashes:
            previous_flashes = flashes
            current_flashes = 0
            for i in range(0, len(grid)):
                for j in range(0, len(grid[i])):
                    if grid[i][j] and grid[i][j] > 9:
                        current_flashes += 1
                        flashed += 1
                        grid[i][j] = None
                        # Up
                        if i > 0 and grid[i - 1][j]:
                            grid[i - 1][j] += 1
                        # Down
                        if i < len(grid) - 1 and grid[i + 1][j]:
                            grid[i + 1][j] += 1
                        #
                        if j > 0 and grid[i][j - 1]:
                            grid[i][j - 1] += 1
                        if j < len(grid[i]) - 1 and grid[i][j + 1]:
                            grid[i][j + 1] += 1
                        if i > 0 and j > 0 and grid[i - 1][j - 1]:
                            grid[i - 1][j - 1] += 1
                        if i > 0 and j < len(grid[i]) - 1 and grid[i - 1][j + 1]:
                            grid[i - 1][j + 1] += 1
                        if i < len(grid) - 1 and j < len(grid[i]) - 1 and grid[i + 1][j + 1]:
                            grid[i + 1][j + 1] += 1
                        if i < len(grid) - 1 and j > 0 and grid[i + 1][j - 1]:
                            grid[i + 1][j - 1] += 1
            flashes += current_flashes
    return steps


if __name__ == '__main__':
    grid = read_file('input.txt')
    # flashes = model_light(grid, 100)
    steps = all_flashing_steps(grid)
    print(steps)
