package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/MiroslavZaprazny/cli-migrate/migrator"
)

func main() {
    dbPath := flag.String("path", "", "")
    flag.Parse()

    switch flag.Arg(0) {
        case "create":
            migrator.Create(dbPath)
        case "up":
            migrator.Up()
        case "down":
            migrator.Down()
        default:
            fmt.Printf("Unsupported action %s", flag.Arg(0))
            os.Exit(2)
    }
}
