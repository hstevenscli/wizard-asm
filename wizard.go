package main

import "fmt"


func main() {

    type gameSpace struct {
        Size int
        Arena [][]int
    }

    g := gameSpace { Size: 10, Arena: make([][]int, 10) }
    for i := range g.Arena {
        g.Arena[i] = make([]int, g.Size)
    }

    fmt.Println(g.Arena)
    pretty_print(g.Arena)
    // Make gameSpace Slice of 10 slices
    // [[] [] [] [] [] [] [] [] [] []]
    var size int = 10
    fmt.Printf("SIZE: %v\n", size)
    // gameSpace := make([][]int, size)

    // // Make each subslice == size
    // for i := range gameSpace {
    //         gameSpace[i] = make([]int, size)
    // }

    // // Spawn players in gameSpace
    // spawn_players( gameSpace )

    // // Pretty print gameSpace
    // pretty_print(gameSpace)
}

// Print out gameSpace in a human-viewable format
func pretty_print( slice [][]int ) {
    for i := range slice {
            fmt.Println(slice[i])
    }
}

func spawn_players( slice [][]int ) {
    slice[0][1] = 1
    slice[0][2] = 2
    fmt.Println("Spawned some players")
}

// Shoot a fireball in the indicated direction
func asm_cast_fireball( slice [][]int ) {
        fmt.Println("FIREBALL")
        fmt.Println(slice)
}


