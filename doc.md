# Documentation/Notes on wizard_asm source code


## Class GameSpace(n)

Params:
* n: size of grid to make

Member Variables:
* mSize - holds the value of n
* mGameSpace - An array of arrays of length n that make up the game space

### get_game_space()

### print_game_space()

Print the gamespace out by looping through all of the arrays in gamespace
Example below
 ```
[
['0', '0'],
['0', '0'],
]

GameSpace array of arrays of length n above would be printed out like so:
00
00

 ```
