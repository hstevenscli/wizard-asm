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
    // Array is 3 long, players come in values 1 & 2, so [0] should
    // stay empty and not accessed, this makes for easier accessing
    // of player array
    // now that im typing this maybe i should just make a map to store
    // this info
	Pinfo map[int]*pInfo
}

type pInfo struct {
	Row int
	Col int
	Prot int
	Mana int
}

type gameOver struct {
	// Player = []
	Player [3]bool
	Message [3]string
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
	// Init Map
	g.Pinfo = make(map[int]*pInfo)
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

	// Init Pinfo structs and assign row, col info
	p1 := pInfo{}
	p2 := pInfo{}
	g.Pinfo[1] = &p1
	g.Pinfo[2] = &p2
	g.Pinfo[1].Row = p1_row
	g.Pinfo[1].Col = p1_col
	g.Pinfo[2].Row = p2_row
	g.Pinfo[2].Col = p2_col

	// Give players mana
    g.P1_mana = 100
    g.P2_mana = 100
	g.Pinfo[1].Mana, g.Pinfo[2].Mana = 100, 100
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
	p_row = g.Pinfo[player].Row
	p_col = g.Pinfo[player].Col
    g.Arena[p_row][p_col] = 0
}

// Update the player location data in gamespace
func update_player_location_data( g *gameSpace, player int, row int, col int ) {
	g.Pinfo[player].Row = row
	g.Pinfo[player].Col = col
}

// Redraw the player in the arena using location data in gamespace
func draw_player( g *gameSpace, player int ){
    var p_row int
    var p_col int
	p_row = g.Pinfo[player].Row
	p_col = g.Pinfo[player].Col
    g.Arena[p_row][p_col] = player

}


// Packages erase_player, update_player_location_data, and draw_player into one function call
func easy_move_wrapper( g *gameSpace, player int, row int, col int ) {
    erase_player(g, player)
    update_player_location_data(g, player, row, col)
    draw_player(g, player)
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
	fmt.Printf("Player %v Died. Player %v Loses\n", player, player)
    fmt.Println("Message:", message)
    gameover.Player[player] = true
	gameover.Message[player] = message
}
