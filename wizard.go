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
    // pretty_print(g.Arena)
    pretty_print(g.Arena)
    asm_cast_fireball(&g, 2, "se", 8)
    pretty_print(g.Arena)
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
func asm_cast_fireball( g *gameSpace, player int, direction string, distance int ) {

    var p_row int
    var p_col int
    if player == 1 {
        p_row = g.P1_loc[0]
        p_col = g.P1_loc[1]
    } else {
        p_row = g.P2_loc[0]
        p_col = g.P2_loc[1]
    }
    var d_row int = p_row
    var d_col int = p_col
    fmt.Println("Testing for player", player)
    fmt.Println("PROW:", d_row)
    fmt.Println("PCOL:", d_col)


    if direction == "n" {
        for i := 0; i < distance; i++ {
            d_row -= 1
            d_row = nlt_zero( d_row )
            if d_row <= 0 {
                fmt.Println("Reached edge, Break")
                break
            }
        }
    }

    if direction == "s" {
        for i := 0; i < distance; i++ {
            d_row += 1
            d_row = ngt_size( d_row, g.Size )
            if d_row >= g.Size - 1 {
                fmt.Println("Reached edge, Break")
                break
            }
        }
    }

    if direction == "e" {
        for i := 0; i < distance; i++ {
            d_col += 1
            d_col = ngt_size( d_col, g.Size )
            if d_col >= g.Size - 1 {
                fmt.Println("Reached edge, Break")
                break
            }
        }
    }

    if direction == "w" {
        for i := 0; i < distance; i++ {
            d_col -= 1
            d_col = nlt_zero( d_col )
            if d_col <= 0 {
                fmt.Println("Reached edge, Break")
                break
            }
        }
    }

    if direction == "nw" {
        for i := 0; i < distance; i++ {
            d_row -= 1
            d_col -= 1
            d_row = nlt_zero( d_row )
            d_col = nlt_zero( d_col )
            if d_col <= 0 || d_row <= 0 {
                fmt.Println("Reached edge, Break")
                break
            }
        }

    }

    if direction == "ne" {
        for i := 0; i < distance; i++ {
            d_row -= 1
            d_col += 1
            d_row = nlt_zero( d_row )
            d_col = ngt_size( d_col, g.Size )
            if d_col >= g.Size - 1 || d_row <= 0 {
                fmt.Println("Reached edge, Break")
                break
            }
        }
    }

    if direction == "sw" {
        for i := 0; i < distance; i++ {
            d_row += 1
            d_col -= 1
            d_row = ngt_size( d_row, g.Size )
            d_col = nlt_zero( d_col )
            if d_col <= 0 || d_row >= g.Size - 1 {
                fmt.Println("Reached edge, Break")
                break
            }
        }
    }

    if direction == "se" {
        for i := 0; i < distance; i++ {
            d_row += 1
            d_col += 1
            d_row = ngt_size( d_row, g.Size )
            d_col = ngt_size( d_col, g.Size )
            if d_col >= g.Size - 1 || d_row >= g.Size - 1 {
                fmt.Println("Reached edge, Break")
                break
            }
        }
    }

    fmt.Println("drow:", d_row)

    // d_row = nlt_zero( d_row )
    // d_col = nlt_zero( d_col )
    // d_row = ngt_size( d_row, g.Size )
    // d_col = ngt_size( d_col, g.Size )
    // fmt.Println("After translations:", d_row, "", d_col)
    g.Arena[d_row][d_col] = 3

}


