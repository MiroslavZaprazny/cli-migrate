package database

import (
	"database/sql"
	"errors"
	"log"
	"strings"
)

type MigrationVersions struct {
    Name string
}

func New(source string) (*sql.DB, error) {
    driver, err := getDriverNameFromSource(source)

    if err != nil {
        return nil, err
    }

    url, err := getUrlFromSource(source)

    if err != nil {
        return nil, err
    }

    openedDb, err := sql.Open(driver, url)

    if err != nil {
        return nil, err
    }

    return openedDb, nil
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
