package file

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type File struct {
    Content string
    Path string
}

func New(path string, content string, direction string) (*File, error) {
    constructedPath, err := contstructPath(path, direction)
    if err != nil {
        return nil, err
    }

    return &File{Path: constructedPath, Content: content}, nil
}

func contstructPath(path string, direction string) (string, error) {
    extIdx := strings.LastIndex(path, ".")
    if extIdx == -1 {
        return "", errors.New(fmt.Sprintf("Couldn't determine the file extension for %s", path))
    }

    lastSlashIdx := strings.LastIndex(path, "/")
    if lastSlashIdx == -1 {
        return "", errors.New("Couldn't determine the file name")
    }

    pwd, err := os.Getwd()

    if err != nil {
        log.Fatal(err.Error())
    }

    return fmt.Sprintf(
        "%s/%s/%s-%s_%s%s",
        pwd,
        path[0:lastSlashIdx],
        time.Now().Format("2006-01-02-15-04-05"),
        path[lastSlashIdx + 1:extIdx],
        direction,
        path[extIdx:]),
    nil
}

func (f *File) WriteContent() {
    err := os.WriteFile(f.Path, []byte(f.Content), 0664)
    if err != nil {
        log.Fatal(err.Error())
    }
}
