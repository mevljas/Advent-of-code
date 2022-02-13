from collections import defaultdict


def read_file(filename):
    """
    Read input file and save pairs into a dictionary.
    :param filename: input file
    :return: pairs, chars, rules
    """
    pairs = defaultdict(int)
    chars = defaultdict(int)
    rules = dict()
    with open(filename, 'r', encoding='UTF-8') as file:
        for line in file:
            line = line.strip()
            if not line:
                continue
            if not "->" in line:
                line = line.strip()
                chars[line[0]] += 1
                for i in range(0, len(line) - 1, 1):
                    a, b = line[i], line[i + 1]
                    pairs[a + b] += 1
                    # chars[a] += 1
                    chars[b] += 1
            else:
                key, value = line.split("->")
                rules[key.strip()] = value.strip()
        return pairs, chars, rules


def par_insertion(pairs, chars, rules, steps):
    """
    Apply steps of par insertion to the polymer template.
    :param pairs: template pairs
    :param chars: counter of characters
    :param steps: number of required steps
    :return: pairs, chars
    """
    for _ in range(steps):
        for (a, b), value in pairs.copy().items():
            insertion = rules[a + b]
            pairs[a + b] -= value
            pairs[a + insertion] += value
            pairs[insertion + b] += value
            chars[insertion] += value

    return pairs, chars


if __name__ == '__main__':
    pairs, chars, rules = read_file('input.txt')
    # pairs, chars = par_insertion(pairs, chars, rules, 10)
    pairs, chars = par_insertion(pairs, chars, rules, 40)
    print(max(chars.values())-min(chars.values()))
