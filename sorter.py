import os
from config import ROOT_DIR
from time import time
from sys import getsizeof


def sorter_1():
    start_split = time()
    file_name = "temp_{}"
    # split input file into several parts
    files = [None]*9  # saves i/o
    temp_numbers = []
    j = 0
    i = 0
    print("Start split and sort!")
    with open(os.path.join(ROOT_DIR, 'gen_numbers.txt')) as f_in:
        for line in f_in:  # 1, 2
            j += 1
            if j % 111111111 == 0:
                files[i] = open(os.path.join(ROOT_DIR, file_name.format(i)), 'w+')
                temp_numbers.sort()  # linear sort by Python
                for z in temp_numbers:
                    files[i].write(str(z) + "\n")
                i += 1
                temp_numbers.clear()
            temp_numbers.append(int(line))
    print("Split and sorting parts time: %s seconds" % (time() - start_split))
    '''i = 7
    for l in range(i+1):
        files[l] = open(os.path.join(ROOT_DIR, file_name.format(l)))'''
    # external sorting
    print("Start external sorting")
    start_sort = time()
    f_out = open(os.path.join(ROOT_DIR, 'sort1_numbers.txt'), 'w')  # 3
    temp_numbers.clear()

    for p in range(i+1):  # 4
        temp_numbers.append(int(files[p].readline()[:11]))
    while True:
        min_i = temp_numbers.index(min(temp_numbers))  # 5.1
        f_out.write(str(temp_numbers[min_i]) + "\n")  # 5.2
        k = files[min_i].readline()[:11]  # 5.3
        if k == "" or k == "\n":  # 5.4
            temp_numbers.pop(min_i)
            files[min_i].close()
            files.pop(min_i)
            os.remove(os.path.join(ROOT_DIR, file_name.format(min_i)))
        else:
            temp_numbers[min_i] = int(k)
        if len(files) == 0:
            break

    print("External sorting time: %s seconds" % (time() - start_sort))


def sorter_2():
    """Converting str to int and sorting just in time"""
    start_time = time()
    sorted_numbers = []
    with open(os.path.join(ROOT_DIR, 'gen_numbers.txt')) as f_in:
        for line in f_in:
            # 8-9**-***-**-**\n
            number = line[:12]
            if len(sorted_numbers) == 0:
                sorted_numbers.append(number)
                continue
            for i in range(len(sorted_numbers)):
                if sorted_numbers[i] < number:
                    sorted_numbers = sorted_numbers[0:i] + [number] + sorted_numbers[i:]
                    break

    with open(os.path.join(ROOT_DIR, 'sort2_numbers.txt'), 'w') as f_out:
        for num in sorted_numbers:
            f_out.write("%s\n" % num)
    print("Linear sorting time: %s seconds" % (time() - start_time))
    print("List memory size: {} bytes".format(getsizeof(sorted_numbers)))


