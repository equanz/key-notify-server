package main

import(
  "os"
  "net/http"
  "net/url"
)

/* Slack botでの通知
 * is_onboard: bool trueの際，Onとして通知
*/
func send_form_message(is_onboard bool) (error){
  const SLACK_ENDPOINT = "https://slack.com/api/chat.postMessage" // endpoint url

   req_val := url.Values{
    "token": {os.Getenv("SLACK_BOT_TOKEN")},
    "channel": {os.Getenv("CHANNEL")},
    "as_user": {"true"}, // use default botname
   }

  if is_onboard == true {
    // on board
    req_val.Add("text", os.Getenv("ON_MESSAGE"))
  } else {
    // out board
    req_val.Add("text", os.Getenv("OFF_MESSAGE"))
  }

  res, err := http.PostForm(SLACK_ENDPOINT, req_val)

  if err != nil {
    return err
  }

  defer res.Body.Close() // async close

  return nil
}
