/* Package: SQLクエリ発行，処理
*/
package sql_query

// import
import(
  "log"
  "os"
  "time"
  "database/sql"
  _ "github.com/lib/pq"
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
  db_local, err := sql.Open("postgres", "host=" + os.Getenv("DB_HOST") + " user=" + os.Getenv("DB_USER") + " dbname=" + os.Getenv("DB_DATABASE") + " password=" + os.Getenv("DB_PASS") + " sslmode=require")
  if err == nil {
    _, err_tz := db_local.Exec("SET TIME ZONE 'Asia/Tokyo'") // setup client timezone
    if err_tz == nil {
      db = db_local
    } else {
      log.Fatal(err_tz)
    }
  } else {
    log.Fatal(err)
  }
}

/* 統計データ全件抽出
*/
func Get_all_statistics() ([]Key_info, error){
  rows, err := db.Query("SELECT * FROM key_info ORDER BY time")
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
  first_time := fd.Format("2006-01-02 15:04:05 -07:00")
  end_time := ed.Format("2006-01-02 15:04:05 -07:00")
  rows, err := db.Query("SELECT * FROM key_info WHERE time >= $1 AND time <= $2 ORDER BY time", first_time, end_time)
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
func Get_latest_state() (Key_info, error){
  nil_key := Key_info{}
  rows, err := db.Query("SELECT * FROM key_info ORDER BY time DESC LIMIT 1")
  if err != nil {
    return nil_key, err
  }
  var info_array []Key_info
  defer rows.Close()
  for rows.Next() {
    var info Key_info
    if err := rows.Scan(&info.Time, &info.State, &info.Key_info_id); err != nil {
      return nil_key, err
    }
    info_array = append(info_array, info)
  }
  return info_array[0], nil
}

/* 指定した日時の一つ前のデータを抽出
 * fd: time.Time 取得する値の指定日時
*/
func Get_before_state(fd time.Time) (Key_info, error){
  nil_key := Key_info{}
  form_time := fd.Format("2006-01-02 15:04:05 -07:00")
  rows, err := db.Query("SELECT * FROM key_info WHERE time <= $1 ORDER BY time DESC LIMIT 1", form_time)
  if err != nil {
    return nil_key, err
  }
  var info_array []Key_info
  defer rows.Close()
  for rows.Next() {
    var info Key_info
    if err := rows.Scan(&info.Time, &info.State, &info.Key_info_id); err != nil {
      return nil_key, err
    }
    info_array = append(info_array, info)
  }
  return info_array[0], nil
}

/* app_idがDBに格納されていれば真
 * app_id: string 要求されたapp_id
*/
func Has_app_id(app_id string) (bool, error){
  rows, err := db.Query("SELECT COUNT(*) FROM app_id WHERE app_id = $1", app_id)
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
  jst := time.FixedZone("Asia/Tokyo", 9 * 60 * 60) //Hour*Minute*Second
  nowJST := now.In(jst)
  second_time := nowJST.Add(time.Second)
  time_format := nowJST.Format("2006-01-02 15:04:05 -07:00")
  second_time_format := second_time.Format("2006-01-02 15:04:05 -07:00")
  before, err_array := Get_latest_state() //get before state

  if err_array == nil {
    if state != before.State {
      // different latest state
      _, err_exec := db.Exec("INSERT INTO key_info (time, state) VALUES ($1, $2)", time_format, state)
      if err_exec == nil {
        return nil
      } else {
        return err_exec
      }
    } else {
      // same state
      if state == "ON" {
        // state ON
        _, err_exec_first := db.Exec("INSERT INTO key_info (time, state) VALUES ($1, $2)", time_format, "OFF")
        if err_exec_first == nil {
          _, err_exec_second := db.Exec("INSERT INTO key_info (time, state) VALUES ($1, $2)", second_time_format, state)
          if err_exec_second == nil {
            return nil
          } else {
            return err_exec_second
          }
        } else {
          return err_exec_first
        }
      } else {
        // state OFF
        _, err_exec_first := db.Exec("INSERT INTO key_info (time, state) VALUES ($1, $2)", time_format, "ON")
        if err_exec_first == nil {
          _, err_exec_second := db.Exec("INSERT INTO key_info (time, state) VALUES ($1, $2)", second_time_format, state)
          if err_exec_second == nil {
            return nil
          } else {
            return err_exec_second
          }
        } else {
          return err_exec_first
        }
      }
    }
  } else {
    return err_array
  }
}
