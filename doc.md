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


### Notes

make a teleport instruction. moves the player to a random location, high mana cost

make the mana Regen take an input but the player will keep doing that same thing depending on how high the number is. eg every 5 mana Regen takes one game cycle.

make fireball a high cost spell.

make a check instruction, it will check whatever is on the stack. make certain instructions return certain values. eg the move spell returns 0 if the move action can't be completed.

an instruction to find out where the enemy is
an instruction to hide yourself
a shield instruction 
a decoy instruction 
lightning/beam instruction
a spreading fire instruction