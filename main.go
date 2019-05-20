package main

import (
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	router := gin.Default()

	router.GET("/kitty", kittyHandler)
	router.GET("/pup", pupHandler)
	router.Run(":8080")
}
