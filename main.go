/* Package: サーバ本体
*/
package main

// import
import(
  "fmt"
  "github.com/equanz/key-notify-server/sql_query"
  "github.com/gin-gonic/gin"
  //"github.com/gin-gonic/gin/binding"
  "os"
  "github.com/nlopes/slack"
  "encoding/json"
  "time"
)

type StatusJSON struct{
  AppID string `json:"app_id" binding:"required"`
}

// serverの作成
func main(){
  r := gin.Default()

  // root
  r.GET("/", func(c *gin.Context){
    c.String(200, "Hello, world!")
  })

  // '/api' group
  api := r.Group("/api")
  {
    /* key on or off
     * :status: string "on", "off"文字列をそれぞれハードウェアのon, offとして処理
    */
    api.POST("/hard/:status", func(c *gin.Context){
      var body StatusJSON
      c.BindJSON(&body) // bind post body
      state := c.Param("status")

      has_app_id, err := sql_query.Has_app_id(body.AppID) // search app_id
      if err == nil{
        if body.AppID != "" && has_app_id == true{
          if state == "on"{
            send_form_message(true)
            c.String(200, "ON")
          } else if state == "off"{
            send_form_message(false)
            c.String(200, "OFF")
          } else{
            // error
            c.String(400, "Bad Request")
          }
        } else{
          // error
          c.String(400, "Bad Request")
        }
      } else{
        // error
        c.String(503, "Service Unavailable")
      }
    })

    /* on/off統計データ(rawデータ)を指定日時から最新まで返す
     * first_date: string(DATETIMEフォーマット) 取得する統計値の指定日時
    */
    api.GET("/statistics", func(c *gin.Context){
      q := c.Request.URL.Query() // query params
      fd, fd_ok := q["first_date"]

      if fd_ok == true {
        time, err := time.Parse("2006-01-02 15:04:05", fd[0]) // parse to format
        if err != nil {
          fmt.Println(err)
          c.String(400, "Bad Request")
        } else {
          info_array := sql_query.Get_statistics(time.String())
          info_array_json, err := json.Marshal(info_array) // generate json bytes from struct
          if err != nil {
            fmt.Println(err)
            c.String(400, "Bad Request")
          } else {
            c.String(200, string(info_array_json)) // stringify and response
          }
        }
      } else{
        info_array := sql_query.Get_all_statistics()
        info_array_json, err := json.Marshal(info_array) // generate json bytes from struct
        if err != nil {
          fmt.Println(err)
          c.String(400, "Bad Request")
        } else {
          c.String(200, string(info_array_json)) // stringify and response
        }
      }
    })
  }

  r.Run()
}

/* Slack botでの通知
 * is_onboard: bool trueの際，Onとして通知
*/
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
    rtm.SendMessage(rtm.NewOutgoingMessage("OFF Board!", os.Getenv("CHANNEL")))
  }
}

