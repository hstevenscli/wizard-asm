package main

import "fmt"


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


    // DIRECTION LOGIC
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
    // END DIRECTION LOGIC

    // Affect all spaces adjacent to where the fireball lands

    fmt.Printf("Fireball landed: [%v, %v]\n", d_row, d_col)
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

