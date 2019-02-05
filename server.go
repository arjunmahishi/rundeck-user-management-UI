package main

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/labstack/echo/middleware"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	e.Static("/", "ui")

	e.GET("/users", getUsers)
	e.PUT("/users", updateUsers)

	e.Logger.Fatal(e.Start(":3000"))
}

func getUsers(c echo.Context) error {
	conts, err := ioutil.ReadFile("/etc/rundeck/rundeck-config.properties")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.String(http.StatusOK, string(conts))
}

func updateUsers(c echo.Context) error {
	propFile, err := os.Open("/etc/rundeck/rundeck-config.properties")
	defer propFile.Close()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.String(http.StatusOK, "OK")
}
