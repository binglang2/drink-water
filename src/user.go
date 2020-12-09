/**
 *
 * @author binglang
 */
package main

import (
	"time"
)

type User struct {
	Id      int
	Name    string
	Mobile  string
	Sex     int
	DtToken string
	Deleted bool
	Created time.Time
	Updated time.Time
}

func AddUser(user User) error {
	sqlStr := "insert into dw_user(name, mobile, sex, dt_token) values(?, ?, ?, ?)"
	_, err := Db.Exec(sqlStr, user.Name, user.Mobile, user.Sex, user.DtToken)
	return err
}

func UpdateNameByMobile(name string, mobile string) error {
	sqlStr := "update dw_user set name=?, updated=current_timestamp where mobile = ?"
	_, err := Db.Exec(sqlStr, name, mobile)
	return err
}

func GetByMobile(mobile string) (*User, error) {
	sqlStr := "select id, name, mobile, sex, dt_token, created, updated from dw_user where mobile=?"
	row := Db.QueryRow(sqlStr, mobile)
	user := User{}
	err := row.Scan(&user.Id, &user.Name, &user.Mobile, &user.Sex, &user.DtToken, &user.Created, &user.Updated)
	return &user, err
}

func SelectUserList() ([]*User, error) {
	sqlStr := "select id, name, mobile, sex, dt_token, created, updated from dw_user where deleted=false"
	rows, err := Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	var users []*User
	for rows.Next() {
		user := &User{}
		err = rows.Scan(&user.Id, &user.Name, &user.Mobile, &user.Sex, &user.DtToken, &user.Created, &user.Updated)
		users = append(users, user)
	}
	return users, err
}

func DeleteUser(id int) error {
	sqlStr := "update dw_user set deleted=true where id=?"
	_, err := Db.Exec(sqlStr, id)
	return err
}
