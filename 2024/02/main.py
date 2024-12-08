import logging


def read_file(filename: str):
    reports = []
    logging.debug("Reading file %s", filename)
    with open(filename, "r", encoding="UTF-8") as file:
        for line in file:
            levels = [int(x) for x in line.split(" ")]
            reports.append(levels)

    return reports


def check_sequence_direction(level) -> bool:
    """
    Check whether all numbers are increasing or decreasing
    :param level: list of numbers
    :return: whether the sequence is increasing or decreasing
    """

    increasing = True
    decreasing = True
    for i in range(1, len(level)):
        if level[i] > level[i - 1]:
            decreasing = False
        elif level[i] < level[i - 1]:
            increasing = False

        if not increasing and not decreasing:
            return False

    return increasing or decreasing


def check_sequence_difference(level) -> bool:
    """
    Check whether adjacent numbers differ by at least 1 and at most three.
    :param level: list of numbers
    :return: whether the sequence is valid
    """

    for i in range(1, len(level)):
        diff = abs(level[i] - level[i - 1])
        if diff < 1 or diff > 3:
            return False

    return True

def problem_damper(level) -> bool:
    """
    Tries to remove each number from the list and check if the sequence is now valid.
    :param level: list of numbers
    :return: whether the modified sequence is valid.
    """

    for i in range(len(level)):
        modified_level = level[:i] + level[i + 1:]
        if check_sequence_difference(modified_level) and check_sequence_direction(modified_level):
            return True

    return False


def analyze_data(reports):
    safe_reports_count = 0

    for level in reports:
        logging.debug("Analyzing level %s", level)
        if check_sequence_direction(level) and check_sequence_difference(level):
            safe_reports_count += 1
        elif problem_damper(level):
            safe_reports_count += 1

    logging.info("Analyzing complete")

    return safe_reports_count


def main():
    """
    Main function of the program.
    """
    logging.basicConfig(level=logging.DEBUG)

    reports = read_file("input.txt")
    safe_reports_count = analyze_data(reports)
    logging.info("The number of safe reports is %d.", safe_reports_count)


if __name__ == "__main__":
    main()
