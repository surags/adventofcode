
inputList = []

resMap = {
    'A': 'B',
    'B': 'C',
    'C': 'A',
    'CA': 6,
    'CB': 0,
    'CC': 3,
    'AA': 3,
    'AB': 6,
    'AC': 0,
    'BA': 0,
    'BB': 3,
    'BC': 6
}

def generateList():
    file1 = open('dec2input.txt', 'r')
    # file1 = open('sample.in', 'r' )
    lines = file1.readlines()
    for line in lines:
        inputs = line.strip().split(' ')
        inputsobj = {
            'action': inputs[0],
            'response': inputs[1]
        }
        inputList.append(inputsobj)

def convertResponse(response):
    if response == 'Y':
        return 'B'
    if response == 'Z':
        return 'C'
    if response == 'X':
        return 'A'

def convertResponseScore(response):
    if response == 'B':
        return 2
    if response == 'C':
        return 3
    if response == 'A':
        return 1

def determineGamePoints(action, response):
    return resMap[action + response]


def calculatePointsPart1(inputList):
    totalScore = 0
    for input in inputList:
        roundScore = 0
        roundScore += determineGamePoints(input['action'], convertResponse(input['response']))
        roundScore += convertResponseScore(convertResponse(input['response']))
        # print(determineGamePoints(input['action'], input['response']))
        # print(convertResponseScore(input['response']))
        # print('Roundscore ', roundScore)
        totalScore += roundScore

    print('Part1: ', totalScore)

def determineGameResponse(action, req):
    if req == 'X':
        # Lose
        options = ['A', 'B', 'C']
        options.remove(action)
        options.remove(resMap[action])
        return options[0]
    if req == 'Y':
        # draw
        return action
    if req == 'Z':
        # Win
        return resMap[action]

def calculatePointsPart2(inputList):
    totalScore = 0
    for input in inputList:
        roundScore = 0
        roundScore += determineGamePoints(input['action'], determineGameResponse(input['action'], input['response']))
        roundScore += convertResponseScore(determineGameResponse(input['action'], input['response']))
        # print(determineGamePoints(input['action'], input['response']))
        # print(convertResponseScore(input['response']))
        # print('Roundscore ', roundScore)
        totalScore += roundScore

    print('Part2: ', totalScore)


generateList()
calculatePointsPart1(inputList)
calculatePointsPart2(inputList)

