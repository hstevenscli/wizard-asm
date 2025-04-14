package main

import (
	"fmt"
	// "strconv"
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
    Username string `bson:"username"`
    Password string `bson:"password"`
    BP battleProgram `bson:"bp"`
    Score float64
}

type report struct {
	// ID primitive.ObjectID `bson:"_id,omitempty"`
	ID      interface{} `bson:"_id,omitempty"`
	Message string `bson:"message"`
	Email string `bson:"email,omitempty"`
}

// type report struct {
// 	ID      primitive.ObjectID `bson:"_id,omitempty" json:"ID"`
// 	Message string             `bson:"message" json:"Message"`
// 	Email   string             `bson:"email,omitempty" json:"Email"`
// }

func getBugReports(c *gin.Context) {
	var reports []report

    mongoClient := getClient(c)
    coll := mongoClient.Database("wizardb").Collection("reports")

    cursor, err := coll.Find(context.TODO(), bson.D{})
    if err != nil {
        c.JSON(500, gin.H{"status": "Database Error", "error": err.Error()})
        return
    }
    defer cursor.Close(context.TODO())

    for cursor.Next(context.TODO()) {
        var r report
        if err := cursor.Decode(&r); err != nil {
            c.JSON(500, gin.H{"status": "Error unpacking cursor into reports", "error": err.Error()})
            return
        }
		fmt.Printf("Decoded report: ID type=%T, ID value=%v, Message=%s, Email=%s\n", 
        r.ID, r.ID, r.Message, r.Email)
		reports = append(reports, r)
    }

    if err := cursor.Err(); err != nil {
        c.JSON(500, gin.H{"status": "Cursor error", "error": err.Error()})
        return
    }

    c.JSON(200, reports)
}

var ErrEmptyBattleProgram = errors.New("Battle Program is empty")

// For getting the mongo client into each handler function
// because the mongo client is started in the main func
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
        c.SetCookie("My_Cookie", "", -1, "/",  getCookieDomain(), false, false)
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
        c.SetCookie("My_Cookie", sessionID, 60*60*24*30, "/",  getCookieDomain(), false, false)

        fmt.Println("session store: ", sessionStore)
        c.JSON(200, gin.H{"status": "successfully logged in"})
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
	c.JSON(http.StatusOK, gin.H{"status": "This version of replay has deprecated"})
}

func postBattleProgram(c *gin.Context) {
	var newBattleProgram battleProgram

	// Call bind json and return if an error occurs
	if err := c.BindJSON(&newBattleProgram); err != nil {
        c.JSON(400, gin.H{"status": "Bad Request"})
		return
	}

    // Get database collection
    mongoClient := getClient(c)
    coll := mongoClient.Database("wizardb").Collection("users")

    filter := bson.D{{ "username", newBattleProgram.User }}
    update := bson.D{{"$set", bson.D{{ "bp", newBattleProgram }} }}
    _, err := coll.UpdateOne(context.TODO(), filter, update)
    if err != nil {
        fmt.Println("Error", err)
        c.JSON(500, gin.H{"status": "server error"})
        return
    }

    // fmt.Println("Matched Count:", updresult.MatchedCount)
    // fmt.Println("Modified Count:", updresult.ModifiedCount)

    c.IndentedJSON(201, gin.H{"status": "Program created/saved successfully"})
}

func postBugReport(c *gin.Context) {
	var newBugReport report
	if err := c.BindJSON(&newBugReport); err != nil {
		c.JSON(400, gin.H{"status": "Bad request; json error"})
		return
	}

	// check length of message
	if len(newBugReport.Message) > 7500 {
		c.JSON(413, gin.H{"status": "Message too long"})
		return
	}

	if len(newBugReport.Email) > 100 {
		c.JSON(413, gin.H{"status": "Email too long"})
		return
	}

	// get database and collection
	mongoClient := getClient(c)
	coll := mongoClient.Database("wizardb").Collection("reports")

	result, err := coll.InsertOne(context.TODO(), newBugReport)
	if err != nil {
		c.JSON(500, gin.H{"status": "Server error"})
		return
	}
	fmt.Println("Result:", result)
	c.JSON(201, gin.H{"status": "Report made successfully"})
}

