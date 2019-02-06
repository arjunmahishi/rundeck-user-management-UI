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

	e.Use(middleware.Logger())

	e.Static("/", "ui")

	e.GET("/users", getUsers)
	e.POST("/users", createUser)
	e.PUT("/users", updateUsers)

	e.Logger.Fatal(e.Start(":3000"))
}

func getUsers(c echo.Context) error {
	users, err := um.GetUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, users)
}

func createUser(c echo.Context) error {
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

	err = um.UpdateUser(bodyJSON.OldUsername, bodyJSON.NewUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.String(http.StatusOK, "OK")
}
