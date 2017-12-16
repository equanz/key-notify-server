package sql_query

import(
  "fmt"
  "os"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

// connection var
var db = create_connection()


func create_connection(){
  db, err := sql.Open("mysql", os.Getenv("MARIA_USER") + ":" + os.Getenv("MARIA_PASS") + "@tcp(" + os.Getenv("MARIA_HOST") +":" + os.Getenv("MARIA_PORT") +")/" + os.Getenv("MARIA_DB"))
  if err == nil {
    return db
  } else {
    log.Fatal(err)
  }
}

func get_all_statistics(){
  rows, err := db.Query("SELECT * FROM key_info")
  if err != nil {
    log.Fatal(err)
  }
  return rows
}

func get_statistics(fd String){
  rows, err := db.Query("SELECT * FROM key_info WHERE time >= ?", fd)
    if err != nil {
    log.Fatal(err)
  }
  return rows
}

