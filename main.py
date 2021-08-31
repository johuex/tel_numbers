from generator import gen
from config import ROOT_DIR
import os
from sorter import sorter_1

if os.path.exists(os.path.join(ROOT_DIR, 'gen_numbers.txt')):
    print("Numbers have been already generated!")
    print()
else:
    print("Numbers are generating!")
    gen()

sorter_1()
# sorter_2()

