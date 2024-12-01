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
            first, second = line.split("   ")
            logging.debug("Read numbers: %s %s.", first, second)
            first_numbers.append(int(first))
            second_numbers.append(int(second))

    logging.info("Read %d first numbers and %d second numbers.", len(first_numbers), len(second_numbers))
    return first_numbers, second_numbers


if __name__ == "__main__":
    first_numbers, second_numbers = read_file("input.txt")
    print(first_numbers)
    print(second_numbers)
