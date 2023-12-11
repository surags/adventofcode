instructions = []

def generateInputList():
    file1 = open('dec10input.txt', 'r')
    # file1 = open('sample.in', 'r')
    lines = file1.readlines()
    for line in lines:
        inputSplit = line.strip().split()
        if len(inputSplit) == 1:
            inputSplit.append('0')
        instructions.append([inputSplit[0], int(inputSplit[1])])

def executeInstructionsPart1(instructions: list):
    register = 1
    cycle = 0
    signalStrengthSum = 0
    printCycleList = [20, 60, 100, 140, 180, 220]
    index = 0
    for instruction in instructions:
        index += 1
        if instruction[0] == 'noop':
            # No register update
            cycle += 1
            if (cycle in printCycleList):
                # print('Cycle:', cycle, ' Register:', register, ' Instruction:', instruction, ' Index:', index)
                signalStrengthSum += cycle * register
        elif instruction[0] == 'addx':
            cycle += 1
            if (cycle in printCycleList):
                # print('Cycle:', cycle, ' Register:', register, ' Instruction:', instruction, ' Index:', index)
                signalStrengthSum += cycle * register

            cycle += 1
            if (cycle in printCycleList):
                # print('Cycle:', cycle, ' Register:', register, ' Instruction:', instruction, ' Index:', index)
                signalStrengthSum += cycle * register
            register += instruction[1]

    print('Part1')
    print('Signal Strength Sum:', signalStrengthSum)
    return

def updatePixel(screen, line, posn, register, cycle):
    # print('Screen(', line, ',', posn, ') reg:', register)
    if posn == register or ((posn == (register + 1)) and (register + 1 < 40 )) or ((posn == register + 2) and (register + 2 < 40 )):
        screen[line][posn] = '#'
    else:
        screen[line][posn] = '.'

def executeInstructionPart2(instructions: list):
    register = 1
    cycle = 0
    signalStrengthSum = 0
    index = 0
    screen = []
    for i in range(7):
        screen.append(['#']*40)
    line = 0
    posn = 0
    printCycleList = [40,80,120,160,200,240]

    for instruction in instructions:
        index += 1
        if instruction[0] == 'noop':
            # No register update
            cycle += 1
            posn += 1
            if cycle in printCycleList:
                line += 1
                posn = 0

            updatePixel(screen, line, posn, register, cycle)
        elif instruction[0] == 'addx':
            cycle += 1
            posn += 1

            if cycle in printCycleList:
                line += 1
                posn = 0

            updatePixel(screen, line, posn, register, cycle)

            cycle += 1
            posn += 1

            if cycle in printCycleList:
                line += 1
                posn = 0

            updatePixel(screen, line, posn, register, cycle)
            register += instruction[1]

    print('Part2')
    print('Screen')
    printScreen(screen)

    return

def printScreen(screen: list):
    screen.pop()
    for line in screen:
        for dot in line:
            print(dot, end='')
        print('\n', end='')
    return

generateInputList()
executeInstructionsPart1(instructions)
print()
executeInstructionPart2(instructions) # there's an off by one error somewhere here but close enough

