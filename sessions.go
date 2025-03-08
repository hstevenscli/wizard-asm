package main

import (
    "time"
	"crypto/rand"
    // "fmt"
    "encoding/hex"
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


