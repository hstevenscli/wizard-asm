package main

import (
	"fmt"
	"log"
	"strings"
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

type frame struct {
	ArenaFrame [][]int
	Player int
	Action string
	Mana int
	Count int
	Args []interface{}
}

type replay struct {
    Frames []frame
	GamoverInfo gameOver
}

// var battleReplay replay

// Add a snapshot of the Arena to the battleReplay
func add_frame_to_replay( arena [][]int, player int, pinfo pInfo, count int, action string, args []interface{}, battleReplay *replay ) {
	// Make frame struct
	f := frame{
		ArenaFrame: deep_copy_arena(arena),
		Player: player,
		Action: action,
		Mana: pinfo.Mana,
		Count: count,
		Args: args,
	}
	// add frame to replay
	battleReplay.Frames = append(battleReplay.Frames, f)
}

func deep_copy_arena ( arena [][]int ) [][]int {
	acopy := make([][]int, len(arena))
	for i:= range arena {
		acopy[i] = make([]int, len(arena[i]))
		copy(acopy[i], arena[i])
	}
	return acopy
}

// For debugging basically
func print_replay( br replay) {
	for i := 0; i < len( br.Frames ); i++ {
		for j := 0; j < len(br.Frames[i].ArenaFrame); j++ {
			fmt.Println(br.Frames[i].ArenaFrame[j])
		}
		fmt.Printf("Player %v | #Actions: %v\n", br.Frames[i].Player, br.Frames[i].Count)
		fmt.Printf("Action: %v | Args: %v\n", br.Frames[i].Action, br.Frames[i].Args)
		fmt.Printf("Mana: %v\n", br.Frames[i].Mana)
		fmt.Println()
	}

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

// Extract an int from the arg interface; return -100 if the an int is unable to be extracted
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

func execute_instruction( g *gameSpace, bp *battleProgram ) (string, []interface{}){
    // fmt.Println("PLAYER", bp.Player, "TAKING AN ACTION")
    var c_args []interface{}
    var instr string
    if bp.Ptr < 0 || bp.Ptr >= len(bp.Instructions) {
        instr = "PTRDEATH"
    } else {
        c_args = bp.Instructions[bp.Ptr].Args
        instr = bp.Instructions[bp.Ptr].Instruction
    }
	instr = strings.ToUpper(instr)

	switch instr {
		// GAME MODIFYING INSTRUCTIONS START
	case "WAIT":
		asm_wait(bp.Player)
	case "MOVE":
		dir := extractStringArg( c_args[0])
		if dir != "" {
			asm_move( g, bp.Player, dir)
		} else {
			log.Println("String not found as argument for Move")
		}
	case "MAGMA":
		arg1 := extractIntArg( c_args[0])
		arg2 := extractIntArg( c_args[1])
		if arg1 != -100 && arg2 != -100 {
			asm_summon_magma( g, bp.Player, arg1, arg2)
		} else {
			log.Println("Int not found as argument for Magma")
		}
	case "TELEPORT":
		arg1 := extractIntArg( c_args[0])
		arg2 := extractIntArg( c_args[1])
		if arg1 != -100 && arg2 != -100 {
			asm_teleport( g, bp.Player, arg1, arg2)
		} else {
			log.Println("Int not found as argument for Teleport")
		}
	case "ACID":
		arg1 := extractIntArg( c_args[0])
		arg2 := extractIntArg( c_args[1])
		if arg1 != -100 && arg2 != -100 {
			asm_summon_acid( g, bp.Player, arg1, arg2)
		} else {
			log.Println("Int not found as argument for Acid")
		}
	case "SHIELD":
		arg1 := extractIntArg( c_args[0] )
		if arg1 != -100 {
			asm_shield( g, bp.Player, arg1)
		} else {
			log.Println("Int not found as argument for Shield")
		}
	case "RECHARGE":
		arg1 := extractIntArg( c_args[0] )
		if arg1 != -100 {
			asm_recharge( g, bp.Player, arg1)
		} else {
			log.Println("Int not found as argument for Recharge")
		}
	case "DIVINATION":
		arg1 := extractIntArg( c_args[0] )
		if arg1 != -100 {
			asm_divination( g, bp.Player, arg1)
		} else {
			log.Println("Int not found as argument for Divination")
		}
	case "LIGHTNING":
		dir := extractStringArg( c_args[0])
		if dir != "" {
			asm_lightning( g, bp.Player, dir)
		} else {
			log.Println("String not found as argument for Lightning")
		}
	// GAME MODIFYING INSTRUCTIONS END
	// INSTRUCTIONS MODIFYING BATTLEPROGRAM START
	case "SLOOP":
		bp.Niterations= extractIntArg( c_args[0]) - 1
		bp.Lstart= bp.Ptr + 1
	case "ELOOP":
		if bp.Niterations > 0 {
			bp.Niterations -= 1
			bp.Ptr = bp.Lstart- 1
		}
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
	case "CJUMP":
    case "PTRDEATH":
        game_over(g, bp.Player, "Died due to Pointer Death")
	default:
		log.Printf("INSTRUCTION NOT RECOGNIZED: %v", bp.Instructions[bp.Ptr].Instruction)
	}
	return instr, c_args


}

func game_loop_temp( g *gameSpace, bp1 battleProgram, bp2 battleProgram, br *replay ) {
	bp1.Player = 1
	bp2.Player = 2
	var count int

	for i := 0; i <= 1000; i++ {
		//P1 Chunk
		p1_action, args := execute_instruction( g, &bp1 )
		bp1.Ptr++
		add_frame_to_replay( g.Arena, 1, *g.Pinfo[1], count, p1_action, args, br)
		if check_gameover(g) {
			break
		}
        fmt.Println("mana in game loop", g.Pinfo[1].Mana)

		//P2 Chunk
		p2_action, args := execute_instruction( g, &bp2 )
		bp2.Ptr++
		add_frame_to_replay( g.Arena, 2, *g.Pinfo[2], count, p2_action, args, br)
		if check_gameover(g) {
			break
		}

		// Increment game counter
		count++
	}
	if count >= 1000 {
		fmt.Println("Game Timed Out")
	}
}
