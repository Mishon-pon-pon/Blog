package dataBase

import "database/sql"

//DBQuery - wrapper for db.Query
func DBQuery(query string, callback func(*sql.Rows)) {
	db, err := sql.Open("sqlite3", "db.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	result, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer result.Close()
	callback(result)
}

//DBExec - wrapper for db.Exec
func DBExec(query string) sql.Result {
	db, err := sql.Open("sqlite3", "db.db")
	defer db.Close()
	if err != nil {
		panic(err)
	}
	result, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
	return result
}
