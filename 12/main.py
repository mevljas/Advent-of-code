from collections import defaultdict
from copy import deepcopy


class Graph:
    paths = []

    def __init__(self):
        """
        Constructor
        """

        # default dictionary to store graph
        self.graph = defaultdict(list)

    def add_ege(self, u, v):
        """
        function to add an edge to graph
        :param u:
        :param v:
        :return:
        """
        self.graph[u].append(v)

    def dfs_util(self, v, visited, path):
        """
        A function used by DFS
        :param v: current node
        :param visited: set of visited nodes
        :param path: list of nodes on a path
        :return: None
        """
        # Deep copy variables
        path = deepcopy(path)
        visited = deepcopy(visited)

        # Mark the current node as visited
        # and print it
        # visited.add(v)
        path.append(v)
        # print(v, end=' ')

        if v == 'end':
            self.paths.append(path)
            return

        if v.islower() and v not in {'start', 'end'}:
            visited.add(v)

        # Recur for all the vertices
        # adjacent to this vertex
        for neighbour in set(self.graph[v]).difference(visited):
            self.dfs_util(neighbour, visited, path)

        return

    def dfs(self, v):
        """
        The function to do DFS traversal. It uses
        recursive DFSUtil()
        :param v: first node
        :return: dfs traversal
        """

        # Create a set to store visited vertices
        visited = {'start'}

        # Call the recursive helper function
        # to find DFS traversal
        self.dfs_util(v, visited, [])
        return self.paths


def read_file(filename):
    """
    Read input file and save cave connections into a graph.
    :param filename: input file
    :return: graph of nodes
    """
    graph = Graph()
    with open(filename, 'r', encoding='UTF-8') as file:
        for line in file:
            a, b = line.strip().split("-")
            graph.add_ege(a, b)
            graph.add_ege(b, a)

    return graph


if __name__ == '__main__':
    graph = read_file('input.txt')
    dfs = graph.dfs('start')
    print(len(dfs))
