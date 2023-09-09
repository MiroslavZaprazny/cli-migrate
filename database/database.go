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

    if hostIdx == -1 {
        log.Fatalf("Couldn't determine host")
    }

    return source[hostIdx:], nil
}

func getDriverNameFromSource(source string) (string, error) {
    index := strings.Index(source, ":")

    if index == -1 {
        return "", errors.New("No driver provided")
    } 

    return source[0:index], nil
}
