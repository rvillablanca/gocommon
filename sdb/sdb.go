package sdb

import "database/sql"

type ServiceDB struct {
	db *sql.DB
	tx *sql.Tx
}

func (sdb *ServiceDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	if sdb.db != nil {
		return sdb.db.Exec(query, args...)
	}

	return sdb.tx.Exec(query, args...)
}

func (sdb *ServiceDB) QueryRow(query string, args ...interface{}) *sql.Row {
	if sdb.db != nil {
		return sdb.db.QueryRow(query, args...)
	}

	return sdb.tx.QueryRow(query, args...)
}

func (sdb *ServiceDB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	if sdb.db != nil {
		return sdb.db.Query(query, args...)
	}

	return sdb.tx.Query(query, args...)
}

func NewFromTx(tx *sql.Tx) *ServiceDB {
	return &ServiceDB{
		db: nil,
		tx: tx,
	}
}

func NewFromDB(db *sql.DB) *ServiceDB {
	return &ServiceDB{
		db: db,
		tx: nil,
	}
}
