package main

import (
	"encoding/json"
	"log"
	"os"
	"fmt"
)

type instruction struct {
    Instruction string `json:"instruction"`
    Args []interface{} `json:"args,omitempty"`
}

type battleProgram struct {
    User string `json:"user"`
    Instructions []instruction `json:"instructions"`
	Ptr int
	Player int
	Lstart int
	Niterations int
}

type battleReplay struct {
    Replay [][][]int

}

// Read some json and put it into a battleProgram struct
func read_json_to_bp( json_file string ) battleProgram {
    // Read json file into data
    data, err := os.ReadFile(json_file)
    if err != nil {
        log.Fatalf("Failed to read file: %v", err)
    }

    var program battleProgram
    err = json.Unmarshal(data, &program)
    if err != nil {
        log.Fatalf("Failed to unmarshal json data: %v", err)
    }

    return program
}

// Extract a string from the arg interface; return an empty string if the arg is not a string
func extractStringArg( arg interface{} ) string {
    switch v := arg.(type) {
    case string:
		return arg.(string)
    default:
        fmt.Printf("Not a string: %T\n", v)
		return ""
    }
}

func extractIntArg( arg interface{} ) int {
	switch v := arg.(type) {
	case float64:
		return int(arg.(float64))
	default:
		fmt.Printf("Not an int: %T\n", v)
		// Return a number unlikely to be given as an argument
		return -100
	}
}


func execute_instruction( g *gameSpace, bp *battleProgram ) string {
    // fmt.Println("PLAYER", bp.Player, "TAKING AN ACTION")
	var r_msg string
    var c_args []interface{}
    var instr string
    if bp.Ptr < 0 || bp.Ptr >= len(bp.Instructions) {
        instr = "PTRDEATH"
    } else {
        c_args = bp.Instructions[bp.Ptr].Args
        instr = bp.Instructions[bp.Ptr].Instruction
    }

	switch instr {
		// GAME MODIFYING INSTRUCTIONS START
	case "WAIT":
		asm_wait(bp.Player)
		r_msg = "WAIT"
	case "MOVE":
		dir := extractStringArg( c_args[0])
		if dir != "" {
			asm_move( g, bp.Player, dir)
		} else {
			log.Println("String not found as argument for Move")
		}
		r_msg = "MOVE"
	case "MAGMA":
		arg1 := extractIntArg( c_args[0])
		arg2 := extractIntArg( c_args[1])
		if arg1 != -100 && arg2 != -100 {
			asm_summon_magma( g, bp.Player, arg1, arg2)
		} else {
			log.Println("Int not found as argument for Magma")
		}
		r_msg = "MAGMA"
	case "TELEPORT":
		arg1 := extractIntArg( c_args[0])
		arg2 := extractIntArg( c_args[1])
		if arg1 != -100 && arg2 != -100 {
			asm_teleport( g, bp.Player, arg1, arg2)
		} else {
			log.Println("Int not found as argument for Teleport")
		}
		r_msg = "TELEPORT"
	case "ACID":
		arg1 := extractIntArg( c_args[0])
		arg2 := extractIntArg( c_args[1])
		if arg1 != -100 && arg2 != -100 {
			asm_summon_acid( g, bp.Player, arg1, arg2)
		} else {
			log.Println("Int not found as argument for Acid")
		}
		r_msg = "ACID"
	case "SHIELD":
		arg1 := extractIntArg( c_args[0] )
		if arg1 != -100 {
			asm_shield( g, bp.Player, arg1)
		} else {
			log.Println("Int not found as argument for Shield")
		}
		r_msg = "SHIELD"
	case "RECHARGE":
		arg1 := extractIntArg( c_args[0] )
		if arg1 != -100 {
			asm_recharge( g, bp.Player, arg1)
		} else {
			log.Println("Int not found as argument for Recharge")
		}
		r_msg = "RECHARGE"
	case "DIVINATION":
		arg1 := extractIntArg( c_args[0] )
		if arg1 != -100 {
			asm_divination( g, bp.Player, arg1)
		} else {
			log.Println("Int not found as argument for Divination")
		}
		r_msg = "DIVINATION"
	case "LIGHTNING":
		dir := extractStringArg( c_args[0])
		if dir != "" {
			asm_lightning( g, bp.Player, dir)
		} else {
			log.Println("String not found as argument for Lightning")
		}
		r_msg = "LIGHTNING"
	// GAME MODIFYING INSTRUCTIONS END
	// INSTRUCTIONS MODIFYING BATTLEPROGRAM START
	case "SLOOP":
		bp.Niterations= extractIntArg( c_args[0]) - 1
		bp.Lstart= bp.Ptr + 1
		r_msg = "SLOOP"
	case "ELOOP":
		if bp.Niterations> 0 {
			bp.Niterations -= 1
			bp.Ptr = bp.Lstart- 1
		}
		r_msg = "ELOOP"
	case "JUMP":
		arg1 := extractIntArg( c_args[0])
		if arg1 != -100 {
			// change this into a call to asm_jump
			// Need logic to see if value passed is valid and it will kill player
			// when executing
			bp.Ptr += arg1
            bp.Ptr--
		} else {
			log.Println("Int not found as argument for Jump")
		}
		r_msg = "JUMP"
	case "CJUMP":
		r_msg = "CJUMP"
    case "PTRDEATH":
        game_over(bp.Player, "Died due to Pointer Death")
		r_msg = "PTRDEATH"
	default:
		log.Printf("INSTRUCTION NOT RECOGNIZED: %v", bp.Instructions[bp.Ptr].Instruction)
		r_msg = "BADINSTRUCTION"
	}
	return r_msg


}
// main function for looping through a players battle program

