/**
 *
 * @author binglang
 */
package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
)

var (
	Db  *sql.DB
	err error
)

func init() {
	Db, err = sql.Open("sqlite3", "./dw.db")
	if err != nil {
		panic(err.Error())
	}
	err = Db.Ping()
	if err != nil {
		panic(err.Error())
	}
}

func InitTables()  {
	files, err := ioutil.ReadFile("./init.sql")
	if err != nil {
		panic(err.Error())
	}
	_, err = Db.Exec(string(files))
	if err != nil {
		panic(err.Error())
	}
}
