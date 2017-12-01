package main

// import
import(
  "github.com/gin-gonic/gin"
  "os"
  "github.com/nlopes/slack"
)

// server
func main(){
  r := gin.Default()

  // root
  r.GET("/", func(c *gin.Context){
    c.String(200, "Hello, world!")
    send_form_message(true)
  })

  // key on
  r.POST("/hard/on", func(c *gin.Context){
    c.String(200, "Key On!")
  })

  r.Run()
}

func send_form_message(is_onboard bool){
  // slack bot api instance
  api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
  // slack rtm instance
  rtm := api.NewRTM()
  go rtm.ManageConnection()

  if is_onboard == true {
    // on board
    rtm.SendMessage(rtm.NewOutgoingMessage("ON Board!", os.Getenv("CHANNEL")))
  } else {
    // out board
    rtm.SendMessage(rtm.NewOutgoingMessage("OUT Board!", os.Getenv("CHANNEL")))
  }
}

