# Documentation/Notes on wizard_asm source code

### Notes

make a check instruction, it will check whatever is on the stack. make certain instructions return certain values. eg the move spell returns 0 if the move action can't be completed.

an instruction to find out where the enemy is
an instruction to hide yourself
a decoy instruction 
lightning/beam instruction
a spreading fire instruction


#### Brainstorming

Divination right now returns a random location in the vicinity of the other player. Might change later

## NEXT ON LIST

Put next item here:
- [x] Conditional
- [ ] crystal ball (for storing numbers)
- [x] Divination
- [x] Add mana cost to spells 
- [ ] Invisibility?
- [ ] Decoy Instruction
- [x] Flesh out frontend; Get Login/Register working
- [x] Hook up backend to mongoDB
- [ ] Implement add instruction

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

ADD - adds the given number to the given instruction
```
instr1 0
instr2 0
add -2 1
```
add 1 to instr1 argument, so after running instr1's arg would go from 0->1
subtract does the opposite

since some args have 2 numbers that they take, like magma, need to have a way to hit both of those numbers



## Struct Info


__battleProgram:__
* Lstart -> marking the start of a loop for SLoop instruction
* Niterations -> used by ELoop to track how long the loop has been going


## Mana Costs

| Spell | Cost |
| - | - |
| Magma | 20   |
| Acid   | 10   |
| Move  |  0  |
| Teleport   | 100  |
| Shield   | 10   |
| Wait   |  0  |
| Recharge | 0 | 
| Divination | 150
| Lightning | 50



## Notes on how the user runs battles


method 1: duels

user makes a request giving a username of someone to duel and providing their own username
lookup the two battleprograms run a battle and return it to the user that made the request

Needed implementations
frontend: a button with a text box to put a username into

backend: receive a request like get(/duels)
need a handler that will lookup both bps from the db, load them into memory and call the run duels function



method 2: tourney?

upon submission, run the bp against a few other randomly selected bps and return the results

needed implementations
frontend: make call to backend tourney handler, maybe get(/tournament)
backend: /tournament handler that will query the db for a number of battleprograms from various users and run the submitted program against all of them.
