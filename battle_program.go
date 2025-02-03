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

func execute_instruction( g *gameSpace, bp *battleProgram, loop_started *int, n_iterations *int ) {

	c_args := bp.Instructions[bp.Ptr].Args


	switch bp.Instructions[bp.Ptr].Instruction {
		// GAME MODIFYING INSTRUCTIONS START
	case "WAIT":
		asm_wait(bp.Player)
		fmt.Println("Wait Happened")
	case "MOVE":
		dir := extractStringArg( c_args[0])
		if dir != "" {
			asm_move( g, bp.Player, dir)
		} else {
			log.Println("String not found as argument for Move")
		}
		fmt.Println("Move Happened")
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
		*n_iterations = extractIntArg( c_args[0]) - 1
		*loop_started = bp.Ptr + 1
		fmt.Println("SLOOP Happened")
	case "ELOOP":
		if *n_iterations > 0 {
			*n_iterations -= 1
			bp.Ptr = *loop_started - 1
		}
		fmt.Println("ELOOP Happened")
	case "JUMP":
		arg1 := extractIntArg( c_args[0])
		if arg1 != -100 {
			// change this into a call to asm_jump
			// Need logic to see if value passed is valid and it will kill player
			// when executing
			bp.Ptr += arg1
		} else {
			log.Println("Int not found as argument for Jump")
		}
	case "CJUMP":
	default:
		log.Printf("INSTRUCTION NOT RECOGNIZED: %v", bp.Instructions[bp.Ptr].Instruction)
	}


}
// main function for looping through a players battle program

//@TODO Need to get the correct player value in here somehow, 
// Either put the player value into the battleProgram like above
// Or extract it from the username or something
func loop_through_battleProgram( g *gameSpace, bp battleProgram ) {
	// For marking where a sloop instruction happens
	bp.Player = 2
	var loop_started int
	var n_iterations int
	var timeout int
	for bp.Ptr < len(bp.Instructions) {
		// fmt.Printf("Executing instr %v: %v\n", bp.Ptr, bp.Instructions[bp.Pointer])
		execute_instruction( g, &bp, &loop_started, &n_iterations)
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
