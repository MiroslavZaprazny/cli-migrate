package database

import (
	"database/sql"
	"errors"
	"fmt"
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

func (db *Db) Execute(query string) error {
    openedDb, err := sql.Open(db.Driver, db.Url)

    if err != nil {
        return err
    }

    fmt.Printf("Executing query: %s\n", query)
    _, err = openedDb.Exec(query)

    if err != nil {
        return err
    }

    return nil
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

   sb.WriteString("?multiStatements=true")

   return sb.String()[hostIdx:], nil
}

func getDriverNameFromSource(source string) (string, error) {
    index := strings.Index(source, ":")

    if index == -1 {
        return "", errors.New("No driver provided")
    } 

    return source[0:index], nil
}
