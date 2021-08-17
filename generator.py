from random import randint
import os


ROOT_DIR = os.path.dirname(os.path.abspath('config.py'))


def gen():
    # 8-9**-***-**-**
    numbers = []
    n = int(input("Кол-во номеров, которые нужно сгенерировать: "))
    print()
    file = open(os.path.join(ROOT_DIR, 'gen_numbers.txt'), 'w')
    i = 0

    while i < n:
        temp_num = 0
        for j in range(9):
            temp_num += randint(0, 9) * 10**(8 - j)
        if temp_num not in numbers:
            numbers.append(temp_num)
            i += 1
            str_num = str(temp_num)
            if len(str_num) < 9:
                for i in range(9-len(str_num)):
                    str_num = "0" + str_num
            file.write("89" + str_num + "\n")
