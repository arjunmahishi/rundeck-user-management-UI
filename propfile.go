package main

import (
	"bytes"
	"io/ioutil"
)

type user struct {
	username string
	roles    []string
}

// UserManager interface to handle user related tasks
type UserManager interface {
	GetUsers() ([]user, error)
	UpdateUsers(users []user) error
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
	return parseProps(conts)
}

func (pf *propsFile) UpdateUsers(users []user) error {
	return nil
}

func parseProps(conts []byte) ([]user, error) {
	records := bytes.Split(conts, []byte("\n"))
	users := []user{}
	for _, record := range records {
		if !bytes.HasPrefix(record, []byte("#")) {
			var u user
			columns := bytes.Split(record, []byte(":"))
			roles := []string{}
			for _, role := range bytes.Split(columns[1], []byte(","))[1:] {
				roles = append(roles, string(role))
			}
			u.username = string(columns[0])
			u.roles = roles
			users = append(users, u)
		}
	}
	return users, nil
}
