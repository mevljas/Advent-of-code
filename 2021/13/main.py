import collections


def read_file(filename):
    """
    Read input file and save dots and folds.
    :param filename: input file
    :return: list of dots, list of folds
    """
    dots = []
    folds = []
    max_x = 0
    max_y = 0
    with open(filename, "r", encoding="UTF-8") as file:
        for line in file:
            if "," in line:
                x, y = line.strip().split(",")
                x = int(x)
                y = int(y)
                dots.append([x, y])
                max_x = max(max_x, x)
                max_y = max(max_y, y)
            if "=" in line:
                _, _, statement = line.split(" ")
                axis, value = statement.split("=")
                folds.append([axis, int(value.strip())])
    max_x += 1
    max_y += 1
    paper = [[0] * max_x for _ in range(max_y)]
    for x, y in dots:
        paper[y][x] = 1

    return paper, folds


def fold_paper_once(paper, folds):
    """
    Fold paper only once
    :param paper: paper with dots
    :param folds: required folds
    :return: folded paper once
    """
    for axis, value in folds:
        if axis == "x":
            for i in range(0, len(paper)):
                for j in range(value + 1, len(paper[i])):
                    if paper[i][j]:
                        paper[i][value - (j - value)] = paper[i][j]
                        paper[i][j] = 0
        elif axis == "y":
            for i in range(value + 1, len(paper)):
                for j in range(0, len(paper[i])):
                    if paper[i][j]:
                        paper[value - (i - value)][j] = paper[i][j]
                        paper[i][j] = 0
        return


def count_dots(paper):
    """
    Count dots on the paper
    :param paper: paper with dots
    :return: number of dots
    """
    counter = 0
    for row in paper:
        for dot in row:
            if dot:
                counter += 1
    return counter


def fold_paper(paper, folds):
    """
    Fold paper
    :param paper: paper with dots
    :param folds: required folds
    :return: folded paper
    """
    for axis, value in folds:
        if axis == "x":
            for i in range(0, len(paper)):
                for j in range(value + 1, len(paper[i])):
                    if paper[i][j]:
                        paper[i][value - (j - value)] = paper[i][j]
                        paper[i][j] = 0
        elif axis == "y":
            for i in range(value + 1, len(paper)):
                for j in range(0, len(paper[i])):
                    if paper[i][j]:
                        paper[value - (i - value)][j] = paper[i][j]
                        paper[i][j] = 0


def save_to_file(paper):
    """
    Save code to a file for easier viewing.
    :param paper: paper with dots
    :return:
    """
    f = open("solution.txt", "w")
    for row in paper:
        for dot in row:
            if dot:
                f.write("#")
            else:
                f.write(".")
        f.write("\n")

    f.close()


if __name__ == "__main__":
    paper, folds = read_file("input.txt")
    # print(paper)
    # print(folds)
    # fold_paper_once(paper, folds)
    fold_paper(paper, folds)
    # print(paper)
    # print(count_dots(paper))
    save_to_file(paper)
