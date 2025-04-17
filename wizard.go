package main

import (
	"context"
	// "crypto/rand"
	// "encoding/json"
	// "fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	// "go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)


func connectToMongo() *mongo.Client {
    // load env variable
    err := godotenv.Load()
	
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    // get mongo uri
    uri := os.Getenv("MONGODB_URI")
    docs := "www.mongodb.com/docs/drivers/go/current/"
    if uri == "" {
        log.Fatal("Set your 'MONGODB_URI' environment variable. " +
			"See: " + docs +
			"usage-examples/#environment-variable")
    }

    // connect with the right options
    client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
        log.Fatal("Error connecting to Mongodb", err)
	}


    return client
}




func mongoMiddleware(client *mongo.Client) gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Set("mongoClient", client)
        c.Next()
    }
}


func authorizeMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authed := validSession(c)
        if !authed {
            c.JSON(401, gin.H{"status": "Session Not Found"})
            c.Abort()
            return
        }
        c.Next()
    }
}

func authorizeMiddlewareAdmin() gin.HandlerFunc {
    return func(c *gin.Context) {
        authed := validSession(c)
        if !authed {
            c.JSON(401, gin.H{"status": "Session Not Found"})
            c.Abort()
            return
        }
		sess := getSessionInfo(c)
		var isAdmin bool
		if sess.Username == "titan" || sess.Username == "titan1" || sess.Username == "bob" {
			isAdmin = true
		}
		if !isAdmin {
            c.JSON(401, gin.H{"status": "Requires Admin Priveleges"})
            c.Abort()
            return
		}
        c.Next()
    }
}

func getCookieDomain() string {
	cookiedomain := os.Getenv("COOKIE_DOMAIN")
	if cookiedomain == "" {
		cookiedomain = "wizard.nothingdangerous.live"
	}
	return cookiedomain
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}


    client := connectToMongo()
    defer func() {
        if err := client.Disconnect(context.TODO()); err != nil {
            log.Fatal("Error disconnecting from MongoDB:", err)
		}
	}()


	router := gin.Default()
	router.Static("/frontend", "./frontend")

    router.Use(mongoMiddleware(client))
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080", "http://localhost:8081", "http://127.0.0.1:8080", "http://127.0.0.1:8081"}, // Change to match your frontend URL
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
        AllowCredentials: true,
    }))

	router.GET("/", func(c *gin.Context) {
		c.File("./frontend/index.html")
	})

	router.GET("/reports", authorizeMiddlewareAdmin(), func(c *gin.Context) {
		c.File("./frontend/bug_reports.html")
	})


    // GET ROUTES
    router.GET("/cookie", cookieHandler)
	router.GET("/battlereplay", getBattleReplay)
	router.GET("/session", getSession)
    router.GET("/battleprogram/:username", getBattleProgramByUsernameHandler)
	router.GET("/duels/:username", authorizeMiddleware(), getDuel)
	router.GET("/duels/random", authorizeMiddleware(), getDuelRandom)
	router.GET("/duels", func(c *gin.Context){
		c.JSON(409, gin.H{"status": "No name provided for user to duel"})
	})
    router.GET("/users/:username", authorizeMiddleware(), getUser)
	router.GET("/bugreports", authorizeMiddlewareAdmin(), getBugReports)
    router.GET("/scoreboard", authorizeMiddleware(), getScoreboard)


    // PUT the stuff in this function into a helper function to authenticate on protected routes
    router.GET("/testsession", authorizeMiddleware(), func(c *gin.Context) {
            c.JSON(200, gin.H{"status": "Boobies"})
    })

    // POST ROUTES
	router.POST("/battleprogram", authorizeMiddleware(), postBattleProgram)
    router.POST("/users", postUsers)
    router.POST("/login", postLogin)
    router.POST("/logout", authorizeMiddleware(), postLogout)
	// router.POST("/game", func(c *gin.Context) {
	// 	runGame()
        // c.JSON(201, gin.H{"msg": "Game has been run"})
	// })
	router.POST("/bugreports", postBugReport)

    // DELETE
    // router.DELETE("/bugreports/:id", func(c *gin.Context) {
    //     id := c.Param("id")
    //     c.JSON(200, gin.H{"status": "Deleted report with id " + id})
    // })
    router.DELETE("/bugreports/:id", authorizeMiddlewareAdmin(), deleteBugReport)

    router.Run("0.0.0.0:"+port)
}
