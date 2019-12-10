package product

import (
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strconv"
)

type ClothingHandler struct {
	rep Repo
}

func New(rep Repo) *ClothingHandler {
	return &ClothingHandler{rep: rep}
}

func (h ClothingHandler) Create(c echo.Context) error {
	var product = GetClothModel(c)
	if err := h.rep.Create(&product); err == nil {
		return c.String(http.StatusOK, "Ok")
	}
	return c.String(http.StatusNotAcceptable, "Failed")
}

func (h ClothingHandler) Get(c echo.Context) error {

	if cloth, err := h.rep.Get(bson.ObjectIdHex(c.QueryParam("id"))); err == nil {
		return c.XML(http.StatusOK, cloth)
	}
	return c.String(http.StatusNotFound, "Wrong id")
}

func (h ClothingHandler) GetCategory(c echo.Context) error {

	if cloth, err := h.rep.GetCategory(c.QueryParam("category")); err == nil {
		return c.XML(http.StatusOK, cloth)
	}
	return c.String(http.StatusInternalServerError, "Error")
}

func (h ClothingHandler) Delete(c echo.Context) error {
	if err := h.rep.Delete(bson.ObjectIdHex(c.QueryParam("id"))); err == nil {
		return c.String(http.StatusOK, "Ok")
	}
	return c.String(http.StatusNotFound, "Not found")
}

func (h ClothingHandler) SaveChanges(c echo.Context) error {
	var product = GetClothModel(c)
	product.Id = bson.ObjectIdHex(c.QueryParam("id"))
	if err := h.rep.Update(&product); err == nil {
		return c.String(http.StatusOK, "Ok")
	} else {
		return err
	}
}

func GetClothModel(c echo.Context) Cloth {
	var price, err = strconv.ParseFloat(c.QueryParam("price"), 32)
	if err != nil {
		price = 0;
	}
	if price < 0 {
		price = -price
	}
	return Cloth{
		Category:     c.QueryParam("category"),
		Manufacturer: c.QueryParam("manufacturer"),
		Name:         c.QueryParam("name"),
		Text:         c.QueryParam("text"),
		Season:       c.QueryParam("season"),
		Size:         c.QueryParam("size"),
		Collection:   c.QueryParam("collection"),
		Price:        price,
	}
}
