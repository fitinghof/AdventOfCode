file = open("input.txt")

lines = file.read().split("\n")

sum = 0
cardCopies = {}
for line in range(len(lines)):
    if not cardCopies.__contains__(line):
        cardCopies[line] = 1
        
    nums = lines[line].replace(", ","").replace("\n","").replace("Card ","").split("|")
    winNums = nums[0].split(" ")
    winNums.pop(0)

    for i in range(len(winNums)):
        if i >= len(winNums): break
        if winNums[i] == " " or winNums[i] == "":
            winNums.pop(i)
            i -= 1

    myNums = nums[1].split(" ")
    for i in range(len(myNums)):
        if i >= len(myNums): break
        if myNums[i] == " " or myNums[i] == "":
            myNums.pop(i)
            i -= 1

    numWins = 0
    for win in winNums:
        for my in myNums:
            if win == my:
                numWins+=1
                print(win, my)
                break
    for i in range(numWins):
        if cardCopies.__contains__(line + i + 1):
            if cardCopies.__contains__(line):
                cardCopies[line + i + 1] += 1*cardCopies[line]
            else: print("ERR")
        else: cardCopies[line + i + 1] = 1 + 1*cardCopies[line]
        
totalCards = 0
print(cardCopies)
for copies in cardCopies.values():
    totalCards += copies
print(totalCards)
        