def read_file(filename):
    """
    Read input file and save the crabs into a list.
    :param filename: input file
    :return: list of crabs
    """
    crabs = []
    with open(filename, "r", encoding="UTF-8") as file:
        for line in file:
            for number in line.split(","):
                crabs.append(int(number))

    return crabs


def find_cheapest_fuel(crabs):
    """
    Find cheapest fuel consumption combination
    :param crabs: list of crabs
    :return: cheapest sum of fuel consumption
    """
    min_fuel_sum = sum(crabs)
    for align_position in range(min(crabs), max(crabs) + 1):
        current_fuel_sum = 0
        for crab_position in crabs:
            current_fuel_sum += abs(align_position - crab_position)

        if current_fuel_sum < min_fuel_sum:
            min_fuel_sum = current_fuel_sum

    return min_fuel_sum


def find_cheapest_fuel_improved(crabs):
    """
    Find cheapest fuel consumption combination with different steps
    :param crabs: list of crabs
    :return: cheapest sum of fuel consumption
    """
    fuel_consumption = lambda crab_position, new_position: sum(
        cost for cost in range(abs(crab_position - new_position) + 1)
    )

    return min(
        [
            sum(
                fuel_consumption(crab_position, new_position) for crab_position in crabs
            )
            for new_position in range(max(crabs))
        ]
    )


if __name__ == "__main__":
    crabs = read_file("input.txt")
    min_fuel_consumption = find_cheapest_fuel(crabs)
    min_fuel_consumption_improved = find_cheapest_fuel_improved(crabs)
    print(crabs)
    print(min_fuel_consumption)
    print(min_fuel_consumption_improved)
