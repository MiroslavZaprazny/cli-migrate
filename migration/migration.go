package migration

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/MiroslavZaprazny/cli-migrate/database"
	"github.com/MiroslavZaprazny/cli-migrate/file"
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

func Up(db *database.Db, files []file.File) {
    openedDb, err := sql.Open(db.Driver, db.Source)

    if err != nil {
        log.Fatal(err)
    }

    for _, file := range files {
        _, err := openedDb.Exec(file.Content)
        if err != nil {
            log.Fatal(err)
        }
    }
}

func Down() {}
