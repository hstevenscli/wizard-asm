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
	PlayerVal int
}

// Read some json and put it into a battleProgram struct
func read_json( json_file string ) battleProgram {
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

// main function for looping through a players battle program

//@TODO Need to get the correct player value in here somehow, 
// Either put the player value into the battleProgram like above
// Or extract it from the username or something
func loop_through_battleProgram( g *gameSpace, bp battleProgram ) {
	for bp.Ptr < len(bp.Instructions) {
		// fmt.Printf("Executing instr %v: %v\n", bp.Ptr, bp.Instructions[bp.Pointer])
		c_int_args := bp.Instructions[bp.Ptr].Args
		switch bp.Instructions[bp.Ptr].Instruction {
		case "WAIT":
			fmt.Println("Casting Wait")
			asm_wait(1)
		case "MOVE":
			fmt.Println("Casting Move")
		case "FIREBALL":
			fmt.Println("Casting Fireball")
			asm_summon_magma( g, 1, int(c_int_args[0].(float64)), int(c_int_args[1].(float64)))
		case "TELEPORT":
			fmt.Println("Casting Teleport")

		}
		bp.Ptr++
	}
}
