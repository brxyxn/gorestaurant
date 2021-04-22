package utility

import "fmt"

// Instead of using panic to abort the program we'll use log the error
// showing details of the error on console.
func logErr(where_when string, err error) string {
	// eg.: logErr('Where or when the error occured', err)
	if err != nil {
		result := fmt.Sprintf(where_when, err)
		return result
	}
	return ""
}
