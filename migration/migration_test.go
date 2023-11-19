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
            path.Join("../testdata/", "test_create_table.sql"),
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
        clearTestFolder(t)
    }
}

func TestUpMigration(t *testing.T) {
    tests := []struct {
        migrationPath string
        expectedFileNames []string
        dbInput string
        dataSource string
        driver string
        migrationContent string
        table string
    } {
        {
            path.Join("../testdata/", "test_create_table1.sql"),
            []string {
                fmt.Sprintf("%s-%s", time.Now().Format("2006-01-02-15-04-05"), "test_create_table1_up.sql"), 
                fmt.Sprintf("%s-%s", time.Now().Format("2006-01-02-15-04-05"), "test_create_table1_down.sql"),
           },
            "mysql://test:password@test-db:3306/test",
            "test:password@tcp(test-db:3306)/test",
            "mysql",
            "CREATE TABLE testing(id int); INSERT INTO `testing` VALUES(1);",
            "testing",
        },

        {
            path.Join("../testdata/", "test_create_table2_test.sql"),
            []string {
                fmt.Sprintf("%s-%s", time.Now().Format("2006-01-02-15-04-05"), "test_create_table2_test_up.sql"), 
                fmt.Sprintf("%s-%s", time.Now().Format("2006-01-02-15-04-05"), "test_create_table2_test_down.sql"),
           },
            "mysql://test:password@test-db:3306/test",
            "test:password@tcp(test-db:3306)/test",
            "mysql",
            "CREATE TABLE uptest2(id int); INSERT INTO `uptest2` VALUES(1);",
            "uptest2",
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
        prepareMigrationContent(t, test.migrationContent, projectPath, "up")
        err = Init(test.dbInput)
        if err != nil {
            t.Fatalf("Failed to init migration db table: %s", err.Error())
        }
        err = Up(test.dbInput, "../testdata/")
        if err != nil {
            t.Fatalf("Error while running migrations: %s", err.Error())
        }

        db, err := sql.Open(test.driver, test.dataSource)
        if err != nil {
            t.Fatalf("Error while connecting to testing database %s", err.Error())
        }

        var id int
        row := db.QueryRow(fmt.Sprintf("SELECT id FROM %s WHERE id = 1", test.table))
        err = row.Scan(&id)

        if err != nil {
            t.Fatalf("Error while retrieving data, %s", err.Error())
        }

        if id != 1 {
            t.Fatalf("Something went wrong and the migration was not executed")
        }

        clearTestFolder(t)
    }
}

func TestResetMigrations(t *testing.T) {

}

func TestInit(t *testing.T) {

}

func prepareMigrationContent(t *testing.T, content string, projectPath string, direction string) {
    t.Log("Preparing migration content")
    entries, err := os.ReadDir(fmt.Sprintf("%s/testdata", projectPath))
    if err != nil {
        t.Error(err)
    }
    for _, entry := range entries {
        dirIdx := strings.LastIndex(entry.Name(), fmt.Sprintf("_%s", direction))
        if dirIdx == -1 {
            continue
        }
        file, err := os.OpenFile(path.Join("../testdata/", entry.Name()), os.O_RDWR, 0644)
        if err != nil {
            t.Fatal(err)
        }
        defer file.Close()
        _, err = file.WriteAt([]byte(content), 0)
        if err != nil {
            t.Fatal(err)
        }
    }
}

func assertMigrationFilesAreCreated(t *testing.T, projectPath string, expectedFileNames []string) {
    entries, err := os.ReadDir(fmt.Sprintf("%s/testdata", projectPath))
    if err != nil {
        t.Error(err.Error())
    }

    //to make this more efficient we could consturct a map of entries in the fs
    //and then just simply lookup if we are expecting it
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

func clearTestFolder(t *testing.T) {
    err := os.RemoveAll("../testdata/")
    if err != nil {
        t.Fatalf("Failed clearing test folder: %s", err.Error())
    }

    err = os.MkdirAll("../testdata/", 0777)
    if err != nil {
        t.Fatalf("Failed creating test folder: %s", err.Error())
    }
}
