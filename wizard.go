package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)


func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {

	router := gin.Default()
    router.GET("/albums", getAlbums)


    // Initialize size 
    var size int = 10

    // Initialize the gamespace to size = 10
    var g gameSpace = init_gamespace(size)

    fmt.Printf("SIZE: %v\n", size)

    // Spawn players in gameSpace
    spawn_players( &g )
	// fmt.Println(g)

    // Pretty print gameSpace
    fmt.Println("==================STARTING ARENA====================")
    pretty_print(g.Arena)
    fmt.Println("==================STARTING ARENA====================")
    fmt.Println()
    program := read_json_to_bp("./program.json")
	program1 := read_json_to_bp("./program1.json")

	game_loop_temp( &g, program, program1)
    // pretty_print(g.Arena)

    router.GET("/bp", func(c *gin.Context){
        c.IndentedJSON(http.StatusOK, program)
        // c.String(http.StatusOK, "HEY THERE")
    })
    router.Run("localhost:8080")
}


