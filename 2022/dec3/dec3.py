import math

inputList = []

def generateList():
    file1 = open('dec3input.txt', 'r')
    # file1 = open('sample.in', 'r')
    lines = file1.readlines()
    for line in lines:
        inputList.append(line.strip())

def calculatePriority(letter):
    if letter.isupper():
            return ord(letter) - ord('A') + 27
    else:
        return ord(letter) - ord('a') + 1

def calculateSumPriority(inputList):
    sumPriority = 0
    for input in inputList:
        compartment1 = input[0:math.floor(len(input)/2)]
        compartment2 = input[math.floor(len(input)/2):len(input)]
        letter = '?'
        for i in compartment1:
            if i in compartment2:
                letter = i
                break

        sumPriority += calculatePriority(letter)

    print('Sum: ', sumPriority)

def calculateBadgePriority(inputList):
    sumPriority = 0
    group = []
    for input in inputList:
        group.append(input)
        if(len(group) == 3):
            a = set([*group[0]]).intersection(set([*group[1]]))
            letter = a.intersection(set([*group[2]])).pop()
            sumPriority += calculatePriority(letter)
            group = []

    print('Sum: ', sumPriority)


generateList()
calculateSumPriority(inputList)
calculateBadgePriority(inputList)
