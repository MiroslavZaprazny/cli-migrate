package file

import (
	"log"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
    err := os.Chdir("/tmp")
    if err != nil {
        log.Fatal(err.Error())
    }
    exitVal := m.Run()
    os.Exit(exitVal)
}

func TestConstructPath(t *testing.T) {
    tests := []struct {
        input string
        expected string
    } {
        //probably not a good idea to check against time.Now this way since time when creating the file could differ
        {"migrations/create_user_table.sql", "/tmp/migrations/" + time.Now().Format("2006-01-02-15-04-05") + "-create_user_table_up.sql"},
    }

    for _, test := range tests {
        file, err := New(test.input, "test migration", "up")

        if err != nil {
            t.Errorf("Error while constructing path %s", err.Error())
        }

        if test.expected != file.Path {
            t.Errorf("Expected path to be %s got %s", test.expected, file.Path)
        }
    }
}
