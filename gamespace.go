package main

import "fmt"
import "math/rand"


// Define gameSpace structure
type gameSpace struct {
    Size int
    Arena [][]int
	Pinfo map[int]*pInfo
	Gameover *gameOver
}

type pInfo struct {
	Row int
	Col int
	Prot int
	Mana int
	CrystalBall [2]int
}

type gameOver struct {
	// Player = []
	Player [3]bool
	Message [3]string
	Conclusion string
	Winner string
	Loser string
	DeathMsg string
}

// Reduce players mana, player dies if mana goes below zero
func deplete_mana( g *gameSpace, player int, mana_cost int ) {
    var manaptr *int = &g.Pinfo[player].Mana
    var new_amount int = *manaptr - mana_cost
    if new_amount < 0 {
        game_over(g, player, "Died of mana depletion")
    } else {
        *manaptr = new_amount
    }
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
    // Init gameover
	gameover := gameOver{}
	g.Gameover = &gameover
	spawn_players(&g)
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

	// Init Pinfo structs and assign row, col info eg record player positions into the struct
	p1 := pInfo{}
	p2 := pInfo{}
	g.Pinfo[1] = &p1
	g.Pinfo[2] = &p2
	g.Pinfo[1].Row = p1_row
	g.Pinfo[1].Col = p1_col
	g.Pinfo[2].Row = p2_row
	g.Pinfo[2].Col = p2_col

	// Give players mana
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

func cleanup_tiles( g *gameSpace, locs [][2]int ) {
    // fmt.Println("Before changing tiles")
    // pretty_print(g.Arena)
    for _, v := range locs {
        // fmt.Printf("Thing: %v at index %d", v, i)
        row := v[0]
        col := v[1]
        // fmt.Println("Row:", row)
        // fmt.Println("Col:", col)


        g.Arena[row][col] = 0
    }
    // fmt.Println("after changing tiles")
    // pretty_print(g.Arena)

}

// Check if a player has triggered a gameover state
func check_gameover(g *gameSpace) bool {
	if g.Gameover.Player[1] || g.Gameover.Player[2] || g.Gameover.Player[0] {
		return true
	}
	return false
}

// Determine who won the game and who lost
func get_winner_loser_info(g *gameSpace, p1 string, p2 string) (string, map[int]float64) {
    // Player values dont match p1 and p2 so wrong user can be declared winner
	var rstring string
    var winner int
    var loser int
    scores := make(map[int]float64)
    if !check_gameover(g) {
		rstring = fmt.Sprintf("Nobody has lost. No winner or loser to identify\n")
        scores[1] = 0.5
        scores[2] = 0.5
		g.Gameover.Winner = "Nobody"
		g.Gameover.Loser = "Nobody"
		g.Gameover.DeathMsg = "The game has ended in a tie"
    } else {
		if g.Gameover.Player[1] {
			winner = 2
			loser = 1
            scores[1] = -1.0
            scores[2] = 1.0
		} else {
			winner = 1
			loser = 2
            scores[1] = 1.0
            scores[2] = -1.0
		}
        if g.Gameover.Player[0] || (g.Gameover.Player[1] && g.Gameover.Player[2]) {
            fmt.Println("Both players have died.")
            rstring = "Both players have died. Cause: " + g.Gameover.Message[0]
            scores[1] = 0.5
            scores[2] = 0.5
			g.Gameover.Winner = "Nobody"
			g.Gameover.Loser = "Nobody"
			g.Gameover.DeathMsg = g.Gameover.Message[0]
        } else {
        	rstring = fmt.Sprintf("Player %v has died.\nDeath message: %v\n\n", loser, g.Gameover.Message[loser])
			g.Gameover.DeathMsg = g.Gameover.Message[loser]
            if winner == 2 {
                    fmt.Println("Not switched")
                    rstring += fmt.Sprintf("Winner is %v\nLoser is %v\n", p2, p1)
                    g.Gameover.Winner = p2
                    g.Gameover.Loser = p1
            } else {
                // winner == 1
                    fmt.Println("Not switched")
                    rstring += fmt.Sprintf("Winner is %v\nLoser is %v\n", p1, p2)
                    g.Gameover.Winner = p1
                    g.Gameover.Loser = p2
            }
        }
	}
    fmt.Println("p1: ", p1)
    fmt.Println("p2: ", p2)
    fmt.Println("winner", winner)
    fmt.Println("loser", loser)
    fmt.Println("g.Gameover:", g.Gameover)
    fmt.Println("g.Gameover.Winner:", g.Gameover.Winner)
    fmt.Println("g.Gameover.Loser:", g.Gameover.Loser)
	return rstring, scores
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

// sets gameover to true for player in the gameOver object as well as a death message
// Need to call get_winner_loser_info() to see game over messages
func game_over(g *gameSpace, player int, message string) {
	// fmt.Printf("Player %v Died. Player %v Loses\n", player, player)
    // fmt.Println("Message:", message)
    g.Gameover.Player[player] = true
	g.Gameover.Message[player] = message
}
