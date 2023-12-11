inputMotions = []

def getTstateSetKey(T_state: dict):
    return 't_x_' + str(T_state['x']) + '_y_' + str(T_state['y'])

def generateInputList():
    file1 = open('dec9input.txt', 'r')
    # file1 = open('sample.in', 'r')
    # file1 = open('sample2.in', 'r')
    lines = file1.readlines()
    for line in lines:
        splitInput = line.strip().split()
        inputMotions.append([splitInput[0], int(splitInput[1])])

def updateTMotion(H_state: dict, T_state: dict, t_uniqueStates: set):
    # should T move
    x_diff = abs(H_state['x'] - T_state['x'])
    y_diff = abs(H_state['y'] - T_state['y'])

    if not (x_diff <= 1 and y_diff <= 1):
        # need to move T
        if x_diff == 0:
            # move up/down
            if H_state['y'] > T_state['y']:
                T_state['y'] += 1
            else:
                T_state['y'] -= 1
        elif y_diff == 0:
            # move left/right
            if H_state['x'] > T_state['x']:
                T_state['x'] += 1
            else:
                T_state['x'] -= 1
        else:
            # move diagonally
            if H_state['x'] > T_state['x']:
                T_state['x'] += 1
            else:
                T_state['x'] -= 1

            if H_state['y'] > T_state['y']:
                T_state['y'] += 1
            else:
                T_state['y'] -= 1

    t_uniqueStates.add(getTstateSetKey(T_state))
    return

def simulateHMotion(x_motion: int, y_motion: int, H_state: dict, T_states: list, uniqueStates: list):
    if (x_motion != 0):
        # Moving left or right
        for i in range(abs(x_motion)):
            if x_motion < 0:
                H_state['x'] -= 1
            else:
                H_state['x'] += 1
            if len(T_states) == 1:
                updateTMotion(H_state, T_states[0], uniqueStates[0])
            else:
                updateTMotion(H_state, T_states[0], uniqueStates[0])
                for i in range(len(T_states) - 1):
                    updateTMotion(T_states[i], T_states[i+1], uniqueStates[i+1])


    else:
        # Moving up or down
        for i in range(abs(y_motion)):
            if y_motion < 0:
                H_state['y'] -= 1
            else:
                H_state['y'] += 1
            if len(T_states) == 1:
                updateTMotion(H_state, T_states[0], uniqueStates[0])
            else:
                updateTMotion(H_state, T_states[0], uniqueStates[0])
                for i in range(len(T_states) - 1):
                    updateTMotion(T_states[i], T_states[i+1], uniqueStates[i+1])

    return

def simulateMotions():
    H_state = {
        'x': 0,
        'y': 0
    }

    T_state = {
        'x': 0,
        'y': 0
    }

    t_uniqueStates = set()
    t_uniqueStates.add(getTstateSetKey(T_state))
    for motion in inputMotions:
        x_motion = 0
        y_motion = 0

        # L -> {-1 x}; R -> {+1 x}; U -> {+1 y}; D -> {-1 y}
        if motion[0] == 'L':
            x_motion = 0 - motion[1]
        elif motion[0] == 'R':
            x_motion = motion[1]
        elif motion[0] == 'U':
            y_motion = motion[1]
        else:
            y_motion = 0 - motion[1]

        simulateHMotion(x_motion, y_motion, H_state, [T_state], [t_uniqueStates])

    print('Part1: ', len(t_uniqueStates))
    return

def simulate10DegreeMotions():
    degree_states = []
    uniqueStates = []
    for i in range(10):
        degree_states.append({'x': 0, 'y': 0})
        uniqueStates.append(set([getTstateSetKey(degree_states[i])]))

    for motion in inputMotions:
        x_motion = 0
        y_motion = 0

        # L -> {-1 x}; R -> {+1 x}; U -> {+1 y}; D -> {-1 y}
        if motion[0] == 'L':
            x_motion = 0 - motion[1]
        elif motion[0] == 'R':
            x_motion = motion[1]
        elif motion[0] == 'U':
            y_motion = motion[1]
        else:
            y_motion = 0 - motion[1]

        simulateHMotion(x_motion, y_motion, degree_states[0], degree_states[1:], uniqueStates[1:])

    print('Part2: ', len(uniqueStates[9]))

    return


generateInputList()
simulateMotions()
simulate10DegreeMotions()
