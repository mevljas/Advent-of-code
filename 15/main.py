import collections
import copy
from warnings import warn
import heapq



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


# A Naive recursive implementation of MCP(Minimum Cost Path) problem

import sys


# Returns cost of minimum cost path from (0,0) to (m, n) in mat[R][C]
def minCost(cost, m, n):
    if (n < 0 or m < 0):
        return sys.maxsize
    elif (m == 0 and n == 0):
        return 0
    else:
        return cost[m][n] + min(minCost(cost, m - 1, n),
                                minCost(cost, m, n - 1))



def main():
    maze = read_file('input.txt')
    maze = multiply_matrix(maze, 5)


    # G Driver program to test above functions
    print(minCost(maze, len(maze) - 1, len(maze[0]) - 1))



if __name__ == '__main__':
   main()
