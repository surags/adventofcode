inputList = []

def generateList():
    file1 = open('dec4input.txt', 'r')
    # file1 = open('sample.in', 'r')
    lines = file1.readlines()
    for line in lines:
        # inputList.append(line.strip())
        lineEntry = {}
        assignments = line.split(',')
        lineEntry['pair1'] = list(map(lambda x: int(x), assignments[0].split('-')))
        lineEntry['pair2'] = list(map(lambda x: int(x), assignments[1].split('-')))

        inputList.append(lineEntry)

def isSet2inSet1(set1: set, set2: set):
    for entry in set2:
        if entry not in set1:
            return False
    return True

def isOverlapping(set1: set, set2: set):
    return not set1.isdisjoint(set2)

def calculateCompleteOverlap(inputList):
    noCompleteOverlap = 0
    for entry in inputList:
        set1 = set(range(entry['pair1'][0], entry['pair1'][1]+1))
        set2 = set(range(entry['pair2'][0], entry['pair2'][1]+1))
        # print('Set1: ', set1)
        # print('Set2: ', set2)
        if isSet2inSet1(set1, set2) or isSet2inSet1(set2, set1):
            # print('Overlap detected')
            noCompleteOverlap += 1

    print('Complete overlap: ', noCompleteOverlap)

def calculateSomeOverlap(inputList):
    noSomeOverlap = 0
    for entry in inputList:
        set1 = set(range(entry['pair1'][0], entry['pair1'][1]+1))
        set2 = set(range(entry['pair2'][0], entry['pair2'][1]+1))
        # print('Set1: ', set1)
        # print('Set2: ', set2)
        if isOverlapping(set1, set2):
            # print('Overlap detected')
            noSomeOverlap += 1

    print('Some overlap: ', noSomeOverlap)

generateList()
calculateCompleteOverlap(inputList)
calculateSomeOverlap(inputList)