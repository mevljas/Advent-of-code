import copy


def read_file(filename):
    """
    Read input file and save random numbers and bingo boards.
    :param filename: input file
    :return: random numbers and bingo boards
    """
    called_numbers = []
    boards = []
    with open(filename, 'r', encoding='UTF-8') as file:
        called_numbers = file.readline().strip().split(',')
        for line in file:
            board = []
            for i in range(0, 5):
                line = file.readline().strip()
                row = []
                for number in line.split(" "):
                    if number.isdigit():
                        row.append([int(number), False])
                board.append(row)
            boards.append(board)

    return [called_numbers, boards]


def Mark_boards(called_number, boards):
    """
    Loop through multiple boards and mark found numbers.
    :param called_number: tombola number
    :param boards: tombola boards
    :return: marked tombola boards
    """
    for i in range(0, len(boards)):
        boards[i] = mark_board(called_number, boards[i])
    return boards


def mark_board(called_number, board):
    """
    Mark found numbers on a boars
    :param called_number:  tombola number
    :param board: tombola board
    :return: marked tombola board
    """
    for j in range(0, len(board)):
        for z in range(0, len(board[j])):
            for k in range(0, len(board[j][z])):
                if int(called_number) == int(board[j][z][0]):
                    board[j][z][1] = True

    return board


def find_first_board(called_numbers, boards):
    """
    Find the board that wins first.
    :param called_numbers: tombola numbers
    :param boards: tombola boards
    :return: score of the winning boards - sum of all unmarked numbers * sum of all unmarked numbers
    """
    for called_number in called_numbers:
        boards = Mark_boards(called_number, boards)
        result = check_boards(boards)
        if result:
            return result * int(called_number)

    return False


def check_boards(boards):
    """
    Check if the game is won
    :param boards: tombola boards
    :return: sum ob unmarked numbers on the board
    """
    for board in boards:
        result = check_board(board)
        if result:
            return result
    return False


def check_board(board):
    """
    Check if the board  won
    :param board: tombola board
    :return: sum ob unmarked numbers on the board
    """
    # Horizontal
    for i in range(0, len(board)):
        for j in range(0, len(board[i])):
            if not bool(board[i][j][1]):
                break
        else:
            return sum_unmarked(board)

    # Vertical
    for j in range(0, len(board[0])):
        for i in range(0, len(board)):
            if not bool(board[i][j][1]):
                break
        else:
            return sum_unmarked(board)

    return False


def sum_unmarked(board):
    """
    Sum of the unmarked pieces on the board
    :param board:
    :return:
    """
    counter = 0
    for row in board:
        for pair in row:
            number, marked = pair
            if not marked:
                counter += number
    return counter


def find_last_board(called_numbers, boards):
    """
    Find the board that wins last.
    :param called_numbers: tombola numbers
    :param boards: tombola boards
    :return: score of the winning boards - sum of all unmarked numbers * sum of all unmarked numbers
    """

    called_number = int(called_numbers[0])
    boards = Mark_boards(called_number, boards)
    boards_copy = copy.deepcopy(boards)
    for board in boards:
        result = check_board(board)
        if result:
            boards_copy.remove(board)
            last_result = result * int(called_number)
            if len(boards_copy) == 0:
                return last_result

    return find_last_board(called_numbers[1:], boards_copy)

if __name__ == '__main__':
    random_numbers, boards = read_file('input.txt')
    # result = find_first_board(random_numbers, boards)
    result = find_last_board(random_numbers, boards)
    print(f'Result is {result}')
