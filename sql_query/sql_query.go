/* Package: SQLクエリ発行，処理
*/
package sql_query

// import
import(
  "log"
  "os"
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
func Get_all_statistics() ([]Key_info){
  rows, err := db.Query("SELECT * FROM key_info")
  if err != nil {
    log.Fatal(err)
  }
  var info_array []Key_info
  defer rows.Close()
  for rows.Next() {
    var info Key_info
    if err := rows.Scan(&info.Time, &info.State, &info.Key_info_id); err != nil {
      log.Fatal(err)
    }
    info_array = append(info_array, info)
  }

  return info_array
}

/* 統計データ指定日時から最新まで抽出
 * fd: string(DATETIMEフォーマット) 取得する統計値の指定日時
*/
func Get_statistics(fd string) ([]Key_info){
  rows, err := db.Query("SELECT * FROM key_info WHERE time >= ?", fd)
    if err != nil {
    log.Fatal(err)
  }
  var info_array []Key_info
  defer rows.Close()
  for rows.Next() {
    var info Key_info
    if err := rows.Scan(&info.Time, &info.State, &info.Key_info_id); err != nil {
      log.Fatal(err)
    }
    info_array = append(info_array, info)
  }
  return info_array
}

