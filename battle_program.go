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

func extractArg( arg interface{} ) {
    switch v := arg.(type) {
    case string:
        fmt.Println("String:", v)
    case float64:
        fmt.Println("Number:", v)
    default:
        fmt.Printf("Something else: %T\n", v)
    }
}
// main function for looping through a players battle program

//@TODO Need to get the correct player value in here somehow, 
// Either put the player value into the battleProgram like above
// Or extract it from the username or something
func loop_through_battleProgram( g *gameSpace, bp battleProgram ) {
	for bp.Ptr < len(bp.Instructions) {
		// fmt.Printf("Executing instr %v: %v\n", bp.Ptr, bp.Instructions[bp.Pointer])
		c_args := bp.Instructions[bp.Ptr].Args
		switch bp.Instructions[bp.Ptr].Instruction {
		case "WAIT":
			fmt.Println("Casting Wait")
			asm_wait(1)
		case "MOVE":
			fmt.Println("Casting Move")
            strArg, ok := c_args[0].(string)
            fmt.Println("StrArg:",strArg)
            if ok {
                asm_move( g, 1, strArg)
            } else {
                log.Println("String not found as argument for Move")
            }
		case "MAGMA":
			fmt.Println("Casting Fireball")
			asm_summon_magma( g, 1, int(c_args[0].(float64)), int(c_args[1].(float64)))
		case "TELEPORT":
			fmt.Println("Casting Teleport")

        default:
            log.Printf("INSTRUCTION NOT RECOGNIZED: %v", bp.Instructions[bp.Ptr].Instruction)
		}
		bp.Ptr++
	}
}
