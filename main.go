package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/MiroslavZaprazny/cli-migrate/file"
	"github.com/MiroslavZaprazny/cli-migrate/migration"
)

func main() {
    // dbPath := flag.String("database", "", "")
    // db := database.New(*dbPath)
    filePath := flag.String("path", "migrations/migration.sql", "Path for the migration file")
    flag.Parse()

    switch flag.Arg(0) {
        case "create":
            file := file.New(*filePath, "--Write your migration here")
            migration.Create(file)
        case "up":
            migration.Up()
        case "down":
            migration.Down()
        default:
            fmt.Printf("Unsupported action %s", flag.Arg(0))
            os.Exit(2)
    }
}
