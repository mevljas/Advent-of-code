def read_file(filename):
    """
    Read commands from a file and save them to a list.
    :param filename: file containing input commands.
    :return: list of read commands
    """

    commands = {}
    with open(filename, "r", encoding="UTF-8") as file:
        for line in file:
            command, unit = line.strip().split(" ")
            unit = int(unit)
            if command in commands:
                commands[command] += unit
            else:
                commands[command] = unit

    return commands


def calculate_position(commands):
    """
    Calculate submarine position after executing commands (aim excluded).
    :param commands: list of read commands
    :return: final location
    """

    horizontal_position = commands.get("forward", 0)
    depth = commands.get("down", 0) - commands.get("up", 0)
    result = horizontal_position * depth
    return result


def calculate_position_aim(filename):
    """
    ead commands from a file and calculate submarine position after executing commands (aim included).
    :param filename: file containing input commands.
    :return: final location
    """

    aim = 0
    horizontal_position = 0
    depth = 0
    with open(filename, "r", encoding="UTF-8") as file:
        for line in file:
            command, unit = line.strip().split(" ")
            unit = int(unit)
            if command == "down":
                aim += unit

            elif command == "up":
                aim -= unit

            elif command == "forward":
                horizontal_position += unit
                depth += unit * aim

    result = horizontal_position * depth

    return result


if __name__ == "__main__":
    commands = read_file("input.txt")
    # result = calculate_position(commands)
    result = calculate_position_aim("input.txt")
    print(result)
