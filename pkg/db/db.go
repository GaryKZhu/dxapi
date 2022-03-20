//db.go - simply db operation

package db

import (
	"database/sql"
	"log"
	_ "os"
	

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

func CheckError (err *error) {
	if *err != nil {
		log.Print((*err).Error())
		return
	}
	*err = nil
}

func GetSQLite3Connection() *sql.DB {
	db, err := sql.Open("sqlite3", "/Users/garyzhu/Projects/dxapi/conf/app.db") // Open the created SQLite File

	if err != nil {
		panic(err.Error())
	}
	return db
}

func GetData() *sql.DB {
	db, err := sql.Open("sqlite3", "/Users/garyzhu/Projects/dxapi/data/data.db") // Open the created SQLite File

	if err != nil {
		panic(err.Error())
	}
	return db
}
func RunQuery(db *sql.DB, statement string) (*sql.Rows, error) {
//	log.Println("##################### Running Query #############################")
//	log.Println(statement)
	rows, err := db.Query(statement)
	CheckError(&err)

	return rows, err
}

