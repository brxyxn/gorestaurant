package utility

import "database/sql"

func ReturnError(where_when string, err error) string {
	return logErr(where_when, err)
}

func CheckCount(rows *sql.Rows) (count int) {
	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			ReturnError("Counting rows: ", err)
		}
	}
	return count
}
