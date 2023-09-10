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
        file, err := file.New(filePath, "", direction)
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
    openedDb, err := sql.Open(db.Driver, db.Url)

    if err != nil {
        return err
    }

    fmt.Printf("Executing query: %s", query)
    _, err = openedDb.Exec(query)

    if err != nil {
        return err
    }

    return nil
}

func Down() {}
