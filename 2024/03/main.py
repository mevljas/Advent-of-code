import logging
import re


def read_file(filename: str):
    input = ""
    logging.debug("Reading file %s", filename)
    with open(filename, "r", encoding="UTF-8") as file:
        input = file.read()

    return input


def find_instruction(input: str):
    return re.findall("(?<=mul\()\d{1,3},\d{1,3}(?=\))", input)


def execute_instructions(instructions: [str]):
    execute_sum = 0
    for instruction in instructions:
        a, b = instruction.split(",")
        a, b = int(a), int(b)
        res = a * b
        execute_sum += res
    return execute_sum


def sum_enabled_multiplications(input: str):
    mul_pattern = re.compile(r"mul\((\d+),(\d+)\)")
    state_change_pattern = re.compile(r"(do\(\)|don't\(\))")

    enabled = True
    sum = 0

    tokens = re.split(r"(?=mul\(|do\(\)|don't\(\))", input)
    for token in tokens:
        # Check for state change
        state_change = state_change_pattern.match(token)
        if state_change:
            if state_change.group() == "do()":
                enabled = True
            elif state_change.group() == "don't()":
                enabled = False
        # Check for mul instructions
        instruction = mul_pattern.match(token)
        if instruction and enabled:
            x, y = map(int, instruction.groups())
            sum += x * y

    return sum


def main():
    """
    Main function of the program.
    """
    logging.basicConfig(level=logging.DEBUG)

    input = read_file("input.txt")
    logging.debug("Input: %s", input)

    instructions = find_instruction(input)
    logging.debug("Instructions: %s", instructions)

    res = execute_instructions(instructions)
    logging.info("Result: %d", res)

    res2 = sum_enabled_multiplications(input)
    logging.info("Result 2: %d", res2)




if __name__ == "__main__":
    main()
