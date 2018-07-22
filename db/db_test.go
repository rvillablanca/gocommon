package db

import "testing"
import _ "github.com/mattn/go-sqlite3"

func testConnection(t *testing.T) {
	MustConnect("sqlite3", ":memory:")
	ConfigureMaxIdleMaxOpen(1, 1)
}
