instructionList = []
crateState = {}

def parseInstructions(instructionListLines: list):
    for line in instructionListLines:
        tokens = line.split()
        instruction = {}
        instruction['action'] = 'move'
        instruction['count'] = int(tokens[1])
        instruction['from'] = tokens[3]
        instruction['to'] = tokens[5]
        instructionList.append(instruction)
    print(instructionList)
    return

def parseCrateState(crateStateLines: list):
    cratesLine = crateStateLines.pop()
    crates = cratesLine.split()
    for crate in crates:
        crateState[crate] = {}
        crateState[crate]['index'] = cratesLine.index(crate)
        crateState[crate]['state'] = []

    while len(crateStateLines) > 0:
        cratesStateLine = crateStateLines.pop()
        for crate in crateState.keys():
            if len(cratesStateLine) > crateState[crate]['index']:
                if cratesStateLine[crateState[crate]['index']] != ' ':
                    crateState[crate]['state'].append(cratesStateLine[crateState[crate]['index']])

    print(crateState)
    return

def parseInputs():
    file1 = open('dec5input.txt', 'r')
    # file1 = open('sample.in', 'r')
    lines = file1.readlines()
    crateStateLines = []
    instructionListLines = []
    for line in lines:
        if(line.startswith('move')):
            instructionListLines.append(line.strip())
        else:
            if len(line.strip()) == 0:
                continue
            else:
                crateStateLines.append(line)

    parseInstructions(instructionListLines)
    parseCrateState(crateStateLines)

def determineTopOfCrates(crateState):
    topOfCrates = ''
    for crate in crateState.keys():
        topOfCrates += crateState[crate]['state'][-1]
    print('Top Of Crates', topOfCrates)
    return

def performOperation(crateState, count, before, after, maintainOrder=False):
    if maintainOrder:
        reorderList=[]
        for iter in range(count):
            reorderList.append(crateState[before]['state'].pop())
        while len(reorderList) > 0:
            crateState[after]['state'].append(reorderList.pop())
    else:
        for iter in range(count):
            crate = crateState[before]['state'].pop()
            crateState[after]['state'].append(crate)
    return

def performOperationsPart1(crateState, instructionList):
    for instruction in instructionList:
        performOperation(crateState, instruction['count'], instruction['from'], instruction['to'])
    return

def performOperationsPart2(crateState, instructionList):
    for instruction in instructionList:
        performOperation(crateState, instruction['count'], instruction['from'], instruction['to'], maintainOrder=True)
    return

parseInputs()
# performOperationsPart1(crateState, instructionList)
# determineTopOfCrates(crateState)
performOperationsPart2(crateState, instructionList)
determineTopOfCrates(crateState)