func getDuel(c *gin.Context) {
	var ses session
	ses = getSessionInfo(c)
    mongoClient := getClient(c)
	user1 := ses.Username
	user2 := c.Param("username")
	bp1, err := getAndLoadBattleProgram(user1, mongoClient)
	// error handling for bp1
	if err != nil {
		if err.Error() == "No user found" {
			c.JSON(409, gin.H{"status": "User1 not found"})
			fmt.Println("Err:", err)
			return
		} else if err.Error() == "Battle Program is empty" {
			c.JSON(409, gin.H{"status": "Battle program 1 is empty"})
			fmt.Println("Err:", err)
			return
		} else {
			c.JSON(500, gin.H{"status": "server error"})
			fmt.Println("Err:", err)
			return
		}
	}
	// error handling for bp2
	bp2, err := getAndLoadBattleProgram(user2, mongoClient)
	if err != nil {
		if err.Error() == "No user found" {
			c.JSON(409, gin.H{"status": "User2 not found"})
			fmt.Println("Err:", err)
			return
		} else if err.Error() == "Battle Program is empty" {
			c.JSON(409, gin.H{"status": "Battle program 2 is empty"})
			fmt.Println("Err:", err)
			return
		} else {
			c.JSON(500, gin.H{"status": "server error"})
			fmt.Println("Err:", err)
			return
		}
	}
	fmt.Printf("Running a game between Player1: %v and Player2: %v\n", user1, user2)
	br := runBattle(bp1, bp2, mongoClient)
	c.JSON(200, br)
}

func getDuelRandom(c *gin.Context) {
	var ses session
    var err error
	ses = getSessionInfo(c)
    mongoClient := getClient(c)
	user1 := ses.Username
	bp1, err := getAndLoadBattleProgram(user1, mongoClient)
	// error handling for bp1
	if err != nil {
		if err.Error() == "No user found" {
			c.JSON(409, gin.H{"status": "User1 not found"})
			fmt.Println("Err:", err)
			return
		} else if err.Error() == "Battle Program is empty" {
			c.JSON(409, gin.H{"status": "Battle program 1 is empty"})
			fmt.Println("Err:", err)
			return
		} else {
			c.JSON(500, gin.H{"status": "server error"})
			fmt.Println("Err:", err)
			return
		}
	}

    var bp2 battleProgram
    for {
        bp2, err = getAndLoadBattleProgramRandom(mongoClient)
        if err == nil && bp1.User != bp2.User {
            break
        }
        if err != ErrEmptyBattleProgram && err != nil {
            // return or handle other types of errors
            fmt.Println("ERROR:", err)
            c.JSON(500, gin.H{"status": "server error"})
            return
        }
        // else: err == ErrEmptyBattleProgram — retry
    }
    fmt.Println("BP1:", bp1)
    fmt.Println("BP2:", bp2)

    // Cases to handle: finds an empty BP, user finds their own BP

    br := runBattle(bp1, bp2, mongoClient)
    c.JSON(200, br)
}


// fetches a battleProgram from the database and returns it to the caller
func getAndLoadBattleProgram(username string, mongoClient *mongo.Client) (battleProgram, error){
    // var userToLookup user
    var found user
	var bp battleProgram

    coll := mongoClient.Database("wizardb").Collection("users")
    filter := bson.D{{ "username", username }}

	// query for users battleprogram
    err := coll.FindOne(context.TODO(), filter).Decode(&found)
    if err != nil {
        if err == mongo.ErrNoDocuments {
			return bp, errors.New("No user found")
        }
		return bp, err
    }

	// check if battleprogram is empty
	if found.BP.User == "" {
		return bp, ErrEmptyBattleProgram
	}
	bp = found.BP
	return bp, nil
}

func getAndLoadBattleProgramRandom(mongoClient *mongo.Client) (battleProgram, error) {
    var bp battleProgram

    coll := mongoClient.Database("wizardb").Collection("users")
    // Create aggregation pipeline with $sample
    pipeline := mongo.Pipeline{
        {{"$sample", bson.D{{"size", 1}}}},
    }

    cursor, err := coll.Aggregate(context.TODO(), pipeline)
    if err != nil {
        return bp, err
    }
    defer cursor.Close(context.TODO())

    var results []user
    if err := cursor.All(context.TODO(), &results); err != nil {
        return bp, err
    }

    if len(results) == 0 {
        return bp, mongo.ErrNoDocuments
    }

	if results[0].BP.User == "" {
		return bp, ErrEmptyBattleProgram
	}

    bp = results[0].BP
    return bp, err
}

func getUserByUsername(username string, mongoClient *mongo.Client) (user, error) {
    var queried_user user
    coll := mongoClient.Database("wizardb").Collection("users")
    filter := bson.D{{ "username", username }}

    err := coll.FindOne(context.TODO(), filter).Decode(&queried_user)

    return queried_user, err
}

