package user

import (
	"crypto/sha256"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

type UserHandler struct {
	rep Repository
}

func New(rep Repository) *UserHandler {
	return &UserHandler{rep: rep}
}

func (u UserHandler) Create(c echo.Context) error {
	email := c.QueryParam("email")
	password := c.QueryParam("password")
	h := sha256.New()
	h.Write([]byte(password))
	usr := &User{Email: email, PasswordHash: h.Sum(nil)}
	err := u.rep.Create(usr)
	if err == nil {
		return c.String(http.StatusOK, "Ok")
	}
	return c.String(http.StatusInternalServerError, "Email not unique")
}

func (u UserHandler) Get(c echo.Context) error {
	email := c.QueryParam("email")
	password := c.QueryParam("password")
	h := sha256.New()
	h.Write([]byte(password))
	usr := &User{Email: email, PasswordHash: h.Sum(nil)}
	res, err := u.rep.Get(usr)
	if err == nil && res != nil{
		return c.XML(http.StatusOK, res)
	}
	return fmt.Errorf("not found")
}

func (u UserHandler) Delete(c echo.Context) error {
	email := c.QueryParam("email")
	err := u.rep.Delete(email)
	if err == nil {
		return c.String(http.StatusOK, "Ok")
	}
	return c.String(http.StatusNotFound, "Not found")
}

func (u UserHandler) SaveUserChange(c echo.Context) error {
	dbEmail := c.QueryParam("dbemail")
	email := c.QueryParam("email")
	password := c.QueryParam("password")

	h := sha256.New()
	h.Write([]byte(password))
	usr := &User{Email: email, PasswordHash: h.Sum(nil)}

	err := u.rep.SaveUserChange(dbEmail, usr)
	return err
}
