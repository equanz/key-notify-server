package main

// import
import(
  "github.com/gin-gonic/gin"
)

// server
func main(){
  r := gin.Default()
  r.POST("/hard/on", func(c *gin.Context){

    c.String(200, "Key On!")
  })

  r.Run()
}

