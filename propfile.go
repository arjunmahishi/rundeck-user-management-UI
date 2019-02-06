package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
)

type user struct {
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
	Password string   `json:"password"`
}

// UserManager interface to handle user related tasks
type UserManager interface {
	GetUsers() ([]user, error)
	CreateUser(newUser user) error
	UpdateUser(oldUsername string, newUser user) error
	DeleteUser(username string) error
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

func (pf *propsFile) CreateUser(newUser user) error {
	newRoleString := ""
	for _, r := range newUser.Roles {
		newRoleString += strings.TrimSpace(r) + ","
	}
	newRoleString = strings.Trim(newRoleString, ",")
	newUserString := fmt.Sprintf("%s:%s,%s", newUser.Username, newUser.Password, newRoleString)

	conts, err := ioutil.ReadFile(pf.path)
	if err != nil {
		return err
	}

	for _, line := range bytes.Split(conts, []byte("\n")) {
		line = bytes.TrimSpace(line)
		if bytes.HasPrefix(line, []byte(newUser.Username)) {
			return fmt.Errorf("User already exists")
		}
	}

	conts = bytes.Join([][]byte{
		conts,
		[]byte(newUserString),
	}, []byte("\n"))

	err = ioutil.WriteFile(pf.path, conts, 0)
	if err != nil {
		return err
	}
	return nil
}

func (pf *propsFile) UpdateUser(oldUsername string, newUser user) error {
	newRoleString := ""
	for _, r := range newUser.Roles {
		newRoleString += strings.TrimSpace(r) + ","
	}
	newRoleString = strings.Trim(newRoleString, ",")
	newUserString := fmt.Sprintf("%s:%s,%s", newUser.Username, newUser.Password, newRoleString)

	conts, err := ioutil.ReadFile(pf.path)
	if err != nil {
		return err
	}

	for _, line := range bytes.Split(conts, []byte("\n")) {
		line = bytes.TrimSpace(line)
		if bytes.HasPrefix(line, []byte(oldUsername)) {
			conts = bytes.Replace(conts, line, []byte(newUserString), -1)
			break
		}
	}

	err = ioutil.WriteFile(pf.path, conts, 0)
	if err != nil {
		return err
	}
	return nil
}

func (pf *propsFile) DeleteUser(username string) error {
	conts, err := ioutil.ReadFile(pf.path)
	if err != nil {
		return err
	}

	for _, line := range bytes.Split(conts, []byte("\n")) {
		line = bytes.TrimSpace(line)
		if bytes.HasPrefix(line, []byte(username)) {
			conts = bytes.Replace(conts, line, []byte(""), -1)
			break
		}
	}

	err = ioutil.WriteFile(pf.path, conts, 0)
	if err != nil {
		return err
	}
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
			userProps := bytes.Split(columns[1], []byte(","))
			for _, role := range userProps[1:] {
				roles = append(roles, string(role))
			}
			u.Username = string(columns[0])
			u.Roles = roles
			u.Password = string(userProps[0])
			users = append(users, u)
		}
	}
	return users
}
