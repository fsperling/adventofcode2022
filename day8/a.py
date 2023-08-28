import os
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

smalldirs = 0
for dir in dirs:
    dirsize = sum(paths.itervalues(prefix=dir))
    if dirsize <= 100000:
        smalldirs += dirsize

print(smalldirs)