// Use this to query for a user and update their score based on if they win or lose
func updateUserScore(username string, modifier float64, mongoClient *mongo.Client) error {
    coll := mongoClient.Database("wizardb").Collection("users")

    filter := bson.M{"username": username}
    update := bson.M{"$inc": bson.M{"score": modifier}}

    _, err := coll.UpdateOne(context.TODO(), filter, update)
    return err
}

func getBattleProgramByUsernameHandler(c *gin.Context) {
    // var userToLookup user
    username := c.Param("username")
    fmt.Println("USername:", username)
    var found user

    // if err := c.BindJSON(&userToLookup); err != nil {
    //     c.JSON(400, gin.H{"status": "bad request"})
    //     return
    // }
    mongoClient := getClient(c)
    coll := mongoClient.Database("wizardb").Collection("users")
    filter := bson.D{{ "username", username }}

    err := coll.FindOne(context.TODO(), filter).Decode(&found)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            c.JSON(409, gin.H{"status": "user not found"})
            return
        }
        c.JSON(500, gin.H{"status": "server error"})
        return
    }

    fmt.Println("Battle program found:", found)
    c.JSON(200, found.BP)
}


func getUser(c *gin.Context) {
    username := c.Param("username")
    mongoClient := getClient(c)
    user, err := getUserByUsername(username, mongoClient)
    user.Password = "*"
    if err != nil {
        c.JSON(500, gin.H{"status": "server error"})
        return
    }
    c.JSON(200, user)
}

func getBattlePrograms(c *gin.Context) {
    // mongoClient := getClient(c)
    // coll := mongoClient.Database("wizardb").Collection("users")

    // err := coll.Find
}

func getSubsetOfBattlePrograms(c *gin.Context) {
    // mongoClient := getClient(c)
    return
}


func cookieHandler(c *gin.Context) {
    cookie, err := c.Cookie("My_Cookie")
    if err != nil {
        sessionID := "123"
        fmt.Println("Number:", sessionID)
        fmt.Println("Cookie My_Cookie not found, Setting cookie My_Cookie with secret value")
        cookie = "notset"
        c.SetCookie("My_Cookie", sessionID, 3600, "/",  getCookieDomain(), false, false)
    }

    fmt.Printf("Cookie Value: %s\n", cookie)
}


func runBattle(bp1 battleProgram, bp2 battleProgram, mongoClient *mongo.Client) replay {
	var size int = 16
	var g gameSpace = init_gamespace(size)
	// spawn_players(&g)

	br := replay{}
	// Put starting arena state into replay
	starting_arena := frame{ 
		ArenaFrame: deep_copy_arena(g.Arena),
		Player: 0,
		PlayerName: "",
		Action: "Starting State",
		Mana: 0,
		Count: -1,
	}
	br.Frames = append(br.Frames, starting_arena)

	game_loop_temp( &g, bp1, bp2, &br)
	print_replay( br )
	// fmt.Println("gameover struct:", g.Gameover)
    message, scores := get_winner_loser_info(&g, bp1.User, bp2.User)
    g.Gameover.Conclusion = message
	br.GameoverInfo = *g.Gameover

    // Update Scores
    fmt.Println("USERS IN FIGHT", bp1.User, bp2.User)
    if bp1.User != "bob" && bp2.User != "bob" {
        updateUserScore(bp1.User, scores[1], mongoClient)
        updateUserScore(bp2.User, scores[2], mongoClient)
    }

    // check users scores
    // user1, err := getUserByUsername(bp1.User, mongoClient)
    // user2, err := getUserByUsername(bp2.User, mongoClient)

    // fmt.Println("ERROR: ", err)

    // fmt.Println("User 1:", user1)
    // fmt.Println("User 2:", user2)
	br.Opp = bp2.User

	return br
}

// Deprecated
// func runGame() {
//     // Initialize size 
//     var size int = 16

//     // Initialize the gamespace to size = 10
//     var g gameSpace = init_gamespace(size)

//     fmt.Printf("SIZE: %v\n", size)

//     // Spawn players in gameSpace
//     spawn_players( &g )
// 	battleReplay := replay{}
// 	gameover := gameOver{}
// 	g.Gameover = &gameover

//     program := read_json_to_bp("./program.json")
// 	program1 := read_json_to_bp("./program1.json")


// 	// Put starting arena state into replay
// 	starting_arena := frame{ 
// 		ArenaFrame: deep_copy_arena(g.Arena),
// 		Player: 0,
// 		Action: "Starting State",
// 		Mana: 0,
// 		Count: -1,
// 	}
// 	battleReplay.Frames = append(battleReplay.Frames, starting_arena)

// 	game_loop_temp( &g, program, program1, &battleReplay)
// 	print_replay( battleReplay )
// 	// get_winner_loser_info(&g)
// }

