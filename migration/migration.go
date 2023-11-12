package migration

import (
	"database/sql"
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

func Up(dbSource string, filePath string) error {
    db, err := database.New(dbSource)
    if err != nil {
        return fmt.Errorf("Failed to create database from source %s", dbSource)
    }
    migrationDir, err := getMigrationDirFromPath(filePath)

    entries, err := os.ReadDir(migrationDir)
    if err != nil {
        return fmt.Errorf("Couldn't get migration entries: %s", err.Error())
    }

    result, err := db.Query("SELECT * FROM migration_versions")
    if err != nil {
        return fmt.Errorf("Couldn't get migrations from the db: %s", err.Error())
    }

    var lookup map[string]struct{}
    for result.Next() {
        var migration database.MigrationVersions
        if err := result.Scan(&migration.Name); err != nil {
            break
        }
        lookup[migration.Name]  = struct{}{} 
    }


    var diff []fs.DirEntry
    for _, entry := range entries {
        if _, ok := lookup[entry.Name()]; ok {
            continue
        }
        diff = append(diff, entry)
    }
    err = executeMigrations(db, diff, migrationDir, "up")
    if err != nil {
        return err
    }

    return nil
}

func Reset(dbSource string, filePath string) error {
    db, err := database.New(dbSource)
    if err != nil {
        log.Fatalf("Failed to create database from source %s", dbSource)
    }

    migrationDir, err := getMigrationDirFromPath(filePath)
    if err != nil {
        return err
    }

    entries, err := os.ReadDir(migrationDir)
    if err != nil {
        return err
    }
    err = executeMigrations(db, entries, migrationDir, "down")

    if err != nil {
        return err
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

func executeMigrations(db *sql.DB, entries []fs.DirEntry, migrationDir string, direction string) error {
    for _, entry := range entries {
        directionIdx := strings.LastIndex(entry.Name(), fmt.Sprintf("_%s", direction))
        if directionIdx == -1 {
            continue
        }
        content, err := os.ReadFile(fmt.Sprintf("%s/%s", migrationDir, entry.Name()))
        if err != nil {
            return fmt.Errorf("Couldn't get the contents of %s migration file: %s", entry.Name(), err.Error())
        }
        fmt.Printf("Executing query for: %s\n", entry.Name())
        query := string(content)
        _, err = db.Exec(query)

        if err != nil {
            return fmt.Errorf("Something went wrong while executing query for migration %s: %s", entry.Name(), err.Error())
        }
    }

    return nil
}

func getMigrationDirFromPath(filePath string) (string, error) {
    pwd, err := os.Getwd()

    if err != nil {
        return "", err
    }

    return fmt.Sprintf("%s/%s", pwd, filePath), nil
}

