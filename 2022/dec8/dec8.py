inputGrid = []
visibilityCount = []
scenicScoreGrid = []

def generateInputGrid():
    file1 = open('dec8input.txt', 'r')
    # file1 = open('sample.in', 'r')
    lines = file1.readlines()
    for line in lines:
        # logs.append(line.strip())
        inputGrid.append(list(map(lambda x: int(x), [*line.strip()])))

def initializeVisibilityGrid():
    length = len(inputGrid)
    width = len(inputGrid[0])
    for i in range(length):
        visibilityCount.append([0]*width)

def initializeScenicScoreGrid():
    length = len(inputGrid)
    width = len(inputGrid[0])
    for i in range(length):
        scenicScoreGrid.append([0]*width)

def calculateUpVisibility(length, width):
    # Iterate over i for each j
    for j in range(width):
        maxTreeHeight = 0
        for i in range(length) :
            if inputGrid[i][j] > maxTreeHeight or i == 0:
                maxTreeHeight = inputGrid[i][j]
                visibilityCount[i][j] += 1
    return

def calculateDownVisibility(length, width):
    # Iterate reverse over i for each j
    for j in range(width):
        maxTreeHeight = 0
        for i in range(length-1, -1, -1) :
            if inputGrid[i][j] > maxTreeHeight or i == length-1:
                maxTreeHeight = inputGrid[i][j]
                visibilityCount[i][j] += 1
    return

def calculateLeftVisibility(length, width):
    # Iterate over j for each i
    for i in range(length):
        maxTreeHeight = 0
        for j in range(width):
            if inputGrid[i][j] > maxTreeHeight or j == 0:
                maxTreeHeight = inputGrid[i][j]
                visibilityCount[i][j] += 1
    return

def calculateRightVisibility(length, width):
    # Iterate reverse over j for each i
    for i in range(length):
        maxTreeHeight = 0
        for j in range(width-1, -1, -1):
            if inputGrid[i][j] > maxTreeHeight or j == width-1:
                maxTreeHeight = inputGrid[i][j]
                visibilityCount[i][j] += 1
    return

def calculateVisibility():
    length = len(inputGrid)
    width = len(inputGrid[0])

    calculateUpVisibility(length, width)
    calculateDownVisibility(length, width)
    calculateLeftVisibility(length, width)
    calculateRightVisibility(length, width)
    return

def calculateScenicScore(i, j, length, width):
    left = 0
    right = 0
    up = 0
    down = 0

    # left
    # decrease j till tall tree
    l = j
    maxHeight = 0
    while l > 0:
        l -= 1
        if inputGrid[l][j] > maxHeight:


    scenicScoreGrid[i][j] = left * right * up * down
    return

def calculateScenicScore():
    length = len(inputGrid)
    width = len(inputGrid[0])

    for i in range(length):
        for j in range(width):
            calculateScenicScore(i, j, length, width)
    return

def countNoOfVisibleTrees(visibilityCount):
    count = 0
    for i in range(len(visibilityCount)):
        for j in range(len(visibilityCount[0])):
            if visibilityCount[i][j] > 0:
                count += 1
    print('Part1: ', count)
    return

def printGrid(grid):
    print()
    print('Grid:')
    for gridLine in grid:
        print(gridLine)


# Setup
generateInputGrid()
initializeVisibilityGrid()
initializeScenicScoreGrid()

print(inputGrid)
print(visibilityCount)
print(scenicScoreGrid)

# Part1
calculateVisibility()
printGrid(visibilityCount)
# print(visibilityCount)
countNoOfVisibleTrees(visibilityCount)

# Part2
calculateScenicScore()

