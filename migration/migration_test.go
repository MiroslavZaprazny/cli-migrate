package migration

import (
	"fmt"
	"os"
	"path"
	"testing"
	"time"
)

func TestCreateMigrationFile(t *testing.T) {
  
    tests := []struct {
        input string
        expectedFileNames []string
    } {
        {
            path.Join("../", "migration/test_create_table.sql"),
            []string {
                fmt.Sprintf("%s-%s", time.Now().Format("2006-01-02-15-04-05"), "test_create_table_up.sql"), 
                fmt.Sprintf("%s-%s", time.Now().Format("2006-01-02-15-04-05"), "test_create_table_down.sql"),
           },
        },
    }

    pwd, err := os.Getwd()
    pwd = path.Join(pwd, "../")
    if err != nil {
       t.Error(err) 
    }
    for _, test := range tests {
        Create(test.input)
        entries, err := os.ReadDir(fmt.Sprintf("%s/migration", pwd))
        if err != nil {
            t.Error(err)
        }

        for _, expectedFile := range test.expectedFileNames {
            var found bool
            for _, entry := range entries {
                if expectedFile == entry.Name() {
                    found = true
                    break
                }
                found = false
            }
            if found == false {
                t.Errorf("File not found %s", expectedFile)
                t.Error(entries)
            }
        }
    }
}
