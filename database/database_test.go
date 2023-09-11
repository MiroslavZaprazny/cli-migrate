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

func TestGetUrlFromSource(t *testing.T) {
    tests := []struct {
        input string
        expected string
    } {
        {"mysql://root:root@localhost:3306/test", "root:root@tcp(localhost:3306)/test"},
        {"postgre://testUser:password@127.0.0.1:3306/test", "testUser:password@tcp(127.0.0.1:3306)/test"},
    }

    for _, test := range tests {
        url, err := getUrlFromSource(test.input)
        if err != nil {
            t.Errorf("Got error %s", err.Error())
        }

        if url != test.expected {
            t.Errorf("Expcted %s got %s", test.expected, url)
        }
    }
}
