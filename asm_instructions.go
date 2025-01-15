package main

import "fmt"


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
                g.Arena[t_row][t_col] = 3
            }
        }
    }
}

func move( g *gameSpace, player int, direction string ) {
}

