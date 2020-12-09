package main

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	//testAddUser()
	m.Run()
}

func TestAddUser(t *testing.T) {
	fmt.Println("testAddUser start:")
	user := User{
		Name:    "test_user",
		Mobile:  "110",
		Sex:     0,
		DtToken: "",
	}
	err := AddUser(user)
	if err != nil {
		t.Errorf("testAddUser error: err=%s", err)
	}
}

func TestGetByMobile(t *testing.T) {
	fmt.Println("TestGetByMobile start:")
	user, err := GetByMobile("110")
	if err != nil {
		t.Errorf("TestGetByMobile error: err=%s", err)
	}
	fmt.Println("user=", user)
}

func TestDeleteUser(t *testing.T) {
	fmt.Println("TestDeleteUser start:")
	err := DeleteUser(1)
	if err != nil {
		t.Errorf("TestDeleteUser error: err=%s", err)
	}
}

func TestSelectUserList(t *testing.T) {
	fmt.Println("TestSelectUserList start:")
	users, err := SelectUserList()
	if err != nil {
		t.Errorf("TestSelectUserList error: err=%s", err)
	}
	fmt.Println("users=", users)
}
