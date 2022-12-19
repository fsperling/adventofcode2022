file_path = "input.txt"
with open(file_path, 'r') as f:
    data = f.read()

line = ''.join(data.splitlines())
i = 0
while (True):
    window = line[i:i+4]
    if len(window) == len(set(window)):
        print(i+4)
        break

    i += 1
