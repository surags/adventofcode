
logs = []
fileSystem = {
    '__size': 0,
    '/': {
        '__size': 0
    }
}

cwd = fileSystem

def generateLog():
    file1 = open('dec7input.txt', 'r')
    # file1 = open('sample.in', 'r')
    lines = file1.readlines()
    for line in lines:
        logs.append(line.strip())

def parseCd(log: str, cwd):
    tokens = log.split()
    if tokens[2] not in cwd.keys():
        cwd[tokens[2]] = {
            '__size': 0,
            '..': cwd
        }
        cwd = cwd[tokens[2]]
    else:
        cwd = cwd[tokens[2]]

    return cwd

def parseLs(log: str, iter: int, cwd):
    lsOutputs = []
    iter += 1
    logLine = logs[iter]
    lsOutputs.append(logLine)
    while (iter + 1 < len(logs)) and (not logs[iter+1].startswith('$')):
        iter += 1
        logLine = logs[iter]
        lsOutputs.append(logLine)

    # print('lsoutput:')
    # print(lsOutputs)

    for lsOutput in lsOutputs:
        tokens = lsOutput.split()
        if tokens[0] == 'dir':
            # dir
            continue
        else:
            cwd[tokens[1]] = int(tokens[0])
            cwd['__size'] += int(tokens[0])

    return cwd

def parseLog(cwd):
    for iter in range(len(logs)):
        log = logs[iter]
        if log.startswith('$'):
            command = ''
            if 'cd' in log:
                command = 'cd'
            else:
                command = 'ls'

            if command == 'cd':
                # print('Parse cd')
                cwd = parseCd(log, cwd)
            else: # ls
                # print('Parse ls')
                cwd = parseLs(log, iter, cwd)

def updateSizes(cwd: dict):
    cwd['__size'] = 0
    for key in cwd.keys():
        if key == '..':
            # skip go back
            cwd['__size'] += 0
            continue

        if type(cwd[key]) == int:
            # file update size
            cwd['__size'] += cwd[key]
        else:
            # dir
            newCwd = cwd[key]
            updateSizes(newCwd)
            cwd['__size'] += cwd[key]['__size']

    if cwd['__size'] < 100000:
        count['dirUnder100000Size'] += cwd['__size']
        count['dirUnder100000'] += 1

    return

def findSmallestDir(cwd:dict, minSize: int):
    if cwd['__size'] >= minSize:
        bestSmallSize = cwd['__size']
    else:
        bestSmallSize = -1
        return -1

    for key in cwd.keys():
        if key == '..':
            # skip go back
            continue

        if type(cwd[key]) == int:
            # skip file
            continue
        else:
            # dir
            newCwd = cwd[key]
            smallSize = findSmallestDir(newCwd, minSize)
            if smallSize == -1:
                # not a viable directory
                continue
            if bestSmallSize == -1:
                bestSmallSize = smallSize
            else:
                if smallSize < bestSmallSize:
                    bestSmallSize = smallSize

    return bestSmallSize


generateLog()
cwd = fileSystem
parseLog(cwd)

count = {'dirUnder100000': 0, 'dirUnder100000Size': 0}
cwd = fileSystem
updateSizes(cwd)

print('Part1: ', count['dirUnder100000Size'])

cwd = fileSystem
minSize = 30000000 - (70000000 - fileSystem['__size'])
print('Part2: ', findSmallestDir(cwd, minSize))
