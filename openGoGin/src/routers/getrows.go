package routers

import (
	"database/sql"
	"errors"
	"fmt"
)

// RowListInfo = fetch all rows based on entries of column2 and 3
func RowListInfo(column2Parm string, column3Parm string, db *sql.DB) (rowList []DataRow, err error) {
	var nextRow = DataRow{}
	rows, err := db.Query("select opentablecolumn1, opentablecolumn2, opentablecolumn3, opentablecolumn4 from opentable where opentablecolumn2 like ? and opentablecolumn3 like ?;", column2Parm, column3Parm)
	if err != nil {
		err = errors.New(fmt.Sprintf("Error after Query: %s", err.Error()))
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&nextRow.Column1, &nextRow.Column2, &nextRow.Column3, &nextRow.Column4)
		if err != nil {
			err = errors.New(fmt.Sprintf("Error after Scan: %s", err.Error()))
			return nil, err
		}
		rowList = append(rowList, nextRow)
	}
	return
}
