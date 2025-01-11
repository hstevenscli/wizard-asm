package main

import "fmt"
import "math/rand"
// import "time"

// Define gameSpace structure
type gameSpace struct {
    Size int
    Arena [][]int
    P1_loc [2]int
    P2_loc [2]int
}

func main() {

    // Initialize size 
    var size int = 10

    // Initialize the gamespace to size = 10
    var g gameSpace = init_gamespace(size)

    fmt.Printf("SIZE: %v\n", size)

    // Spawn players in gameSpace
    spawn_players( &g )

    // Pretty print gameSpace
    pretty_print(g.Arena)
    asm_cast_fireball(g, 1, "nw", 4)
}

// Print out gameSpace in a human-viewable format
func pretty_print( slice [][]int ) {
    for i := range slice {
            fmt.Println(slice[i])
    }
}

// Init and fill the gamespace and gamespace.arena
func init_gamespace( size int ) gameSpace {
    var g gameSpace = gameSpace { Size: size, Arena: make([][]int, size) }
    for i := range g.Arena {
        g.Arena[i] = make([]int, size)
    }
    return g
}

// Spawn some players into the gamespace
func spawn_players( g *gameSpace ) {
    var p1_row int = rand.Intn(g.Size)
    var p1_col int = rand.Intn(g.Size)
    var p2_row int = rand.Intn(g.Size)
    var p2_col int = rand.Intn(g.Size)

    // Place Player 1
    g.Arena[p1_row][p1_col] = 1

    // Prevent deleting player 1
    if p1_row == p2_row && p1_col == p2_col {
        p2_row = (p2_row + 1) % g.Size
    }
    // Place Player 2
    g.Arena[p2_row][p2_col] = 2

    // Record player positions in struct
    g.P1_loc[0] = p1_row
    g.P1_loc[1] = p1_col
    g.P2_loc[0] = p2_row
    g.P2_loc[1] = p2_col

    fmt.Println("Spawned some players")
}

// Shoot a fireball in the indicated direction
func asm_cast_fireball( g gameSpace, player int, direction string, distance int ) {

    // var north_directions = map[string]bool{
    //     "n": true,
    //     "nw": true,
    //     "ne": true,
    // }

    var p_row int
    var p_col int
    if player == 1 {
        p_row = g.P1_loc[0]
        p_col = g.P1_loc[1]
    } else {
        p_row = g.P2_loc[0]
        p_col = g.P2_loc[1]
    }
    fmt.Println(g.P1_loc)
    fmt.Println("Row", p_row)
    fmt.Println("Col", p_col)

    for i, char := range direction {
        fmt.Printf("%v, %c\n", i, char)
        if char == 'n' {
            fmt.Println("N")
        }
        if char == 's' {
            fmt.Println("S")
        }
        if char == 'e' {
            fmt.Println("E")
        }
        if char == 'w' {
            fmt.Println("W")
        }
    }
    // var south_directions = map[string]bool{
    //     "s": true,
    //     "sw": true,
    //     "se": true,
    // }

    // if north_directions[direction] {
    //     fmt.Println("NORTH")
    // } else {
    //     fmt.Println("SOUTH")
    // }

    if direction == "n" || direction == "nw" || direction == "ne" {
        return
    }



}


