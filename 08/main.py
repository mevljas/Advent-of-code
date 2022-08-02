def read_file(filename):
    """
    Read input file and save the signals into a list.
    :param filename: input file
    :return: list of signals
    """
    signals = []
    with open(filename, "r", encoding="UTF-8") as file:
        for line in file:
            unique_signal_patterns, four_digit_output = line.rstrip().split(" | ")
            signals.append(
                [unique_signal_patterns.split(" "), four_digit_output.split(" ")]
            )

    return signals


def count_unique_number_of_segments(signals):
    """
    Count number of unique segments
    :param signals: list of signals
    :return: counter
    """
    counter = 0
    for signal in signals:
        unique_signal_patterns, four_digit_output = signal
        for segments in four_digit_output:
            if len(segments) in (2, 4, 3, 7):
                counter += 1

    return counter


def decode_digit(digit, decode_table):
    """
    Decode digit and return number
    :param digit: input data digit
    :param decode_table: table for decoding digits
    :return: value of the digit
    """
    for key, item in decode_table.items():
        if len(item) == len(digit) and set(filter(str.isalnum, digit)) == item:
            return key

    return False


def sum_output_values(signals):
    """
    Sum decoded output values
    :param signals: input signals
    :return: sum of output values
    """
    sum = 0
    for signal in signals:
        unique_signal_patterns, four_digit_output = signal
        decode_table = generate_decode_table(unique_signal_patterns)
        output_value = ""
        for digit in four_digit_output:
            output_value += str(decode_digit(digit, decode_table))

        sum += int(output_value)

    return sum


def generate_decode_table(unique_signal_patterns):
    """
    Generate table for easier decoding of input data
    :param unique_signal_patterns: input data
    :return: decoding table
    """
    dt = {}  # decode table
    for signal in unique_signal_patterns:
        if len(signal) == 2:
            dt[1] = set(signal)
        elif len(signal) == 4:
            dt[4] = set(signal)
        elif len(signal) == 3:
            dt[7] = set(signal)
        elif len(signal) == 7:
            dt[8] = set(signal)

    # We already have 1, 4, 7, 8
    # find other codes
    for signal in unique_signal_patterns:
        length = len(signal)
        if length in (2, 4, 3, 7):
            continue

        signal = set(signal)
        if length == 6:
            if len(dt[1].difference(signal)) == 1:
                dt[6] = signal
            elif len(dt[4].difference(signal)) == 1:
                dt[0] = signal
            else:
                dt[9] = signal
        elif length == 5:
            if len(dt[1].difference(signal)) == 0:
                dt[3] = signal
            elif len(dt[4].difference(signal)) == 2:
                dt[2] = signal
            else:
                dt[5] = signal

    return dt


if __name__ == "__main__":
    signals = read_file("input.txt")
    # print(signals)
    # unique_number_of_segments = count_unique_number_of_segments(signals)
    # print(unique_number_of_segments)
    sum_of_output_values = sum_output_values(signals)
    print(sum_of_output_values)
