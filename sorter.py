import os
from config import ROOT_DIR


def sorter():
    f_in = open(os.path.join(ROOT_DIR, 'gen_numbers.txt'))
    f_out = open(os.path.join(ROOT_DIR, 'sort_numbers.txt'), 'w')
    sorted_numbers = []
    for line in f_in:
        # 8-9**-***-**-**
        number = int(line[:12])
        if len(sorted_numbers) == 0:
            sorted_numbers.append(number)
            continue
        for i in range(len(sorted_numbers)):
            if sorted_numbers[i] < number:
                sorted_numbers = sorted_numbers[0:i] + [number] + sorted_numbers[i:]
                break

    for num in sorted_numbers:
        f_out.write("%d\n" % num)

    f_in.close()
    f_out.close()
