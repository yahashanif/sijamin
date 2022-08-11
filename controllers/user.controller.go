package controllers

import (
	"net/http"
	"rest-api-sijamin/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

func LoginAuditor(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	result, err := models.LoginAuditor(username, password)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, result)
}

func ProdiYangAuditor(c echo.Context) error {
	idAuditor := c.FormValue("idAuditor")
	idAuditorInt, _ := strconv.Atoi(idAuditor)
	result, err := models.ProdiYangAuditor(idAuditorInt)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, result)

}
