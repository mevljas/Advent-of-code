import statistics


def read_file(filename):
    """
    Read input file and save the lines into a list.
    :param filename: input file
    :return: list of input lines
    """
    navigation_subsystem = []
    with open(filename, 'r', encoding='UTF-8') as file:
        for line in file:
            navigation_subsystem.append(list(line.strip()))
    return navigation_subsystem


def calculate_syntax_error_score(navigation_subsystem):
    """
    Calculate syntax error score
    :param navigation_subsystem:  list of chunks
    :return: syntax error score and filtered chunks
    """
    score = 0
    uncomplete_lines = []

    for line in navigation_subsystem:
        stack = []
        corrupted_line = False
        for character in line:
            if character in ('(', '[', '{', '<'):
                stack.append(character)
            else:
                stack_pop = stack.pop() or ""
                if character == ')' and stack_pop != '(':
                    score += 3
                    corrupted_line = True
                    continue

                elif character == ']' and stack_pop != '[':
                    score += 57
                    corrupted_line = True
                    continue

                elif character == '}' and stack_pop != '{':
                    score += 1197
                    corrupted_line = True
                    continue

                elif character == '>' and stack_pop != '<':
                    score += 25137
                    corrupted_line = True
                    continue
        if not corrupted_line:
            uncomplete_lines.append(stack)

    return score, uncomplete_lines


def repair_navigation_subsystem(uncomplete_lines):
    """
    Calculate middle score
    :param uncomplete_lines: chunks with missing values
    :return: median score of all lines
    """
    scores = []
    for line in uncomplete_lines:
        score = 0
        for character in reversed(line):
            score *= 5
            if character == '(':
                score += 1
            elif character == '[':
                score += 2
            elif character == '{':
                score += 3
            elif character == '<':
                score += 4

        scores.append(score)

    return statistics.median(scores)


if __name__ == '__main__':
    navigation_subsystem = read_file('input.txt')
    syntax_error_score, uncomplete_lines = calculate_syntax_error_score(navigation_subsystem)
    print(syntax_error_score)
    middle_score = repair_navigation_subsystem(uncomplete_lines)
    print(middle_score)
