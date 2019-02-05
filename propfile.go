package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

type user struct {
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
	Password string
}

// UserManager interface to handle user related tasks
type UserManager interface {
	GetUsers() ([]user, error)
	UpdateUser(oldUsername, newUsername string, newRoles []string) error
}

// NewUserManager is a constructor for UserManager
func NewUserManager() UserManager {
	return &propsFile{path: "/etc/rundeck/realm.properties"}
}

type propsFile struct {
	path string
}

func (pf *propsFile) GetUsers() ([]user, error) {
	conts, err := ioutil.ReadFile(pf.path)
	if err != nil {
		return nil, err
	}
	return parseProps(conts), nil
}

func (pf *propsFile) UpdateUser(oldUsername, newUsername string, newRoles []string) error {
	users, err := pf.GetUsers()
	if err != nil {
		return err
	}

	var currUser user
	for _, currUser = range users {
		if currUser.Username == oldUsername {
			break
		}
	}
	fmt.Println(currUser)
	return nil
}

func parseProps(conts []byte) []user {
	conts = bytes.TrimSpace(conts)
	records := bytes.Split(conts, []byte("\n"))
	users := []user{}
	for _, record := range records {
		record = bytes.TrimSpace(record)
		if !bytes.HasPrefix(record, []byte("#")) && !bytes.Equal(record, []byte("")) {
			var u user
			columns := bytes.Split(record, []byte(":"))
			roles := []string{}
			for _, role := range bytes.Split(columns[1], []byte(","))[1:] {
				roles = append(roles, string(role))
			}
			u.Username = string(columns[0])
			u.Roles = roles
			users = append(users, u)
		}
	}
	return users
}
