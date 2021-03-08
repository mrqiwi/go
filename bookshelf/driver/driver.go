package driver

import (
    "os"
    "log"
    "database/sql"

    "github.com/lib/pq"
)


func ConnectDB() *sql.DB {
    var db *sql.DB
    pgURL, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
    if err != nil {
        log.Fatal(err)
    }

    db, err = sql.Open("postgres", pgURL)
    if err != nil {
        log.Fatal(err)
    }

    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    return db
}
