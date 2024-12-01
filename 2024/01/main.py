import logging


def read_file(filename: str) -> tuple[list[int], list[int]]:
    """
    Read numbers from a file. The file should have two numbers per line, separated by three spaces.
    :param filename: The name of the file to read.
    :return: A tuple with two lists of integers.
    """

    first_numbers = []
    second_numbers = []
    logging.debug("Reading file %s", filename)
    with open(filename, "r", encoding="UTF-8") as file:
        for line in file:
            first, second = line.strip().split("   ")
            # logging.debug("Read numbers: %s %s.", first, second)
            first_numbers.append(int(first))
            second_numbers.append(int(second))

    logging.info("Read %d first numbers and %d second numbers.", len(first_numbers), len(second_numbers))
    logging.debug("First numbers: %s", first_numbers)
    logging.debug("Second numbers: %s", second_numbers)
    return first_numbers, second_numbers


def sort_numbers(first_numbers: list[int], second_numbers: list[int]) -> (list[int], list[int]):
    """
    Sorts the number by value.
    :param first_numbers: a list of first numbers.
    :param second_numbers: a list of second numbers.
    :return: a tuple of sorted lists.
    """

    logging.debug("Sorting numbers.")

    return sorted(first_numbers), sorted(second_numbers)




def calculate_distance(first: list[ int], second: list[int]) -> int:
    """
    Calculates difference between numbers with matching position in both lists.
    For example between first in the first list and the first in the second list.
    Then sums all the differences.
    :param first: sorted list of numbers.
    :param second: sorted list of numbers.
    :return: the sum of differences.
    """

    logging.debug("Calculating distance.")
    distance = 0
    for i in range(len(first)):
        distance += abs(first[i] - second[i])
    logging.info("The distance is %d.", distance)
    return distance

def calculate_similarity_score(first: list[int], second: list[int]) -> int:
    """
    Calculates similarity score between two lists of numbers.
    Calculates a total similarity score by adding up each number in the left list after multiplying it
    by the number of times that number appears in the right list.
    :param first: sorted list of numbers.
    :param second: sorted list of numbers.
    :return: similarity score.
    """

    logging.debug("Calculating similarity score.")
    score = 0
    for i in range(len(first)):
        score += first[i] * second.count(first[i])
    logging.info("The similarity score is %d.", score)
    return score


def main():
    """
    Main function of the program.
    """

    logging.basicConfig(level=logging.DEBUG)
    first_numbers, second_numbers = read_file("input.txt")
    first_sorted, second_sorted = sort_numbers(first_numbers, second_numbers)
    distance = calculate_distance(first_sorted, second_sorted)
    sim_score = calculate_similarity_score(first_sorted, second_sorted)


if __name__ == "__main__":
    main()
