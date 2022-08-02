def read_file(filename):
    """
    Read report from a file and save them to a list.
    :param filename: file containing input report.
    :return: list of read binary numbers
    """

    report = []
    with open(filename, "r", encoding="UTF-8") as file:
        for line in file:
            report.append(line.strip())

    return report


def power_consumption(report):
    """
    Decode and check power consumption
    :param report: list of read binary numbers
    :return: power consumption
    """
    counter_of_bits = [None] * len(report[0])

    for position in report:
        for j in range(0, len(position)):
            if not counter_of_bits[j]:
                counter_of_bits[j] = {"one_counter": 0, "zero_counter": 0}
            if position[j] == "1":
                counter_of_bits[j]["one_counter"] += 1
            else:
                counter_of_bits[j]["zero_counter"] += 1

    gamma_rate = ""
    epsilon_rate = ""
    for bits in counter_of_bits:
        if bits["one_counter"] > bits["zero_counter"]:
            gamma_rate += "1"
            epsilon_rate += "0"
        else:
            gamma_rate += "0"
            epsilon_rate += "1"

    result = int(gamma_rate, 2) * int(epsilon_rate, 2)
    return result


def life_support_rating(report):
    """
    Determine the life support rating value.
    :param report:  list of read binary numbers
    :return: life support rating
    """
    oxygen = oxygen_generator_rating(report, 0)
    CO2 = CO2_scrubber_rating(report, 0)

    return int(oxygen, 2) * int(CO2, 2)


def oxygen_generator_rating(report, index):
    """
    Determine the oxygen generator rating value.
    :param report:  list of read binary numbers
    :return: oxygen generator rating
    """

    ones = 0
    zeros = 0

    if len(report) == 1:
        return report[0]

    for number in report:
        if int(number[index]):
            ones += 1
        else:
            zeros += 1

    if ones >= zeros:
        return oxygen_generator_rating([x for x in report if int(x[index])], index + 1)
    else:
        return oxygen_generator_rating(
            [x for x in report if not int(x[index])], index + 1
        )


def CO2_scrubber_rating(report, index):
    """
    Determine the CO2 scrubber rating value.
    :param report:  list of read binary numbers
    :return: CO2 scrubber rating
    """
    ones = 0
    zeros = 0

    if len(report) == 1:
        return report[0]

    for number in report:
        if int(number[index]):
            ones += 1
        else:
            zeros += 1

    if ones >= zeros:
        return CO2_scrubber_rating([x for x in report if not int(x[index])], index + 1)
    else:
        return CO2_scrubber_rating([x for x in report if int(x[index])], index + 1)


if __name__ == "__main__":
    report = read_file("input.txt")
    result = life_support_rating(report)
    print(result)
