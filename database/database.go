package database

import (
	"errors"
	"log"
	"strings"
)

type Db struct {
    Url string
    Driver string
}

func New(source string) *Db {
    driver, err := getDriverNameFromSource(source)

    if err != nil {
        log.Fatal(err.Error())
    }

    url, err := getUrlFromSource(source)

    if err != nil {
        log.Fatal(err)
    }

    return &Db{Url: url, Driver: driver}
}

func getUrlFromSource(source string) (string, error) {
    hostIdx := strings.Index(source, "://") + 3
    dbNameIdx := strings.LastIndex(source, "/")

    if hostIdx == -1 {
        log.Fatalf("Couldn't determine host")
    }

    if dbNameIdx == -1 {
        log.Fatalf("Couldn't determine database name")
    }

    var sb strings.Builder
    for i, char := range source {
        if string(char) == "@" {
            sb.WriteString("@tcp(")
            continue
        }

        if i == dbNameIdx {
            sb.WriteString(")")
        }

        sb.WriteRune(char)
   }

   return sb.String()[hostIdx:], nil
}

func getDriverNameFromSource(source string) (string, error) {
    index := strings.Index(source, ":")

    if index == -1 {
        return "", errors.New("No driver provided")
    } 

    return source[0:index], nil
}
