package main

import (
    "time"
	"crypto/rand"
    // "fmt"
    "encoding/hex"
	"github.com/gin-gonic/gin"
    "fmt"
)

type session struct {
    Username string
    Timestamp time.Time
}

var sessionStore = make(map[string]session)

func generateSessionID () (string, error) {
    bytes := make([]byte, 16)
    _, err := rand.Read(bytes)
    if err != nil {
        return "", err
    }
    return hex.EncodeToString(bytes), nil
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
