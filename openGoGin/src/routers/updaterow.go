package routers

import (
	"database/sql"
	"errors"
	"fmt"
)

// UpdateRow = update all rows for a given column1
func UpdateRow(rowUpdate DataRow, db *sql.DB) (updateResults int64, err error) {

	stmt, err := db.Prepare("update rowUpdateUsers set opentablecolumn2 = ?, opentablecolumn3 = ?, opentablecolumn4 = ? where opentablecolumn1=?;")
	if err != nil {
		err = errors.New(fmt.Sprintf("Error after Prepare of update: %s", err.Error()))
		return 0, err
	}
	updated, err := stmt.Exec(rowUpdate.Column2, rowUpdate.Column3, rowUpdate.Column4, rowUpdate.Column1)
	if err != nil {
		err = errors.New(fmt.Sprintf("Error after Exec of update: %s", err.Error()))
		return 0, err
	}
	updateResults, err = updated.RowsAffected()

	return // updateResults, err do not need to be explictly mentioned on the last return
}
