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
- [ ] Conditional and crystal ball (for storing numbers)
- [x] Divination
- [x] Add mana cost to spells 
- [ ] Invisibility?
- [ ] Decoy Instruction
- [ ] Flesh out frontend; Get Login/Register working
- [ ] Hook up backend to mongoDB

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
