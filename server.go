package main

import (
	"os"
	"encoding/json"
	"log"
)

type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// type user struct {
//     Username string             `json:"username"`
//     BattleProgram battleProgram `json:"battleprogram"`
// }

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
