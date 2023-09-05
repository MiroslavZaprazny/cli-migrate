package migration

import (
	"log"
	"os"

	"github.com/MiroslavZaprazny/cli-migrate/file"
)

func Create(file *file.File) {
    _, err := os.Create(file.Path)
    if err != nil {
        log.Fatal(err.Error())
    }
    file.WriteContent()
}

func Up() {}

func Down() {}
