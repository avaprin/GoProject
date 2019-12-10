package main

import (
	"courseproject/product"
	"courseproject/user"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	pr := product.NewRepo()
	ph := product.New(pr)

	e.GET("/CreateProduct", ph.Create)
	e.GET("/GetProduct", ph.Get)
	e.GET("/GetCategory", ph.GetCategory)
	e.GET("/DeleteProduct", ph.Delete)
	e.GET("/UpdateProduct", ph.SaveChanges)


	ur := user.NewUserRepo()
	uh := user.New(ur)

	e.GET("/Register", uh.Create)
	e.GET("/GetUser", uh.Get)
	e.GET("/DeleteUser", uh.Delete)
	e.GET("/SaveUserChange", uh.SaveUserChange)

	failure := e.Start(":8888")
	e.Logger.Fatal(failure)
}
