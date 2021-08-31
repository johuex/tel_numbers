import os
from config import ROOT_DIR
from time import time
from random import randint


def gen():
    # 8-9**-***-**-**
    start_time = time()
    file = open(os.path.join(ROOT_DIR, 'gen_numbers.txt'), 'w')
    for i in range(27):
        if i == 26:
            numbers = [j for j in range(9000000000 + 37037037*i, 9000000000 + 37037037*(i+1)+1)]
        else:
            numbers = [j for j in range(9000000000 + 37037037*i, 9000000000 + 37037037*(i+1))]
        len_num = len(numbers) - 1
        for j in range(len_num//2):
            end_i = len_num-j
            rand_i = randint(0, end_i//2-1)
            numbers[end_i], numbers[rand_i] = numbers[rand_i], numbers[end_i]
        file.write(''.join('8' + str(line) + '\n' for line in numbers))
        numbers.clear()
    print('Phone numbers were generated in {} seconds!'.format((time() - start_time)))
    print()
    file.close()
