package main

import "fmt"


func main() {

    // Initialize size 
    var size int = 10

    // Initialize the gamespace to size = 10
    var g gameSpace = init_gamespace(size)

    fmt.Printf("SIZE: %v\n", size)

    // Spawn players in gameSpace
    spawn_players( &g )

    // Pretty print gameSpace
    // pretty_print(g.Arena)
    pretty_print(g.Arena)
    // asm_summon_magma(&g, 2, 1, 1)
    fmt.Println()
	// asm_summon_acid(&g, 2, 0, 0)
    // asm_shield(&g, 2, 1)
    // asm_lightning(&g, 1, "s")
    // asm_divination(&g, 1, 10)
    program := read_json_to_bp("./program.json")
	program1 := read_json_to_bp("./program1.json")
	// loop_through_battleProgram( &g, program )
	game_loop_temp( &g, program, program1)
    pretty_print(g.Arena)

	
	// for i, instruction := range program.Instructions {
	// 	fmt.Println(i)
	// 	fmt.Println("Spell Name:", instruction.Instruction, "Args:", instruction.Args)
	// }

}


