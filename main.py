import random

class GameSpace:
    def __init__(self, n: int):
        self.mSize = n
        self.mGameSpace = []
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

    # Randomly place players in the gamespace
    def spawn_players(self):
        i = random.randint(0, self.mSize - 1)
        j = random.randint(0, self.mSize - 1)
        k = random.randint(0, self.mSize - 1)
        l = random.randint(0, self.mSize - 1)
        self.mGameSpace[i][j] = "A"
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

    def instr_shoot_fireball(self, direction, distance, player):
        #direction can be the following:
        # valid_directions = ["n", "w", "s", "e", "nw", "ne", "sw", "se"]
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





        pass



    # Print the gamespace out by looping through all of the arrays in gamespace
    def print_game_space(self):
        for i in self.mGameSpace:
            for j in i:
                print(j,end="")
            print()

def main():
    gs = GameSpace(16)
    gs.spawn_players()
    for arr in gs.get_game_space():
        print(arr)
    gs.print_game_space()


main()
