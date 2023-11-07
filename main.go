package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/MiroslavZaprazny/cli-migrate/migration"
)

func main() {
    filePath := flag.String("path", "", "Path for the migration file")
    dbSource := flag.String("database", "", "Database source name")
    flag.Parse()

    switch flag.Arg(0) {
        case "create":
            if *filePath == "" {
                //TODO: if no path is create the migration file in pwd?
                log.Fatal("Please provide a path to your migration folder when creating a migration.")
            }
                migration.Create(*filePath)
        case "up":
            if *filePath == "" {
                //TODO: if no path is provided look in the pwd?
                log.Fatal("Please provide a path to your migration folder when running this command.")
            }
            if *dbSource == "" {
                log.Fatal("Please provide your databse source.")
            }

            err := migration.Migrate(*dbSource, *filePath, "down")

            if err != nil {
                log.Fatal(err)
            }
        case "down":
            if *filePath == "" {
                //TODO: if no path is provided look in the pwd?
                log.Fatal("Please provide a path to your migration folder when running this command.")
            }
            if *dbSource == "" {
                log.Fatal("Please provide your databse source.")
            }
            migration.Migrate(*dbSource, *filePath, "down")
        default:
            fmt.Printf("Unsupported action %s", flag.Arg(0))
            os.Exit(2)
    }
}
