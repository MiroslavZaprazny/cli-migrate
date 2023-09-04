package database

import (
	"errors"
	"log"
	"strings"
)

type Db struct {
    Path string
    Driver string
}

func New(path string) *Db {
    driver, err := getDriverNameFromPath(path)

    if err != nil {
        log.Fatal(err.Error())
    }

    return &Db{Path: path, Driver: driver}
}

func getDriverNameFromPath(path string) (string, error) {
    index := strings.Index(path, ":")

    if index == -1 {
        return "", errors.New("No driver provided")
    } 

    return path[0:index], nil
}
