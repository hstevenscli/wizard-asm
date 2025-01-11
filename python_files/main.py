import random

class GameSpace:
    def __init__(self, n: int):
        self.mSize = n
        self.mGameSpace = []
        self.mPloc_A = (-1, -1)
        self.mPloc_B = (-1, -1)
        self.init_gamespace()

    # Initialize the game space to a nxn grid
    def init_gamespace(self):
        self.mGameSpace = [
        ]
        for i in range(self.mSize):
            self.mGameSpace.append([0 for i in range(self.mSize)])

    # Return the array that contains the arrays of length n.
    def get_game_space(self):
        return self.mGameSpace

    # Randomly place players in the gamespace, return their location
    def spawn_players(self):
        # i = random.randint(0, self.mSize - 1)
        # j = random.randint(0, self.mSize - 1)
        # k = random.randint(0, self.mSize - 1)
        # l = random.randint(0, self.mSize - 1)
        i = 1
        j = 1
        k = 3
        l = 1
        self.mGameSpace[i][j] = "A"
        ploc_a = (i, j)
        while k == i and l == j:
            if k == i:
                k = (k + 1) % self.mSize
            # Never reach the else?
            # else:
            #     l = (l + 1) % self.mSize
            # Could also just redo k and l
            # k = random.randint(0, self.mSize - 1)
            # l = random.randint(0, self.mSize - 1)
        self.mGameSpace[k][l] = "B"
        ploc_b = (k, l)
        self.mPloc_A = ploc_a
        self.mPloc_B = ploc_b

    def instr_shoot_fireball(self, direction, distance, player):
        #direction can be the following:
        # valid_directions = ["n", "w", "s", "e", "nw", "ne", "sw", "se"]
        if player == "a":
            ploc = self.mPloc_A
        else:
            ploc = self.mPloc_B
        p_row = ploc[0]
        p_col = ploc[1]
        print("Player", player, "located at ({},{})".format(p_row, p_col))
        '''
        Notes for me on directions
        n - decrease row
        s - increase row
        e - increse col
        w - decrease col
        ne - decrease row, increase col
        nw - decrease row, decrease col
        se - increase row, increase col
        sw - increase row, decrease col

        row = i
        col = j

        Distance:
        just tells you how many times to increment or dencrement i and/or j
        '''
        # fireball has a radius of  1
        if direction in ["n", "nw", "ne"]:
            # decrease row
            d_row = p_row - distance
            d_col = p_col
            if d_row < 0:
                d_row = 0
            #if w decrease col
            if direction == "nw":
                d_col = p_col - distance
            #if e increase col
            if direction == "ne":
                d_col = p_col + distance

        if direction in ["s", "se", "sw"]:
            # increase row
            d_row = p_row + distance
            d_col = p_col
            if d_row > self.mSize - 1:
                d_row = self.mSize - 1
            #if w decrease col
            if direction == "sw":
                d_col = p_col - distance
            #if e increase col
            if direction == "se":
                d_col = p_col + distance

        if direction == "e":
            d_row = p_row
            d_col = p_col + distance

        if direction == "w":
            d_row = p_row
            d_col = p_col - distance

        # keep d_col within bounds
        if d_col < 0:
            d_col = 0

        if d_col > self.mSize - 1:
            d_col = self.mSize - 1


        #9 locations
        # drow and dcol +- 1 
        # drow -1 and dcol +-1
        # drow +1 and dcol +-1
        PHIT = False

        print("Tile where fireball will land is at ({},{})".format(d_row, d_col))
        for row in range(-1, 2):
            for col in range(-1, 2):
                print("Checking: {},{}".format(d_row + row, d_col + col))
                b_row = d_row + row
                b_col = d_col + col
                b_row = max(b_row, 0)
                b_col = max(b_col, 0)

                if self.mGameSpace[d_row + row][d_col + col] in ("A", "B"):
                    print("Player at {},{} in blast radius".format(d_row + row, d_col + col))
                    print("V at coords:", self.mGameSpace[d_row + row][d_col + col])

        print(self.mGameSpace[d_row][d_col])
        if self.mGameSpace[d_row][d_col] == "A" or self.mGameSpace[d_row][d_col] == "B":
            print("HIT PLAYER")


    # Print the gamespace out by looping through all of the arrays in gamespace
    def print_game_space(self):
        for i in self.mGameSpace:
            for j in i:
                print(j,end="")
            print()

def main():
    gs = GameSpace(4)
    gs.spawn_players()
    for arr in gs.get_game_space():
        print(arr)
    gs.print_game_space()
    gs.instr_shoot_fireball("n", 3, "b")


main()
