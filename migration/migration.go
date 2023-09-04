package migration

import (
	"log"
	"os"

	"github.com/MiroslavZaprazny/cli-migrate/file"
)

func Create(file *file.File) {
    path, err := file.ContstructPath()
    if err != nil {
        log.Fatal(err.Error())
    }

    _, err = os.Create(path)
    if err != nil {
        log.Fatal(err.Error())
    }
}

func Up() {}

func Down() {}
