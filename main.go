package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/MiroslavZaprazny/cli-migrate/database"
	"github.com/MiroslavZaprazny/cli-migrate/migration"
)

func main() {
    filePath := flag.String("path", "", "Path for the migration file")
    dbSource := flag.String("database", "", "Database source name")
    flag.Parse()

    switch flag.Arg(0) {
        case "create":
            if *filePath == "" {
                log.Fatal("Please provide a path to your migration folder when creating a migration.")
            }
                migration.Create(*filePath)
        case "up":
            db := database.New(*dbSource)
            migration.Up(db)
        case "down":
            migration.Down()
        default:
            fmt.Printf("Unsupported action %s", flag.Arg(0))
            os.Exit(2)
    }
}
