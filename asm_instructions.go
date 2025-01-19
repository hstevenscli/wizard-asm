package main

import "fmt"
// import "errors"


// Shoot a fireball in the indicated direction
func asm_summon_magma( g *gameSpace, player int, row int, col int ) {

    var d_row int = row
    var d_col int = col
    fmt.Println("Testing for player", player)
    fmt.Println("PROW:", d_row)
    fmt.Println("PCOL:", d_col)

    // Changing to coords instead of direction

    //TODO MAYBE
    // An arr to track where the magma pops up, depending on final implementation, may or may not need this
    // var marr[9][2]int
    // fmt.Println(marr)

    // Summon Magma in 3x3 square
    fmt.Printf("Magma Center: [%v, %v]\n", d_row, d_col)
    for row := -1; row < 2; row++ {
        for col := -1; col < 2; col++ {
            t_row := d_row + row
            t_col := d_col + col
            var valid bool = is_valid_loc(t_row, t_col, g.Size)
            fmt.Printf("Burning loc: [%v, %v]| Valid loc: %v\n", t_row, t_col, valid)
            if valid {
				player_val, p_hit := check_player(t_row, t_col, g)
				if p_hit {
					game_over(player_val)
				}
				// FOR DEBUGGING, ERASE LATER
                g.Arena[t_row][t_col] = 3
            }
        }
    }
}

// Summon an acid puddle on target location
func asm_summon_acid ( g *gameSpace, player int, row int, col int ) {
    var d_row int = row
    var d_col int = col
	fmt.Println("Called summon acid, player:", player)
    fmt.Println("PROW:", d_row)
    fmt.Println("PCOL:", d_col)
	fmt.Println(g)
	var valid bool = is_valid_loc(d_row, d_col, g.Size)
	if valid {
		
	}
}



func asm_move( g *gameSpace, player int, direction string ) {

    // Using player int to make a pointer to easily access g.PX_loc
    var p_loc *[2]int
    if player == 1 {
        p_loc = &g.P1_loc
    } else {
        p_loc = &g.P2_loc
    }

    var d_row int
    var d_col int

    switch direction {
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
    }
    // Keep destination within valid index
    d_row = within_valid_range( d_row, g.Size )
    d_col = within_valid_range( d_col, g.Size )

    // If space is clear move the player there
    // In this case a space is clear if there is not a player there
    if g.Arena[d_row][d_col] != 1 && g.Arena[d_row][d_col] != 2 {
        if g.Arena[d_row][d_col] != 0 {
            fmt.Println("You stepped in Acid!! (I think). Player stepped in:", g.Arena[d_row][d_col])
            game_over(player)
        }
        easy_move_wrapper(g, player, d_row, d_col)
    }
}

