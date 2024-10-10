package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	r.GET("/bingo", func(c *gin.Context) {
		n := rand.Intn(74) + 1
		c.String(http.StatusOK, strconv.Itoa(n))
	})

	r.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
