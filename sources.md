# Keep a list of all sources

## ChatGPT
chatgpt helped me make small changes to make code simpler and in some cases possibly more efficient

like this for loop in asm_cast_fireball
```go
for i, char := range direction {
    fmt.Printf("%v, %c\n", i, char)
}
```
just these lines, not the logic inside the for loop

also for quick questions about documentation and finding errors while learning go


q: how can i check if a value is in an array in go?

a: 

q:     if direction == "n" || direction == "nw" || direction == "ne" {
        return
    }

is there a more succint way to write this?

Brainstorming spells.
Question:
```
Go – Ten base instructions finished.
    1. Fireball (AoE)
    2. Lightning Bolt (Line Attack)
    3. Move
    4. Teleport (Random Location)
    5. Shield (Protection From Damage)
    6. Divination (Find where other player is)
    7. Time Jump (Jump asm)
    8. Conditional Jump 
    9. Acid Puddle (Lingering Damage on the ground)
    10. Splash (Gets Player wet)

	 Able to run a (singular) ‘battle program’ in a simple, static game environment

for a simple game where two wizards must battle by writing a simple assembly like program i have these ten instructions. time jump and conditional jump work as a sort of loop (time jump) and if clause (conditional jump) time jump just jumps up or down the program the specified amount.

what are some other instructions I could include and implement to allow for more sophisticated strategies in this way? 

also some other 'spell' instructions like the other attack instructions

```

Answer:
```
Here are some suggestions for additional instructions you could include, divided into utility/control-flow, buffs/debuffs, and attack/damage-related spells. These instructions can enhance strategy and creativity for players in your game.

Utility/Control-Flow Instructions
Wait (NOP - No Operation):

Does nothing for a turn, allowing the player to stall or time their actions more precisely.
Could synergize with delayed effects or counterplay.
Counterspell:

Negates the next spell cast by the opponent. Adds a layer of prediction and strategy.
Summon Familiar:

Summons an ally that performs basic actions (e.g., attacks, shields, or distracts).
Delay:

Executes the next instruction after a specified number of turns, enabling delayed spells or effects.
Loop (Mark/EndLoop):

Creates a defined loop in the program, allowing for repetitive behaviors.
Memory/Variable Store (Set/Get):

Store and retrieve data in a small memory location. For example:
Set A 10 (stores 10 in memory location A).
Jump If A > 0.
Great for adding tactical complexity and state management.
Buffs/Debuffs
Haste:

Increases the number of instructions executed per turn, allowing faster action execution.
Slow:

Reduces the number of instructions the opponent executes per turn.
Weaken:

Reduces the power or effectiveness of the opponent's next spell.
Confusion:

Randomizes the opponent’s next set of instructions.
Gravity Well:

Slows down movement and teleportation in a specific area.
Camouflage:

Makes the player invisible for a number of turns, preventing them from being targeted directly.
Attack/Damage-Related Spells
Ice Spike (Line Attack + Freeze):
A straight-line attack like Lightning Bolt, but it also freezes the target for a turn, preventing movement.
Meteor (Delayed AoE):
Targets an area, but lands after 1-2 turns, requiring prediction to hit.
Chain Lightning:
Jumps between multiple targets if they're within a certain range of each other.
Poison Cloud:
Creates a lingering AoE that deals small amounts of damage over several turns.
Vortex:
Pulls all entities in an area toward a specific point, disrupting movement.
Thunderclap (AoE + Stun):
Deals damage in an area and stuns affected players for a turn.
Eruption:
Creates random small AoE explosions over several turns in a defined area.
Leech Life:
Steals health from the opponent and heals the caster.
Defensive/Healing Spells
Barrier:
Creates a barrier in a line or arc that blocks attacks.
Regenerate:
Heals the player over time, restoring health each turn.
Mirror Shield:
Reflects the next spell back at the caster.
Absorb Energy:
Converts a portion of incoming spell damage into health or mana.
Area Manipulation/Environmental Effects
Fog:
Creates an area where targeting is random or impossible.
Wall of Fire/Ice:
Creates a barrier that deals damage to anyone crossing it.
Quicksand:
Creates an area that immobilizes any player trying to move through it.
Gust:
Pushes players or objects away from the caster.
Earthquake:
Damages all players on the ground and destroys environmental features.
Advanced Mechanics
Teleport Target:
Moves the opponent to a random location, displacing their positioning.
Mana Drain:
Reduces the opponent’s available energy or ability to cast spells.
Program Swap:
Temporarily swaps a portion of your program with the opponent’s.
Clone:
Creates a duplicate of the player that mimics some instructions but disappears after taking damage.
Amplify:
Increases the power or area of effect of the next spell.
Phase Shift:
Makes the player untargetable for a turn, but they also cannot attack.
Summon Object:
Creates a destructible object in the environment that blocks movement or attacks.
Suggestions for Sophistication
Dynamic Environments:

Instructions that interact with the environment (e.g., "Set Fire to Grass," "Freeze Water").
Combos:

Spells that synergize (e.g., Lightning Bolt does extra damage to wet players from Splash).
Resource Management:

Introduce mana or cooldown systems for certain instructions to add a resource-management layer.
These instructions can open up more creative strategies and make the programming aspect more engaging while adding depth to gameplay. Let me know if you'd like implementation ideas for specific ones!
```













## Go Tutorial by Alex Mux
[Learn GO Fast: Full Tutorial](https://www.youtube.com/watch?v=8uiZC0l4Ajw&t=1624s)

## Go Maps in Action

[From go.dev](https://go.dev/blog/maps)
