package main

import (
    "fmt"
    "math/rand"
    "strings"
)

// var gameover gameOver

// Summon magma in 3x3 square centered on target row,col
func asm_summon_magma( g *gameSpace, player int, row int, col int ) {

    var d_row int = row
    var d_col int = col
    // fmt.Println("Testing for player", player)
    // fmt.Println("PROW:", d_row)
    // fmt.Println("PCOL:", d_col)

    // Changing to coords instead of direction

    //TODO MAYBE
    // An arr to track where the magma pops up, depending on final implementation, may or may not need this
    // var marr[9][2]int
    // fmt.Println(marr)

    // Summon Magma in 3x3 square
    // fmt.Printf("Magma Center: [%v, %v]\n", d_row, d_col)
    for row := -1; row < 2; row++ {
        for col := -1; col < 2; col++ {
            t_row := d_row + row
            t_col := d_col + col
            var valid bool = is_valid_loc(t_row, t_col, g.Size)
            // fmt.Printf("Burning loc: [%v, %v]| Valid loc: %v\n", t_row, t_col, valid)
            if valid {
				player_val, p_hit := check_player(t_row, t_col, g)
				if p_hit {
					game_over(g, player_val, "Burned in magma")
				}
				// FOR DEBUGGING, ERASE LATER
                g.Arena[t_row][t_col] = 3
            }
        }
    }
    deplete_mana(g, player, 20)
}

// Summon an acid puddle on target location
func asm_summon_acid ( g *gameSpace, player int, row int, col int ) {
    var d_row int = row
    var d_col int = col
	var valid bool = is_valid_loc(d_row, d_col, g.Size)
	if valid {
        player_val, p_hit := check_player(d_row, d_col, g)
        if p_hit {
            game_over(g, player_val, "Burned in acid")
        }
        // Turn the space into acid
        g.Arena[d_row][d_col] = 3
	}
    deplete_mana(g, player, 10)
}



func asm_move( g *gameSpace, player int, direction string ) {

	// making this (slice p_loc) to fit with the old code in the switch statements
	var row, col int
	row = g.Pinfo[player].Row
	col = g.Pinfo[player].Col
    p_loc := []int{row, col}

    var d_row int
    var d_col int

    dir := strings.ToLower(direction)
    switch dir {
    case "n":
        d_row = p_loc[0] - 1
        d_col = p_loc[1]
    case "s":
        d_row = p_loc[0] + 1
        d_col = p_loc[1]
    case "w":
        d_row = p_loc[0]
        d_col = p_loc[1] - 1
    case "e":
        d_row = p_loc[0]
        d_col = p_loc[1] + 1
    default:
        d_row = p_loc[0]
        d_col = p_loc[1]
    }
    // Keep destination within valid index
    d_row = within_valid_range( d_row, g.Size )
    d_col = within_valid_range( d_col, g.Size )
    // fmt.Println("DRow:", d_row)
    // fmt.Println("Dcol:", d_col)

    // If space is clear move the player there
    // In this case a space is clear if there is not a player there
    if g.Arena[d_row][d_col] != 1 && g.Arena[d_row][d_col] != 2 {
        if g.Arena[d_row][d_col] == 3 {
            fmt.Printf("You stepped in Acid!! (I think). Player %v stepped in: %v\n", player, g.Arena[d_row][d_col])
            game_over(g, player, "Stepped in acid")
        }
        easy_move_wrapper(g, player, d_row, d_col)
    }
}

// Teleport to target location; Random location if row,int are -1
func asm_teleport( g *gameSpace, player int, row int, col int ) {
    // Random teleport
    if row <= -1 || col <= -1 {
        row = rand.Intn(g.Size)
        col = rand.Intn(g.Size)

    // Teleport to specified location
    } else {
        row = within_valid_range( row, g.Size )
        col = within_valid_range( col, g.Size )
    }
    // Check for a player collision
    player_val, p_hit := check_player(row, col, g) 
    if p_hit && player_val != player {
        fmt.Println("Teleport double death")
        game_over(g, 0, "Teleportation death")
    }
    // Draw changes to g.Arena
    easy_move_wrapper(g, player, row, col)
    deplete_mana(g, player, 100)
}

// Apply a protection to g.PX_prot, valid values: 1 fire, 2 lightning, 3 acid
// @TODO Need to implement checks in necessary places for protections
// At minimum in asm_summon_magma/acid/lightning asm_move/teleport
func asm_shield( g *gameSpace, player int, d_type int ) {
	g.Pinfo[player].Prot = d_type
    deplete_mana(g, player, 10)
}

func asm_wait( player int ) {
}

func asm_recharge( g *gameSpace, player int, amount int ) {
	g.Pinfo[player].Mana += amount
    if g.Pinfo[player].Mana > 200 {
        game_over(g, player, "Mana Overcharge")
    }
}

// Hone in on location of enemy, player = player who cast it,
// level range 0-5 inclusive 0 is most accurate, 5 is least accurate
func asm_divination( g *gameSpace, player int, level int ) {
	var tprow, tpcol int
	if player == 1 {
		fmt.Println(1)
		tprow = g.Pinfo[2].Row
		tpcol = g.Pinfo[2].Col
	} else {
		tprow = g.Pinfo[1].Row
		tpcol = g.Pinfo[1].Col
	}
	row_shift := rand.Intn(level * 2 + 1) - level
	col_shift := rand.Intn(level * 2 + 1) - level
	r_row := tprow + row_shift
	r_col := tpcol + col_shift
	g.Pinfo[player].CrystalBall[0] = r_row
	g.Pinfo[player].CrystalBall[1] = r_col
    deplete_mana(g, player, 200 - (level * 10))
}


func asm_lightning( g *gameSpace, player int, direction string ) {

    dir := strings.ToLower(direction)

	var row, col int
	row = g.Pinfo[player].Row
	col = g.Pinfo[player].Col
    p_loc := []int{row, col}

    var s_row int = p_loc[0]
    var s_col int = p_loc[1]
    var row_mod int = 0
    var col_mod int = 0

    switch dir {
    case "n":
        row_mod = -1
    case "s":
        row_mod = 1
    case "e":
        col_mod = 1
    case "w":
        col_mod = -1
    case "ne":
        row_mod = -1
        col_mod = 1
    case "nw":
        row_mod = -1
        col_mod = -1
    case "se":
        row_mod = 1
        col_mod = 1
    case "sw":
        row_mod = 1
        col_mod = -1
    }

    for i:= 0; i < g.Size-1; i++ {
        s_row = s_row + row_mod
        s_col = s_col + col_mod
        if !is_valid_loc(s_row, s_col, g.Size ) {
            break
        }
        player_val, p_hit := check_player(s_row, s_col, g)
        if p_hit {
            game_over(g, player_val, "Struck by lightning")

        }
        // For debugging, remove later
        g.Arena[s_row][s_col] = 7
    }
    deplete_mana(g, player, 100)
}
