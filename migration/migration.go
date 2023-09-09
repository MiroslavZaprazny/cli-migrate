package migration

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/MiroslavZaprazny/cli-migrate/database"
	"github.com/MiroslavZaprazny/cli-migrate/file"
    //TODO: add to another module?
    _ "github.com/go-sql-driver/mysql"
)

var directions = []string{"up", "down"}

func Create(filePath string) {
    for _, direction := range directions {
        file, err := file.New(filePath, fmt.Sprintf("--Write your %s migration here", direction), direction)
        if err != nil {
            log.Fatal(err)
        }
        _, err = os.Create(file.Path)
        if err != nil {
            log.Fatal(err.Error())
        }
        file.WriteContent()
    }
}

func Up(db *database.Db, query string) error {
    // openedDb, err := sql.Open(db.Driver, db.Url)
    openedDb, err := sql.Open(db.Driver, "root:root@tcp(localhost:3306)/test")

    if err != nil {
        return err
    }

    _, err = openedDb.Exec(query)

    if err != nil {
        return err
    }

    return nil
}

func Down() {}
