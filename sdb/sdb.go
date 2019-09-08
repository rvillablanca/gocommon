package sdb

import "database/sql"

type ServiceDB struct {
	DB *sql.DB
	TX *sql.Tx
}

func (sdb *ServiceDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	if sdb.DB != nil {
		return sdb.DB.Exec(query, args...)
	}

	return sdb.TX.Exec(query, args...)
}

func (sdb *ServiceDB) QueryRow(query string, args ...interface{}) *sql.Row {
	if sdb.DB != nil {
		return sdb.DB.QueryRow(query, args...)
	}

	return sdb.TX.QueryRow(query, args...)
}

func (sdb *ServiceDB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	if sdb.DB != nil {
		return sdb.DB.Query(query, args...)
	}

	return sdb.TX.Query(query, args...)
}

func NewFromTx(tx *sql.Tx) *ServiceDB {
	return &ServiceDB{
		DB: nil,
		TX: tx,
	}
}

func NewFromDB(db *sql.DB) *ServiceDB {
	return &ServiceDB{
		DB: db,
		TX: nil,
	}
}
