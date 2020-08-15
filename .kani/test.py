import sys

action = sys.argv[1]

if action == "s":
    print('statusが表示されています')

if action == "j":
    f = open('test.java')
    data = f.read()
    f.close()

    print(type(data))
    lines = data.split('\n')
    print(type(lines))
    for line in lines:
        print(line)
