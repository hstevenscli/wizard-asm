package main

import (
	"encoding/json"
	"log"
	"os"
)

type instruction struct {
    Command string `json:"command"`
    Args []interface{} `json:"args,omitempty"`
}

type battleProgram struct {
    User string `json:"user"`
    Instructions []instruction `json:"instructions"`
}

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

