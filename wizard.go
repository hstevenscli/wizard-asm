package main

import "fmt"


func main() {

    // Initialize size 
    var size int = 3

    // Initialize the gamespace to size = 10
    var g gameSpace = init_gamespace(size)

    fmt.Printf("SIZE: %v\n", size)

    // Spawn players in gameSpace
    spawn_players( &g )

    // Pretty print gameSpace
    // pretty_print(g.Arena)
    pretty_print(g.Arena)
    // asm_summon_magma(&g, 2, 1, 1)
    fmt.Println()
	// asm_summon_acid(&g, 2, 0, 0)
    // asm_shield(&g, 2, 1)
    // asm_lightning_bolt(&g, 1, "s")
    asm_divination(&g, 1, 10)
    pretty_print(g.Arena)

    program := read_json("battleprogram.json")
    fmt.Println("program:", program)



}

// Make sure n is between 0 inclusive and size exclusive
func within_valid_range( n int, size int ) int {
    if n < 0 {
        return 0
    }
    if n >= size {
        return size - 1
    }
    return n
}

func game_over(player int, message string) {
	fmt.Printf("Player %v Hit. Player %v Loses\n", player, player)
    fmt.Println("Message:", message)
    gameover = true
}
