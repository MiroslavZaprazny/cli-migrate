package database

import "testing"

func TestGetDriverNameFromSource(t *testing.T) {
    tests := []struct {
        input string
        expected string
    } {
        {"mysql://root:root@localhost:3306", "mysql"},
        {"postgre://root:rootlocalhost:3306", "postgre"},
    }

    for _, test := range tests {
        driver, err := getDriverNameFromSource(test.input)
        if err != nil {
            t.Errorf("Got error %s", err.Error())
        }

        if driver != test.expected {
            t.Errorf("Expcted %s got %s", test.expected, driver)
        }
    }
}
