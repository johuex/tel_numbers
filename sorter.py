import os
from config import ROOT_DIR
from time import time


def sorter_1():
    start_split = time()
    file_name = "temp_{}"
    # split input file into several parts
    files = [0]*27
    temp_numbers = []
    print("Start split and sort!")

    # NEW
    with open(os.path.join(ROOT_DIR, 'gen_numbers.txt')) as f_in:
        #f_in.seek(2)
        for i in range(27):
            temp_numbers = f_in.read(12*37037037).split('\n')  # read each 1/27 part
            temp_numbers.sort()  # Timsort in Python
            files[i] = open(os.path.join(ROOT_DIR, file_name.format(i)), 'w+')
            files[i].write(''.join(str(line) + '\n' for line in temp_numbers))
            temp_numbers.clear()
        print("Split and sorting parts time: {} seconds".format((time() - start_split)))

    # external sorting
    buffer = []
    print("Start external sorting")
    start_sort = time()
    f_out = open(os.path.join(ROOT_DIR, 'sort1_numbers.txt'), 'w')  # 3
    files_copy = files.copy()

    for p in range(27):  # 4
        files[p].seek(2)  # cursor in each file to it`s start
        temp_numbers.append(files[p].readline()[:11])
    while True:
        min_i = temp_numbers.index(min(temp_numbers))  # 5.1
        if len(buffer) == 37037037:  # 5.2
            f_out.write(''.join(line for line in buffer)) # 5.2
            buffer.clear()
        else:
            buffer.append(temp_numbers[min_i]+'\n')
        k = files[min_i].readline()  # 5.3
        if k == "" or k == "\n" or k is None:  # 5.4
            temp_numbers.pop(min_i)
            files[min_i].close()
            files.pop(min_i)
        else:
            temp_numbers[min_i] = k[:11]
        if len(files) == 0:
            if len(buffer) > 0:
                f_out.write(''.join(line for line in buffer))
            break
        if len(temp_numbers) == 0:
            if len(buffer) > 0:
                f_out.write(''.join(line for line in buffer))
            break

    for w in range(27):  # 6
        try:
            files_copy[w].close()
        except:
            pass
        finally:
            os.remove(os.path.join(ROOT_DIR, file_name.format(w)))
    f_out.close()
    print("External sorting time: {} seconds".format((time() - start_sort)))


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


