package main

import (
    "fmt"
    // "strconv"
    "time"
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
	BP battleProgram
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
    // Comparing the password with the hash
    err := bcrypt.CompareHashAndPassword(hashedPassword, password)
    fmt.Println(err) // nil means it is a match
    if err == nil {
        return true
    }
    return false
}

// check if the cookie exists and if it has a valid session id
func validSession(c *gin.Context) bool {
    cookie, err := c.Cookie("My_Cookie")
    if err != nil {
        fmt.Println("No cookie please login")
        // c.JSON(403, gin.H{"status": "Cookie not found"})
        return false
    }
    if _, exists := sessionStore[cookie]; !exists {
        fmt.Println("Session not found please log in")
        // c.JSON(403, gin.H{"status": "Session not found"})
        return false
    }
    // c.JSON(200, gin.H{"status": "Boobies"})
    return true
}

// ALWAYS MAKE SURE SESSION IS VALID BEFORE USING THIS FUNCTION
func getSessionInfo(c *gin.Context) session {
	cookie, _ := c.Cookie("My_Cookie")
	ses, _ := sessionStore[cookie]
	return ses
}

func getSession(c *gin.Context) {
	authed := validSession(c)
	if !authed {
		c.JSON(401, gin.H{"status": "Unauthenticated"})
	} else {
		ses := getSessionInfo(c)
		fmt.Println("SESSIONINFO:", ses)
		c.JSON(200, gin.H{"session": ses})
	}
}

func postLogout(c *gin.Context) {
    cookie, err := c.Cookie("My_Cookie")
    if err != nil {
        c.JSON(409, gin.H{"status": "Already logged out/cookie not found"})
    } else {
        c.SetCookie("My_Cookie", "", -1, "/",  "localhost", false, false)
        delete(sessionStore, cookie)
        c.JSON(200, gin.H{"status": "Logged out successfully"})
    }
}

func postLogin(c *gin.Context) {
    var loginUser user
    var result user
    if err :=  c.BindJSON(&loginUser); err != nil {
        c.JSON(400, gin.H{"status": "Bad Request"})
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
        c.JSON(500, gin.H{"status": "server error"})
        return
    }
    // Encrypt password and verify here

    authenticated := authenticatePassword([]byte(result.Password), []byte(loginUser.Password))

    // Put sessionID in session store
    if authenticated {
        sessionID, err := generateSessionID()
        // Session ID Error
        if err != nil {
            fmt.Println("Error Generating ID")
            c.JSON(500, gin.H{"status": "Internal server error, please try again later"})
            return
        }
        sessionStore[sessionID] = session{ Username: loginUser.Username, Timestamp: time.Now().Add(time.Hour *24*30) }
        c.SetCookie("My_Cookie", sessionID, 60*60*24*30, "/",  "localhost", false, false)

        fmt.Println("session store: ", sessionStore)
        c.JSON(200, gin.H{"status": sessionID})
    } else {
        c.JSON(401, gin.H{"status": "unauthorized", "message": "invalid username or password"})
    }
}

func postUsers(c *gin.Context) {
    // capture username and password of user that is about to be created
    var newUser user
    mongoClient := getClient(c)
    coll := mongoClient.Database("wizardb").Collection("users")

    // Get username and password from the request
    if err := c.BindJSON(&newUser); err != nil {
        c.JSON(400, gin.H{"status": "Bad Request"})
        return
    }

    // make sure the username isnt already taken
    filter := bson.D{{ "username", newUser.Username }}
    // .Err() will query without needing to save a result into a struct
    err := coll.FindOne(context.TODO(), filter).Err()
    // err is nil when it finds something
    if err == nil {
        // username is not free, we need to return
        c.JSON(409, gin.H{"status": "Username already taken"})
        return
    } else if err != mongo.ErrNoDocuments {
        c.JSON(500, gin.H{"error": "Internal Server Error"})
        return
    }

    // encrypt password
    // Hashing the password with the default cost of 10
    password := []byte(newUser.Password)
    hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
    if err != nil {
        c.JSON(500, gin.H{"error": "Internal Server Error"})
        panic(err)
    }
    newUser.Password = string(hashedPassword)

    _, err = coll.InsertOne(context.TODO(), newUser)
    if err != nil {
        panic(err)
    }

    c.JSON(http.StatusOK, gin.H{"status": "User created successfully"})
}

func getBattleReplay(c *gin.Context) {
	c.JSON(http.StatusOK, battleReplay)
}

func postBattleProgram(c *gin.Context) {
	var newBattleProgram battleProgram

	// Call bind json and return if an error occurs
	if err := c.BindJSON(&newBattleProgram); err != nil {
        c.JSON(400, gin.H{"status": "Bad Request"})
		return
	}
    // TODO Save the program to the DB

	fmt.Println("NewBP:", newBattleProgram)



    c.IndentedJSON(201, gin.H{"status": "Program created successfully"})
}

func cookieHandler(c *gin.Context) {
    cookie, err := c.Cookie("My_Cookie")
    if err != nil {
        sessionID := "123"
        fmt.Println("Number:", sessionID)
        fmt.Println("Cookie My_Cookie not found, Setting cookie My_Cookie with secret value")
        cookie = "notset"
        c.SetCookie("My_Cookie", sessionID, 3600, "/",  "localhost", false, false)
    }

    fmt.Printf("Cookie Value: %s\n", cookie)
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

