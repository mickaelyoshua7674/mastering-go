package main
import (
    "database/sql"
    "fmt"
    "os"

    _"modernc.org/sqlite"
)

func main() {
    db, err := sql.Open("sqlite", "test.db")
    if err != nil {
        fmt.Println("Error connecting:", err)
        return
    }
    defer db.Close()

    var version string
    err = db.QueryRow("SELECT SQLITE_VERSION()").Scan(&version)
    if err != nil {
        fmt.Println("Version error:", err)
        return
    }

    fmt.Println("SQLite3 version:", version)
    os.Remove("test.db")
}

