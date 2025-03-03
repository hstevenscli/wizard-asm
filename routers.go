package main

import (
    "fmt"
    "context"
    "net/http"
	"github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/v2/mongo"
    "go.mongodb.org/mongo-driver/v2/bson"
    "golang.org/x/crypto/bcrypt"
)

type user struct {
    Username string `bson:"username"`
    Password string `bson:"password"`
}

func getClient(c *gin.Context) *mongo.Client {
    client, exists := c.Get("mongoClient")
    if !exists {
        fmt.Println("doesnt exist")
    }
    mongoClient := client.(*mongo.Client)
    return mongoClient

}

func authenticatePassword(hashedPassword []byte, password []byte) bool {
    // use bcrypt to compare passwords
    password = []byte("MyDarkSecret")

    // Hashing the password with the default cost of 10
    hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
    if err != nil {
        panic(err)
    }

    // Comparing the password with the hash
    err = bcrypt.CompareHashAndPassword(hashedPassword, password)
    fmt.Println(err) // nil means it is a match

    return true
}

func postLogin(c *gin.Context) {
    var loginUser user
    var result user
    if err :=  c.BindJSON(&loginUser); err != nil {
        fmt.Println("Bad format with username and password")
        return
    }
    // check if user and password are in db
    client := getClient(c)
    coll := client.Database("wizardb").Collection("users")

    filter := bson.D{{ "username", loginUser.Username }}


    // find appropriate user
    err := coll.FindOne(context.TODO(), filter).Decode(&result)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            fmt.Println("user not found")
            c.JSON(401, gin.H{"status": "unauthorized", "message": "invalid username or password"})
            return
        }
        c.JSON(500, gin.H{"message": "server error"})
        return
    }
    // Encrypt password and verify here

    authenticated := authenticatePassword([]byte(result.Password), []byte(loginUser.Password))
    if authenticated {
        c.JSON(200, gin.H{"status":"success"})
    } else {
        c.JSON(401, gin.H{"status": "unauthorized", "message": "invalid username or password"})
    }
}

func postUsers(c *gin.Context) {
    var newUser user
    mongoClient := getClient(c)
    coll := mongoClient.Database("wizardb").Collection("users")

    // Get username and password from the request
    if err := c.BindJSON(&newUser); err != nil {
        fmt.Println("Bad request")
        return
    }

    result, err := coll.InsertOne(context.TODO(), newUser)
    if err != nil {
        panic(err)
    }

    c.JSON(http.StatusOK, gin.H{"msg": result})
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getBattleReplay(c *gin.Context) {
	c.JSON(http.StatusOK, battleReplay)
}

func postBattleProgram(c *gin.Context) {
	var newBattleProgram battleProgram

	// Call bind json and return if an error occurs
	if err := c.BindJSON(&newBattleProgram); err != nil {
		return
	}

	// Save program to the database
	fmt.Println(newBattleProgram)

	// Modify battleProgram
	// newBattleProgram.Player = 2

	c.IndentedJSON(http.StatusCreated, newBattleProgram)
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
	battleReplay.Frames = append(battleReplay.Frames, starting_arena)

	game_loop_temp( &g, program, program1)
	print_replay( battleReplay )
	get_winner_loser_info()
}

