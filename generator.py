import os
from config import ROOT_DIR
from time import time
from sys import getsizeof
from random import shuffle


def gen():
    # 8-9**-***-**-**
    start_time = time()
    file = open(os.path.join(ROOT_DIR, 'gen_numbers.txt'), 'w')
    numbers = []
    for i in range(9):
        if i == 8:
            numbers = [j for j in range(9000000000 + 111111111*i, 9000000000 + 111111111*(i+1)+1)]
        else:
            numbers = [j for j in range(9000000000 + 111111111*i, 9000000000 + 111111111*(i+1))]
        shuffle(numbers)
        shuffle(numbers)
        for j in numbers:
            file.write("8" + str(j) + "\n")
        size_ = getsizeof(numbers)
        numbers.clear()

    print('Phone numbers were generated in {} minutes!'.format((time() - start_time)/60))
    print("Generated List memory size: {} bytes".format(size_))
    print()
    file.close()
