package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(n int) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v += n
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value() int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	return c.v
}

func main() {
	rand.Seed(time.Now().UnixNano())

	counter := SafeCounter{}

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

	r.GET("/count", func(c *gin.Context) {
		v := counter.Value()

		c.String(http.StatusOK, "counter: %d", v)
	})

	r.GET("/count/:n", func(c *gin.Context) {
		nStr := c.Param("n")
		n, _ := strconv.Atoi(nStr)
		counter.Inc(n)
		v := counter.Value()

		c.String(http.StatusOK, "counter: %d", v)
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
