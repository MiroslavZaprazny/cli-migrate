package database

import (
	"errors"
	"log"
	"strings"
)

type Db struct {
    Source string
    Driver string
}

func New(source string) *Db {
    driver, err := getDriverNameFromSource(source)

    if err != nil {
        log.Fatal(err.Error())
    }

    return &Db{Source: source, Driver: driver}
}

func getDriverNameFromSource(source string) (string, error) {
    index := strings.Index(source, ":")

    if index == -1 {
        return "", errors.New("No driver provided")
    } 

    return source[0:index], nil
}
