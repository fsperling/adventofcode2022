import os
import sys
import pygtrie

file_path = "input.txt"
with open(file_path, 'r') as f:
    data = f.read()

currentpath = "/"
paths = pygtrie.StringTrie(separator=os.path.sep)
dirs = []

for line in data.splitlines():
    if line.startswith("$"):
        if "cd" == line.split(' ')[1]:
            if ".." == line.split(' ')[2]:
                head, tail = os.path.split(currentpath)
                currentpath = head
            else:
                currentpath = os.path.join(currentpath, line.split(' ')[2])
    elif line.startswith("dir"):
        dirs.append(os.path.join(currentpath, line.split(' ')[1]))
    else:
        paths[os.path.join(currentpath, line.split(' ')[1])] = int(line.split(' ')[0])

totalsize = sum(paths.itervalues())
freespace = 70000000 - totalsize
minimumrequired = 30000000 - freespace

difference = sys.maxsize
dirsize_smallestdiff = 0

for dir in dirs:
    dirsize = sum(paths.itervalues(prefix=dir))
    if dirsize >= minimumrequired:
        diff = dirsize - minimumrequired
        if diff < difference:
            difference = diff
            dirsize_smallestdiff = dirsize

print(dirsize_smallestdiff)
