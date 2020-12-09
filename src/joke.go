/**
 *
 * @author binglang
 */
package main

import (
	"strings"
)

func AddJoke(content string) error {
	sqlStr := "insert into dw_joke(content) values(?)"
	_, err := Db.Exec(sqlStr, content)
	return err
}

func SelectJokeList() ([]string, error) {
	sqlStr := "select content from dw_joke where deleted=false"
	rows, err := Db.Query(sqlStr)
	if err != nil {
		if strings.Contains(err.Error(), "no such table") {
			InitTables()
		}
		return nil, err
	}
	var contents []string
	for rows.Next() {
		content := ""
		err = rows.Scan(&content)
		contents = append(contents, content)
	}
	return contents, err
}
