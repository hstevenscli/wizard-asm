package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)


func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getBattleReplay(c *gin.Context) {
	c.JSON(http.StatusOK, battleReplay)
}

func runGame() {
    // Initialize size 
    var size int = 16

    // Initialize the gamespace to size = 10
    var g gameSpace = init_gamespace(size)

    fmt.Printf("SIZE: %v\n", size)

    // Spawn players in gameSpace
    spawn_players( &g )
	battleReplay = replay{}
	gameover = gameOver{}

    program := read_json_to_bp("./program.json")
	program1 := read_json_to_bp("./program1.json")


	// Put starting arena state into replay
	starting_arena := frame{ 
		ArenaFrame: deep_copy_arena(g.Arena),
		Player: 0,
		Action: "Starting State",
		Mana: 0,
		Count: -1,
	}
	battleReplay.Frame = append(battleReplay.Frame, starting_arena)

	game_loop_temp( &g, program, program1)
	print_replay( battleReplay )
	get_winner_loser_info()
}

func main() {
    // // Initialize size 
    // var size int = 4

    // // Initialize the gamespace to size = 10
    // var g gameSpace = init_gamespace(size)

    // // Spawn players in gameSpace
    // spawn_players( &g )
	// pretty_print(g.Arena)
	// asm_divination( &g, 1, 0)
	// fmt.Println(g.Pinfo[1].CrystalBall)
	// fmt.Println(g.Pinfo[1].Mana)
	// get_winner_loser_info()

	router := gin.Default()
    router.GET("/albums", getAlbums)
	router.GET("/br", getBattleReplay)
	router.GET("/game", func(c *gin.Context) {
		runGame()
		c.String(200, "Game has been run")
	})



    router.Run("localhost:8080")
}


