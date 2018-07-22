package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

var (

	// DB is the global reference to database.
	DB *sqlx.DB
)

// MustConnect allows connect to database. It panics if there is an error because connection is considered
// critical and the package can't work without a database connection.
func MustConnect(driverName, dataSourceName string) {
	DB = sqlx.MustOpen(driverName, dataSourceName)

	err := DB.Ping()
	if err != nil {
		panic(fmt.Sprintf("could not ping to database: %v", err))
	}

}

// ConfigureMaxIdleMaxOpen allows set max idle connections and max open connections on pooling
func ConfigureMaxIdleMaxOpen(maxIdle, maxOpen int) {
	DB.SetMaxIdleConns(maxIdle)
	DB.SetMaxOpenConns(maxOpen)
}
