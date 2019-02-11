package main

import (
	"github.com/labstack/echo"
)

func validateUser(username, password string, c echo.Context) (bool, error) {
	users, err := um.GetUsers()
	if err != nil {
		return false, err
	}

	for _, user := range users {
		if user.Username == username && user.Password == password {
			c.Set("user", user)
			return true, nil
		}
	}
	return false, nil
}

func validateAccess(c echo.Context, scope string) bool {
	for _, role := range getCurrUser(c).Roles {
		if role == scope {
			return true
		}
	}
	return false
}

func getCurrUser(c echo.Context) user {
	return c.Get("user").(user)
}

func createAllowance(c echo.Context) []bool {
	if validateAccess(c, "admin") {
		return []bool{true, true, true}
	}
	return []bool{false, false, false}
}
