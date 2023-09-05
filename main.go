package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/MiroslavZaprazny/cli-migrate/file"
	"github.com/MiroslavZaprazny/cli-migrate/migration"
)

func main() {
    filePath := flag.String("path", "", "Path for the migration file")
    flag.Parse()

    switch flag.Arg(0) {
        case "create":
            if *filePath == "" {
                log.Fatal("Please provide a path to your migration folder when creating a migration.")
            }
            file, err := file.New(*filePath, "--Write your migration here")
            if err != nil {
                log.Fatal(err)
            }
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
