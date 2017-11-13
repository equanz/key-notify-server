package main

// import
import(
  "github.com/gin-gonic/gin"
)

// server
func main(){
  r := gin.Default()

  // root
  r.GET("/", func(c *gin.Context){
    c.String(200, "Hello, world!")
  })

  // key on
  r.POST("/hard/on", func(c *gin.Context){
    c.String(200, "Key On!")
  })

  r.Run()
}

