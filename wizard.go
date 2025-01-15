package main

import "fmt"
// import "time"


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
    asm_summon_magma(&g, 2, 1, 1)
    pretty_print(g.Arena)
}

// No less than zero, return zero if n is less than zero
func nlt_zero( n int ) int {
    if n < 0 {
        return 0
    }
    return n
}

// No greater than size, return size if n is greater than size
func ngt_size( n int, size int ) int {
    if n >= size {
        return size - 1
    }
    return n
}

func game_over(player int) {
	fmt.Printf("Player %v Hit. Player %v Loses\n", player, player)
}
