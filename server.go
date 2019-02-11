package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/middleware"

	"github.com/labstack/echo"
)

var um = NewUserManager()

func main() {
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} method=${method} uri=${uri} status=${status}\n",
	}))

	e.Use(middleware.BasicAuth(validateUser))

	e.Static("/", "ui")
	e.GET("/users", getUsers)
	e.POST("/users", createUser)
	e.PUT("/users", updateUsers)
	e.DELETE("/users", deleteUser)

	e.GET("/logout", func(c echo.Context) error {
		return c.String(http.StatusUnauthorized, "Logged out")
	})

	e.Logger.Fatal(e.Start(":4180"))
}

func getUsers(c echo.Context) error {
	users, err := um.GetUsers()
	currUser := getCurrUser(c)
	for i, u := range users {
		if u.Username == currUser.Username {
			temp := users[0]
			users[0] = users[i]
			users[i] = temp
		}
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"users": users, "allowance": createAllowance(c)})
}

func createUser(c echo.Context) error {
	if !validateAccess(c, "admin") {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorised"})
	}

	body := c.Request().Body
	defer body.Close()

	raw, err := ioutil.ReadAll(body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	var newUser user
	json.Unmarshal(raw, &newUser)
	err = um.CreateUser(newUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.String(http.StatusOK, "OK")
}

func updateUsers(c echo.Context) error {
	body := c.Request().Body
	defer body.Close()

	raw, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	var bodyJSON struct {
		OldUsername string
		NewUser     user
	}
	json.Unmarshal(raw, &bodyJSON)

	if !validateAccess(c, "admin") && getCurrUser(c).Username != bodyJSON.OldUsername {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorised"})
	}

	err = um.UpdateUser(bodyJSON.OldUsername, bodyJSON.NewUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.String(http.StatusOK, "OK")
}

func deleteUser(c echo.Context) error {
	if !validateAccess(c, "admin") {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorised"})
	}

	body := c.Request().Body
	defer body.Close()

	raw, err := ioutil.ReadAll(body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	var bodyJSON struct {
		Username string
	}
	json.Unmarshal(raw, &bodyJSON)

	err = um.DeleteUser(bodyJSON.Username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.String(http.StatusOK, "OK")
}
