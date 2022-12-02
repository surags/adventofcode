inputList = []

def generateList():
    file1 = open('dec1input.txt', 'r')
    lines = file1.readlines()
    for line in lines:
        inputList.append(line.strip())

def getMostElf(input: list):
    elf_list = []
    elf_index = 0
    elf_list.append(0)
    for entry in input:
        if entry == '\n' or entry == '':
            elf_index += 1
            elf_list.append(0)
            continue
        elf_list[elf_index] = elf_list[elf_index] + int(entry)

    elf_list.sort(reverse=True)
    print('Part1: ' + str(elf_list[0]))
    print('Part2: ' + str(elf_list[0] + elf_list[1] + elf_list[2]))
    return elf_list[0]

generateList()
getMostElf(inputList)