package main

import (
  "net/http"
  "fmt"
  "strconv"
  "github.com/gin-gonic/gin"
  "geekshubs.com/demo-server/internal/sum"
)

func main() {
  r := gin.Default()

  r.GET("/", func(c *gin.Context) {

	x, err := strconv.Atoi(c.DefaultQuery("x", "2"))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}
	
	y, err := strconv.Atoi(c.DefaultQuery("y", "2"))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}

	fmt.Println(c.DefaultQuery("y", "2"))
	fmt.Println(c.DefaultQuery("x", "2"))

    c.JSON(http.StatusOK, gin.H{"data": sum.Sum(x,y)})
  })

  r.Run()
}
