package main 

import "github.com/gin-going/gin"

func main(){

  r := gin.Default() 
  r.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "ping", 
    })
  })

  r.run(":3001")
}