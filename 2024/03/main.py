import logging


def read_file(filename: str):
    input = []
    logging.debug("Reading file %s", filename)
    with open(filename, "r", encoding="UTF-8") as file:
        for line in file:
            input.append(line)

    return reports


def main():
    """
    Main function of the program.
    """
    logging.basicConfig(level=logging.DEBUG)

    input = read_file("input.txt")


if __name__ == "__main__":
    main()
