def read_file(filename):
    """
    Read number from a file and save them to a list.
    :param filename: file containing input numbers.
    :return: list of read numbers
    """
    numbers = []
    with open(filename, "r", encoding="UTF-8") as file:
        for line in file:
            numbers.append(int(line.strip()))

    return numbers


def count_greater_numbers(numbers):
    """
    Count cases when the next number was greater than the one before it.
    :param numbers: list of numbers.
    :return:
    """
    previous_number = None
    counter = 0
    for number in numbers:
        if previous_number and number >= previous_number:
            counter += 1
        previous_number = number

    print(f"Number of greater values is {counter}.")


def count_greater_windows(numbers):
    """
    Count cases when the next window ob numbers was greater than the one before it.
    :param numbers: list of numbers.
    :return:
    """
    previous_sum = None
    counter = 0

    for i in range(0, len(numbers) - 2):
        current_sum = sum(numbers[i : i + 3])
        if previous_sum and current_sum > previous_sum:
            counter += 1

        previous_sum = current_sum

    print(f"Number of greater windows is {counter}.")


if __name__ == "__main__":
    numbers = read_file("input.txt")
    count_greater_numbers(numbers)
    count_greater_windows(numbers)
