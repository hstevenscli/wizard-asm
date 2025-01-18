package main

import "fmt"
// import "time"


func main() {

    // Initialize size 
    var size int = 4

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
    move(&g, 2, "n")
    pretty_print(g.Arena)
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

func game_over(player int) {
	fmt.Printf("Player %v Hit. Player %v Loses\n", player, player)
}
