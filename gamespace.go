package main

import "fmt"
import "math/rand"


// Define gameSpace structure
type gameSpace struct {
    Size int
    Arena [][]int
    P1_loc [2]int
    P2_loc [2]int
    P1_prot int
    P2_prot int
    P1_mana int
    P2_mana int
}

type pInfo struct {
	ploc [2]int
	prot int
	mana int
}

type gameOver struct {
    p1 bool
    p2 bool
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
    g.P1_mana = 100
    g.P2_mana = 100
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

// Make sure the location is valid first
func check_player( r int, c int, g *gameSpace ) (int, bool) {
	var loc int = g.Arena[r][c]
	if loc == 1 || loc == 2 {
		return loc, true
	}
	return 0, false
}

func is_valid_loc( r int, c int, size int)bool {
    if r < 0 || r > size - 1 || c < 0 || c > size - 1 {
        return false
    }
    return true
}

// Remove the player symbol from the arena
func erase_player( g *gameSpace, player int ){
    var p_row int
    var p_col int
    if player == 1 {
        p_row = g.P1_loc[0]
        p_col = g.P1_loc[1]
    } else {
        p_row = g.P2_loc[0]
        p_col = g.P2_loc[1]
    }
    g.Arena[p_row][p_col] = 0
}

// Update the player location data in gamespace
func update_player_location_data( g *gameSpace, player int, row int, col int ) {
    if player == 1 {
        g.P1_loc[0] = row
        g.P1_loc[1] = col
    } else {
        g.P2_loc[0] = row
        g.P2_loc[1] = col
    }
}

// Redraw the player in the arena using location data in gamespace
func draw_player( g *gameSpace, player int ){
    var p_row int
    var p_col int
    if player == 1 {
        p_row = g.P1_loc[0]
        p_col = g.P1_loc[1]
    } else {
        p_row = g.P2_loc[0]
        p_col = g.P2_loc[1]
    }
    g.Arena[p_row][p_col] = player

}


// Packages erase_player, update_player_location_data, and draw_player into one function call
func easy_move_wrapper( g *gameSpace, player int, row int, col int ) {
    erase_player(g, player)
    update_player_location_data(g, player, row, col)
    draw_player(g, player)
}
