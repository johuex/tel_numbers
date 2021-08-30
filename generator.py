import os
from config import ROOT_DIR
from time import time
from random import shuffle, randint, choice


def gen():
    # 8-9**-***-**-**
    start_time = time()
    file = open(os.path.join(ROOT_DIR, 'gen_numbers.txt'), 'w')
    for i in range(27):
        if i == 26:
            numbers = [str(j) + '\n' for j in range(9000000000 + 37037037*i, 9000000000 + 37037037*(i+1)+1)]
        else:
            numbers = [str(j) + '\n' for j in range(9000000000 + 37037037*i, 9000000000 + 37037037*(i+1))]
        # print("Generating number: {} seconds".format(time() - start_time))
        shuffle_time = time()
        len_num = len(numbers) - 1
        for j in range(len_num//2):
            end_i = len_num-j
            rand_i = randint(0, end_i//2-1)
            numbers[end_i], numbers[rand_i] = numbers[rand_i], numbers[end_i]
        # print("Shuffle list: {} seconds".format(time() - shuffle_time))
        write_time = time()
        file.writelines(numbers)
        # print("Write time: {} seconds".format(time() - write_time))
        clear_time = time()
        numbers.clear()
        # print("Clear time: {} seconds".format(time() - clear_time))

    print('Phone numbers were generated in {} seconds!'.format((time() - start_time)))
    print()
    file.close()
