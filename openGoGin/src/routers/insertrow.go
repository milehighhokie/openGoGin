package routers

import (
	"database/sql"
	"errors"
	"fmt"
)

// CreateRow = insert a single row
func CreateRow(request DataRow, db *sql.DB) (insertResponse int64, err error) {

	// Prepare/Exec helps prevent SQL injection
	stmt, err := db.Prepare("insert into opentable (opentablecolumn1, opentablecolumn2, opentablecolumn3, opentablecolumn4) values (?,?,?,?);")
	if err != nil {
		err = errors.New(fmt.Sprintf("Error after Prepare of insert: %s", err.Error()))
		return 0, err
	}
	inserted, err := stmt.Exec(request.Column1, request.Column2, request.Column3, request.Column4)
	if err != nil {
		err = errors.New(fmt.Sprintf("Error after Exec of insert: %s", err.Error()))
		return 0, err
	}
	insertResponse, err = inserted.RowsAffected()

	return // insertResponse, err do not need to be explictly mentioned on the last return
}
