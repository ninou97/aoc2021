# https://adventofcode.com/2021/day/4


boards = []
boards_bingos = {}

with open ("input.txt", "r") as f:
    bingo = f.readline()
    bingo = [int(x) for x in bingo.strip().split(",")]
    count = 0
    board = []
    for line in f:
        if count % 5 == 0 and count != 0:
            boards.append(board)
            board = []
            count = 0
        if line != "\n":
            row = line.split()
            row = [int(x) for x in row]

            board.append(row)
            count += 1
    boards.append(board) # because file can end before being able to append last board

for i, board in enumerate(boards):
    boards_bingos[i] = set()


def check_bingo():
    for bingo_num in bingo:
        for i, board in enumerate(boards):
            for y, row in enumerate(board):
                for x, num in enumerate(row):
                    if num == bingo_num:
                        boards_bingos[i].add((x, y))
                        setX = boards_bingos[i]
                        # check if you have a bingo horizontally and vertically
                        c = 0
                        oldX = x
                        oldY = y
                        for x in range(5):
                            if (x, y) in setX:
                                c += 1
                        if c >= 5: yield (i, num)
                        c = 0
                        for y in range(5):
                            if (oldX, y) in setX:
                                c += 1
                        if c >= 5: yield (i, num)

yielded = []
won = set()

generator = check_bingo()
i, bingo_num = next(generator)
yielded.append((i, bingo_num))
won.add(i)

sum = 0
for y, row in enumerate(boards[i]):
    for x, num in enumerate(row):
        if (x, y) not in boards_bingos[i]:
            sum += boards[i][y][x]
answer = sum * bingo_num
print("part 1 answer:", answer)


# part 2

for val in generator:
    if val[0] in won: continue # against already won multiple times
    yielded.append(val)
    if len(yielded) == len(boards): break
    won.add(val[0])

i, bingo_num = yielded[-1]
sum = 0
for y, row in enumerate(boards[i]):
    for x, num in enumerate(row):
        if (x, y) not in boards_bingos[i]:
            sum += boards[i][y][x]
answer = sum * bingo_num
print("part 2 answer:", answer)


