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

func New(path string, content string) *File {
    return &File{Path: path, Content: content}
}

func (f *File) ContstructPath() (string, error) {
    extIdx := strings.Index(f.Path, ".")
    if extIdx == -1 {
        return "", errors.New("Couldn't determine the file extension")
    }

    lastSlashIdx := strings.LastIndex(f.Path, "/")
    if lastSlashIdx == -1 {
        return "", errors.New("Couldn't determine the file name")
    }

    pwd, err := os.Getwd()

    if err != nil {
        log.Fatal(err.Error())
    }

    return fmt.Sprintf(
        "%s/%s/%s-%s%s", 
        pwd,
        f.Path[0:lastSlashIdx], 
        time.Now().Format("2006-01-02-15-04-05"), 
        f.Path[lastSlashIdx + 1:extIdx], 
        f.Path[extIdx:]), 
    nil
}
