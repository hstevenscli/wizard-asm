# Documentation/Notes on wizard_asm source code

### Notes

return player location whenver they use a move or teleport instr. location is returned when spawning players

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


#### Update 01/14/25

change fireball to magma eruption
let the spell take in two numbers, row and col, to allow for better strategies around that spell
let most spells work in this way, these will mesh better with the divination/find enemy spell

divination returns a row,col somwhere in an area around the enemy with accuracy increasing as more time and mana is spent on the spell


### Checklist Delete when Complete
    6. X Divination (Find where other player is)
    8. Conditional Jump 
    12. Invisibility
    13. Add mana costs to all of the spells

## NEXT ON LIST

Make a JSON object and read from it, turn the json data into a battle program and run it

```
loop_through_battle_program
```
need to change the main loop to a for loop that lasts for 1000 game ticks for example:

loop 1000 {
    p1 program chunk
    p2 program chunk
}
and have two main chunks inside the loop, so every iteration of the loop will call each players program once. each chunk will track where each battle program is and if needed not move on if a sloop or eloop instruction is called. this way each battle program gets an instruction called every game tick


## ASM INSTRUCTIONS

```go
func asm_summon_magma( g *gameSpace, player int, row int, col int ) 
func asm_summon_acid ( g *gameSpace, player int, row int, col int ) 
func asm_move( g *gameSpace, player int, direction string ) 
func asm_teleport( g *gameSpace, player int, row int, col int ) 
func asm_shield( g *gameSpace, player int, d_type int ) 
func asm_wait( player int ) 
func asm_recharge( g *gameSpace, player int, amount int ) 
func asm_divination( g *gameSpace, player int, intensity int ) 
func asm_lightning_bolt( g *gameSpace, player int, direction string
```

## Struct Info


__battleProgram:__
* Lstart -> marking the start of a loop for SLoop instruction
* Niterations -> used by ELoop to track how long the loop has been going


## Mana Costs

summon magma
summon acid 
move
teleport
shield
wait
recharge
divination
lightning bolt
