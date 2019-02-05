package main

import (
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

func updateUsers(c echo.Context) error {
	var users []user
	err := um.UpdateUsers(users)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.String(http.StatusOK, "OK")
}
