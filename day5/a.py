import math
from collections import deque
from pprint import pprint

file_path = "input.txt"
with open(file_path, 'r') as f:
    data = f.read()

stacklist = []

for line in data.splitlines():
    if not stacklist:
        x = range(0, math.ceil(len(line)/4))
        for n in x:
            stacklist.append(deque())

    if "]" in line:
        sep = list(line)
        x = range(1, len(sep), 4)
        for n in x:
            if sep[n] not in " ":
                stacklist[int((n)/4)].appendleft(sep[n])

    elif "move" in line:
        parts = line.split(' ')
        amount = int(parts[1])
        orig = int(parts[3])
        dest = int(parts[5])
        for i in range(amount):
            stacklist[dest-1].append(stacklist[orig-1].pop())

res = ""
for s in stacklist:
    res += s.pop()
print(res)
