package routers

import (
	"database/sql"
	"errors"
	"fmt"
)

func RowInfo(value string, db *sql.DB) (rowResponse DataRow, err error) {

	rowResponse = DataRow{}
	row := db.QueryRow("select opentablecolumn1, opentablecolumn2, opentablecolumn3, opentablecolumn4 from opentable where opentablecolumn1 = ?;", value)
	err = row.Scan(&rowResponse.Column1, &rowResponse.Column2, &rowResponse.Column3, &rowResponse.Column4)
	if err != nil {
		err = errors.New(fmt.Sprintf("Error after QueryRow and Scan: %s", err.Error()))
	}

	return
}
