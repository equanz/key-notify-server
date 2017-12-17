package sql_query

import(
  "fmt"
  "log"
  "os"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

// connection var
var db *sql.DB

func init(){
  db_local, err := sql.Open("mysql", os.Getenv("MARIA_USER") + ":" + os.Getenv("MARIA_PASS") + "@tcp(" + os.Getenv("MARIA_HOST") +":" + os.Getenv("MARIA_PORT") +")/" + os.Getenv("MARIA_DB"))
  if err == nil {
    db = db_local
  } else {
    log.Fatal(err)
  }
}

func Get_all_statistics() (*sql.Rows){
  rows, err := db.Query("SELECT * FROM key_info")
  if err != nil {
    log.Fatal(err)
  }
  return rows
}

func Get_statistics(fd string) (*sql.Rows){
  rows, err := db.Query("SELECT * FROM key_info WHERE time >= ?", fd)
    if err != nil {
    log.Fatal(err)
  }
  return rows
}

