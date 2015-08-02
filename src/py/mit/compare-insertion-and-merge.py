import math

def insertion_time(n):
    return 8 * n * n

def merge_time(n):
    return 64 * n * math.log(n, 2)

for n in range(2, 100):
    if merge_time(n) <= insertion_time(n):
        print(n)
        break
