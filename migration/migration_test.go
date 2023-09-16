package migration

import (
	"database/sql"
	"fmt"
	"os"
	"path"
	"strings"
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
    projectPath := path.Join(pwd, "../")
    if err != nil {
       t.Error(err) 
    }
    for _, test := range tests {
        Create(test.input)
        assertMigrationFilesAreCreated(t, projectPath, test.expectedFileNames)
    }
}

func TestUpMigration(t *testing.T) {
    tests := []struct {
        migrationPath string
        expectedFileNames []string
        dbInput string
        dataSource string
        driver string
    } {
        {
            path.Join("../", "migration/test_create_table.sql"),
            []string {
                fmt.Sprintf("%s-%s", time.Now().Format("2006-01-02-15-04-05"), "test_create_table_up.sql"), 
                fmt.Sprintf("%s-%s", time.Now().Format("2006-01-02-15-04-05"), "test_create_table_down.sql"),
           },
            "mysql://test:password@test-db:3306/test",
            "test:password@tcp(test-db:3306)/test",
            "mysql",
        },
    }

    pwd, err := os.Getwd()
    projectPath := path.Join(pwd, "../")
    if err != nil {
       t.Error(err) 
    }

    for _, test := range tests {
        Create(test.migrationPath)
        assertMigrationFilesAreCreated(t, projectPath, test.expectedFileNames)

        entries, err := os.ReadDir(fmt.Sprintf("%s/migration", projectPath))
        if err != nil {
            t.Error(err)
        }

        for _, entry := range entries {
            upIndex := strings.LastIndex(entry.Name(), "_up")
            if upIndex == -1 {
                continue
            }
            file, err := os.OpenFile(entry.Name(), os.O_RDWR, 0644)
            if err != nil {
                t.Error(err)
            }
            defer file.Close()
            _, err = file.WriteAt([]byte("CREATE TABLE testing(id int); INSERT INTO `testing` VALUES(1);"), 0)
            if err != nil {
                t.Error(err)
            }
        }
        Up(test.dbInput, path.Join(test.migrationPath, "../"))

        db, err := sql.Open(test.driver, test.dataSource)
        if err != nil {
            t.Errorf("Error while connecting to testing database %s", err.Error())
        }

        var id int
        row := db.QueryRow("SELECT id FROM `testing` WHERE id = 1")
        err = row.Scan(&id)

        if err != nil {
            t.Errorf("Error while retrieving data, %s", err.Error())
        }

        if id != 1 {
            t.Errorf("Something went wrong and the migration was not executed")
        }
    }
}

func assertMigrationFilesAreCreated(t *testing.T, projectPath string, expectedFileNames []string) {
    entries, err := os.ReadDir(fmt.Sprintf("%s/migration", projectPath))
    if err != nil {
        t.Error(err.Error())
    }

    for _, expectedFile := range expectedFileNames {
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

