import copy


def read_file(filename):
    """
    Read input file and save risk levels into a list.
    :param filename: input file
    :return: matrix
    """
    matrix = []
    with open(filename, 'r', encoding='UTF-8') as file:
        for line in file:
            line = [int(d) for d in line.strip()]
            matrix.append(line)

        return matrix


def build_graph(matrix):
    graph = []
    for i in range(0, len(matrix)):
        for j in range(0, len(matrix[i])):
            current_node = i * len(matrix) + j
            if i < len(matrix) - 1:
                # down
                graph.append((current_node + len(matrix[i + 1]), current_node, matrix[i][j]))
            if i > 0:
                # up
                graph.append((current_node - len(matrix[i - 1]), current_node, matrix[i][j]))
            if j < len(matrix[0]) - 1:
                # right
                graph.append((current_node + 1, current_node, matrix[i][j]))
            if j > 0:
                # left
                graph.append((current_node - 1, current_node, matrix[i][j]))

    return graph


def multiply_matrix(matrix, n):
    new_matrix = []
    for line in matrix:
        new_line = []
        for i in range(0, n):
            for number in line:
                new_number = number + i
                new_number = new_number - 9 if new_number > 9 else new_number
                new_line.append(new_number)
        new_matrix.append(new_line)

    matrix = copy.deepcopy(new_matrix)
    new_matrix = []
    for i in range(0, n):
        for line in matrix:
            new_line = []
            for number in line:
                new_number = number + i
                new_number = new_number - 9 if new_number > 9 else new_number
                new_line.append(new_number)
            new_matrix.append(new_line)

    return new_matrix


if __name__ == '__main__':
    matrix = read_file('input.txt')
    multiplied_matrix = multiply_matrix(matrix, 5)
    graph = build_graph(matrix)
