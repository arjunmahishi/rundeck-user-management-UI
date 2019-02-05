package main

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
	return &propsFile{path: "/etc/rundeck/rundeck-config.properties"}
}

type propsFile struct {
	path string
}

func (pf *propsFile) GetUsers() ([]user, error) {
	return nil, nil
}

func (pf *propsFile) UpdateUsers(users []user) error {
	return nil
}
