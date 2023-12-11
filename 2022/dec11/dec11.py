import math

def generateInput():
    file1 = open('dec11input.txt', 'r')
    # file1 = open('sample.in', 'r')
    lines = file1.readlines()

    monkeyLineDetails = []

    commonDivisor = 1

    for line in lines:
        if line.strip() == '':
            monkeyList.append(monkeyLineDetails)
            monkeyLineDetails = []
            continue
        monkeyLineDetails.append(line.strip())
    monkeyList.append(monkeyLineDetails)

    for monkeyLineDetails in monkeyList:
        monkey = {}
        for line in monkeyLineDetails:
            if line.startswith('Monkey '):
                monkey['id'] = line.split()[1].split(':')[0]
                continue
            if line.startswith('Starting items:'):
                monkey['items'] = list(map(lambda x: int(x), line.split(':')[1].split(',')))
                continue
            if line.startswith('Operation'):
                operation = line.split('=')[1].strip()
                monkey['op'] = eval('lambda old: ' + operation)
                continue
            if line.startswith('Test:'):
                monkey['divisibleBy'] = int(line.split()[3])
                commonDivisor *= monkey['divisibleBy']
                continue
            if line.startswith('If true'):
                monkey['divisibleByTrue'] = line.split()[5]
                continue
            if line.startswith('If false'):
                monkey['divisibleByFalse'] = line.split()[5]
                continue

        monkeyBehavior[monkey['id']] = monkey
    return commonDivisor


def monkeySeeMonkeyDo(monkeyBehavior, rounds, monkeyShenaniganCount, worryAdjust, commonDivisor):

    for i in range(rounds):
        for key in monkeyBehavior.keys():
            monkey = monkeyBehavior[key]
            for item in monkey['items']:
                # monkey['items'].pop(0)
                worryLevel = monkey['op'](item)
                # print(item, ':', worryLevel)
                if worryAdjust:
                    worryLevel = math.floor(worryLevel / 3)
                else:
                    # mod by common divisor to keep numbers small
                    worryLevel = worryLevel % commonDivisor

                if worryLevel % monkey['divisibleBy'] == 0:
                    # print('Worrylevel', worryLevel, 'to', monkey['divisibleByTrue'])
                    monkeyBehavior[monkey['divisibleByTrue']]['items'].append(worryLevel)
                else:
                    # print('Worrylevel', worryLevel, 'to', monkey['divisibleByFalse'])
                    monkeyBehavior[monkey['divisibleByFalse']]['items'].append(worryLevel)

                monkeyShenaniganCount[int(key)] += 1
            monkey['items'].clear()
    return

def printMonkeyItems(monkeyBehavior):
    for key in monkeyBehavior.keys():
            monkey = monkeyBehavior[key]
            print(monkey['id'], ': ', monkey['items'])
    return


instructions = []
monkeyList = []
monkeyBehavior = {}
commonDivisor = generateInput()
monkeyShenaniganCount = [0]*len(monkeyBehavior.keys())
printMonkeyItems(monkeyBehavior)


# Part 1
monkeySeeMonkeyDo(monkeyBehavior, 20, monkeyShenaniganCount, True, commonDivisor)
printMonkeyItems(monkeyBehavior)
monkeyShenaniganCount.sort(reverse=True)
print('Part 1:', monkeyShenaniganCount[0] * monkeyShenaniganCount[1])

# Part 2
commonDivisor = 1
instructions = []
monkeyList = []
monkeyBehavior = {}
commonDivisor = generateInput()
monkeyShenaniganCount = [0]*len(monkeyBehavior.keys())

monkeySeeMonkeyDo(monkeyBehavior, 10000, monkeyShenaniganCount, False, commonDivisor)
printMonkeyItems(monkeyBehavior)
monkeyShenaniganCount.sort(reverse=True)
print('Part 2:', monkeyShenaniganCount[0] * monkeyShenaniganCount[1])