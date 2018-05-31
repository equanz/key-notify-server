/* Package: SQLクエリ発行，処理
*/
package sql_query

// import
import(
  "log"
  "os"
  "time"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

// connection var
var db *sql.DB

type Key_info struct{
  Time string
  State string
  Key_info_id int
}

// package initialize
func init(){
  db_local, err := sql.Open("mysql", os.Getenv("MARIA_USER") + ":" + os.Getenv("MARIA_PASS") + "@tcp(" + os.Getenv("MARIA_HOST") +":" + os.Getenv("MARIA_PORT") +")/" + os.Getenv("MARIA_DB"))
  if err == nil {
    db = db_local
  } else {
    log.Fatal(err)
  }
}

/* 統計データ全件抽出
*/
func Get_all_statistics() ([]Key_info, error){
  rows, err := db.Query("SELECT * FROM key_info")
  if err != nil {
    return nil, err
  }
  var info_array []Key_info
  defer rows.Close()
  for rows.Next() {
    var info Key_info
    if err := rows.Scan(&info.Time, &info.State, &info.Key_info_id); err != nil {
      return nil, err
    }
    info_array = append(info_array, info) // append to array
  }

  return info_array, nil
}

/* 統計データ指定日時から最新まで抽出
 * fd: time.Time 取得する統計値の指定日時
*/
func Get_statistics(fd time.Time) ([]Key_info, error){
  parse_time, err_parse := time.Parse("2006-01-02 15:04:05", fd.String())
  if err_parse != nil {
    return nil, err_parse
  }
  rows, err_que := db.Query("SELECT * FROM key_info WHERE time >= ? ORDER BY time DESC", parse_time)
  if err_que != nil {
    return nil, err_que
  }
  var info_array []Key_info
  defer rows.Close()
  for rows.Next() {
    var info Key_info
    if err := rows.Scan(&info.Time, &info.State, &info.Key_info_id); err != nil {
      return nil, err
    }
    info_array = append(info_array, info)
  }
  return info_array, nil
}

/* app_idがDBに格納されていれば真
 * app_id: string 要求されたapp_id
*/
func Has_app_id(app_id string) (bool, error){
  rows, err := db.Query("SELECT COUNT(*) FROM app_id WHERE app_id = ?", app_id)
  if err != nil {
    return false, err
  }
  var db_count int
  defer rows.Close()
  for rows.Next() {
    if err := rows.Scan(&db_count); err != nil {
      return false, err
    }
  }
  if db_count == 0 {
    return false, nil
  } else {
    return true, nil
  }
}

/* 鍵の状態をデータベースに挿入
 * state: string(enum("ON","OFF")) 鍵の状態
*/
func Insert_status(state string) (error){
  now := time.Now()
  jst := time.FixedZone("Asia/Tokyo", 9*60*60)
  nowJST := now.In(jst)
  time_shape := nowJST.Format("2006-01-02 15:04:05")
  _, err_exec := db.Exec("INSERT INTO `key_info` (`time`, `state`) VALUES (?, ?)", time_shape, state)
  if err_exec != nil{
    return err_exec
  }
  return nil
}
