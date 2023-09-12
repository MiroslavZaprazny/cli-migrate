package migration

import (
	"fmt"
	"log"
	"os"
	"strings"

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

func Up(dbSource string, filePath string) error {
    db := database.New(dbSource)
    pwd, err := os.Getwd()

    if err != nil {
        log.Fatal(err)
    }
    migrationPath := fmt.Sprintf("%s/%s", pwd, filePath)
    entries, err := os.ReadDir(migrationPath)

    if err != nil {
        log.Fatal(err)
    }

    for _, entry := range entries {
        upDirection := strings.LastIndex(entry.Name(), "_up")
        downDirecrion := strings.LastIndex(entry.Name(), "_down")
        if upDirection == -1 && downDirecrion == -1 {
            continue
        }

        content, err := os.ReadFile(fmt.Sprintf("%s/%s", migrationPath, entry.Name()))
        if err != nil {
            log.Fatal(err)
        }
        err = db.Execute(string(content))

        if err != nil {
            log.Fatal(err)
        }
    }

    return nil
}

func Down() {}