//@TODO Need to get the correct player value in here somehow, 
// Either put the player value into the battleProgram like above
// Or extract it from the username or something
func loop_through_battleProgram( g *gameSpace, bp battleProgram ) {
	// For marking where a sloop instruction happens
	bp.Player = 2
	var timeout int
	for bp.Ptr < len(bp.Instructions) {
		// fmt.Printf("Executing instr %v: %v\n", bp.Ptr, bp.Instructions[bp.Pointer])
		execute_instruction( g, &bp)
		bp.Ptr++
		timeout++
		fmt.Println("Timeout: ", timeout)
		fmt.Println()
		if timeout >= 1000 {
			fmt.Println("TIMEOUT!")
			break
		}
	}
	fmt.Println("Loop ended. Game over")
}

func check_gameover() bool {
	if gameover.Player[1] || gameover.Player[2] || gameover.Player[0] {
		return true
	}
	return false
}

func get_winner_loser_info() {
    if !check_gameover() {
        fmt.Println("Nobody has lost. No winner or loser to identify")
    } else {
		var winner int
		var loser int
		if gameover.Player[1] {
			winner = 2
			loser = 1
		} else {
			winner = 1
			loser = 2
		}
        if gameover.Player[0] {
            fmt.Println("Both players have died.")
        } else {
            fmt.Printf("Player %v has died.\nDeath message: %v\n", loser, gameover.Message[loser])
            fmt.Printf("Winner is player %v\nLoser is player %v\n", winner, loser)
        }
	}
}

func game_loop_temp( g *gameSpace, bp1 battleProgram, bp2 battleProgram ) {
	bp1.Player = 1
	bp2.Player = 2
	var count int
	// If a player builds a program wrong the program crashes, fix it so that the 
	// player loses if their bp.Ptr leaves the scope of their program
	for i := 0; i <= 1000; i++ {
		//P1 Chunk
		p1_action := execute_instruction( g, &bp1 )
		bp1.Ptr++

		fmt.Println("===================================")
		fmt.Printf("Player 1 | #Actions: %v\n", count+1)
		fmt.Printf("Action: %v\n", p1_action)
		fmt.Printf("Mana: %v\n", g.Pinfo[1].Mana)
		fmt.Println("===================================")
		fmt.Println()
		pretty_print(g.Arena)
		fmt.Println()
		fmt.Println()
		fmt.Println()
		if check_gameover() {
			break
		}

		//P2 Chunk
		p2_action := execute_instruction( g, &bp2 )
		bp2.Ptr++

		fmt.Println("===================================")
		fmt.Printf("Player 2 | #Actions: %v\n", count+1)
		fmt.Printf("Action: %v\n", p2_action)
		fmt.Printf("Mana: %v\n", g.Pinfo[2].Mana)
		fmt.Println("===================================")
		fmt.Println()
		pretty_print(g.Arena)
		fmt.Println()
		fmt.Println()
		fmt.Println()
		if check_gameover() {
			break
		}

		// Increment game counter
		count++
	}
	if count >= 1000 {
		fmt.Println("Game Timed Out")
	}
	get_winner_loser_info()
}
