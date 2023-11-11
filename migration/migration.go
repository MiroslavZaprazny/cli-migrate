package migration

import (
	"fmt"
	"io/fs"
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

func Migrate(dbSource string, filePath string, direction string) error {
    db, err := database.New(dbSource)
    if err != nil {
        log.Fatalf("Failed to create database from source %s", dbSource)
    }
    pwd, err := os.Getwd()

    if err != nil {
        log.Fatal(err)
    }
    migrationPath := fmt.Sprintf("%s/%s", pwd, filePath)
    entries, err := os.ReadDir(migrationPath)
    if err != nil {
        log.Fatal(err)
    }
    // get all migrations from db
    // create a diff from the things we have in db and in the fs
    // only loop through the diff
    //TODO: move this to a private method?
    //TODO: this only works for up migration, adjust this for down migrations aswell
    result, err := db.Query("SELECT * FROM migration_versions")
    var migrations []database.MigrationVersions
    for result.Next() {
        var migration database.MigrationVersions
        if err := result.Scan(&migration.Name); err != nil {
            break
        }
        migrations = append(migrations, migration)
    }

    //TODO: we can maybe move this to the for loop above
    var lookup map[string]struct{}
    for _, migration := range migrations {
        lookup[migration.Name] = struct{}{}
    }

    var diff []fs.DirEntry
    for _, entry := range entries {
        if _, ok := lookup[entry.Name()]; ok {
            continue
        }
        diff = append(diff, entry)
    }

    for _, entry := range diff {
        directionIdx := strings.LastIndex(entry.Name(), fmt.Sprintf("_%s", direction))
        if directionIdx == -1 {
            continue
        }
        log.Println(entry.Name())

        content, err := os.ReadFile(fmt.Sprintf("%s/%s", migrationPath, entry.Name()))
        if err != nil {
            log.Fatal(err)
        }
        query := string(content)
        fmt.Printf("Executing query: %s\n", query)
        _, err = db.Exec(query)

        if err != nil {
            log.Fatal(err)
        }
    }

    return nil
}

func Init(dbSource string) {
    db, err := database.New(dbSource)
    if err != nil {
        log.Fatalf("Failed to create db from source: %s", dbSource)        
    }

    _, err = db.Exec("CREATE TABLE IF NOT EXISTS migration_versions(version_name varchar(255)")
    if err != nil {
        log.Fatal("Failed to create migration_versions table")
    }
}

