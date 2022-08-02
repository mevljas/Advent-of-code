def read_file(filename):
    """
    Read input file and save the fish int bins grouped by age.
    :param filename: input file
    :return: dictionary of fish grouped by age
    """
    fish = dict.fromkeys([0, 1, 2, 3, 4, 5, 6, 7, 8], 0)
    with open(filename, "r", encoding="UTF-8") as file:
        for line in file:
            for number in line.split(","):
                fish[int(number)] += 1

    return fish


def simulate_fish(fish, days):
    """
    Simulate fish lives after number of days
    :param fish: dictionary of fish
    :param days: number of days
    :return: modified dictionary of fish
    """
    for i in range(0, days):
        moving_fish = 0
        for j in range(8, -1, -1):
            new_fish = moving_fish
            if j == 0:
                fish[6] += fish[j]
                fish[8] += fish[j]
                fish[j] = new_fish
            else:
                moving_fish = fish[j]
                fish[j] = new_fish

    return fish


def count_fish(fish):
    """
    Count how many fish are alive.
    :param fish: dictionary of fish
    :return: number of living fish
    """
    counter = 0
    for f in fish.values():
        counter += f

    return counter


if __name__ == "__main__":
    fish = read_file("input.txt")
    # fish = simulate_fish(fish, 80)
    fish = simulate_fish(fish, 256)
    print(count_fish(fish))
