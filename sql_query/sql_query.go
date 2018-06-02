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

/* 統計データ指定日時から指定日時まで抽出
 * fd: time.Time 統計値の取得を開始する指定日時
 * ed: time.Time 統計値の取得を終わる指定日時
*/
func Get_statistics(fd time.Time,ed time.Time) ([]Key_info, error){
  first_time := fd.Format("2006-01-02 15:04:05")
  end_time := ed.Format("2006-01-02 15:04:05")
  rows, err := db.Query("SELECT * FROM key_info WHERE time >= ? AND time <= ? ORDER BY time", first_time, end_time)
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
    info_array = append(info_array, info)
  }
  return info_array, nil
}

/* 最新データ一件を抽出
*/
func Get_latest_state() ([]Key_info, error){
  rows, err := db.Query("SELECT * FROM key_info ORDER BY time DESC LIMIT 1")
  if err != nil {
    return nil , err
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

/* 指定した日時の一つ前のデータを抽出
 * fd: time.Time 取得する値の指定日時
*/
func Get_before_state(fd time.Time) ([]Key_info, error){
  form_time := fd.Format("2006-01-02 15:04:05")
  rows, err := db.Query("SELECT * FROM key_info WHERE time <= ? ORDER BY time DESC LIMIT 1", form_time)
  if err != nil {
    return nil , err
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
  jst := time.FixedZone("Asia/Tokyo", 9*60*60) //Hour*Minute*Second
  nowJST := now.In(jst)
  time_format := nowJST.Format("2006-01-02 15:04:05")
  _, err_exec := db.Exec("INSERT INTO `key_info` (`time`, `state`) VALUES (?, ?)", time_format, state)
  if err_exec != nil{
    return err_exec
  }
  return nil
}
