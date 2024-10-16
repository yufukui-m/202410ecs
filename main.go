package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"os"
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

	r.GET("/image", func(c *gin.Context) {
		data, err := os.ReadFile("./file/test.jpg")
		if err != nil {
			c.String(http.StatusInternalServerError, "error")
			return
		}
		c.Header("Content-Type", "image/jpeg")
		_, _ = c.Writer.Write(data)
	})

	r.GET("/wait/:sec", func(c *gin.Context) {
		secStr := c.Param("sec")
		sec, _ := strconv.Atoi(secStr)
		time.Sleep(time.Duration(sec) * time.Second)
		c.String(http.StatusOK, "waited %d second(s)", sec)
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
