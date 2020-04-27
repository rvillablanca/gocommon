package sdb

import (
	"database/sql"
)

func RunTX(sdb ServiceDB, txFunc func(*sql.Tx, ...interface{}) error, params ...interface{}) (err error) {
	if sdb.TX != nil {
		err = txFunc(sdb.TX, params...)
		return err
	}

	return RunTXWith(sdb.DB, txFunc, params...)

}

func RunTXWith(db *sql.DB, txFunc func(*sql.Tx, ...interface{}) error, params ...interface{}) (err error) {
	tx, errBegin := db.Begin()
	if errBegin != nil {
		return
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	err = txFunc(tx, params...)
	return err
}
