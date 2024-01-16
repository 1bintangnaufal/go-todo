package controllers

import (
	"go-todo/models"
	"go-todo/services"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ItemControl struct {
	itemService services.ItemService
	validate    validator.Validate
}

func (control ItemControl) ReadItems(c echo.Context) error {
	pageSize, _ := strconv.Atoi(c.QueryParam("page_size"))
	pageNumber, _ := strconv.Atoi(c.QueryParam("page_number"))
	filterTitle := c.QueryParam("filter_title")
	filterDesc := c.QueryParam("filter_desc")
	result := control.itemService.ReadItems(pageSize, pageNumber, filterTitle, filterDesc)
	return c.JSON(http.StatusOK, result)
}
func (control ItemControl) ReadItem(c echo.Context) error {
	id, _ := strconv.Atoi(c.QueryParam("id"))
	result := control.itemService.ReadItem(id)

	return c.JSON(http.StatusOK, result)
}

func (control ItemControl) InsertItem(c echo.Context) error {
	type payload struct {
		Title       string   `json:"title" validate:"required,min=1,max=100"`
		Description string   `json:"description" validate:"required,min=1,max=1000"`
		File        []string `json:"file"`
	}
	validator := new(payload)

	if err := c.Bind(validator); err != nil {
		return err
	}

	result := control.itemService.InsertItem(models.Item{
		Title:       validator.Title,
		Description: validator.Description,
		File:        validator.File,
	})

	return c.JSON(http.StatusOK, result)
}
func (control ItemControl) UpdateItem(c echo.Context) error {
	type payload struct {
		Title       string   `json:"title" validate:"required,min=1,max=100"`
		Description string   `json:"description" validate:"required,min=1,max=1000"`
		File        []string `json:"file"`
	}
	validator := new(payload)

	if err := c.Bind(validator); err != nil {
		return err
	}

	// walaupun aslinya integer, parameter bawaan echo itu string jadi kalau tipenya int
	// eger tetap harus di convert
	// pakai anonymous variable karena strconv expect untuk return 2 value
	id, _ := strconv.Atoi(c.Param("id"))
	result := control.itemService.UpdateItem(id, models.Item{
		Title:       validator.Title,
		Description: validator.Description,
		File:        validator.File,
	})

	return c.JSON(http.StatusOK, result)
}
func (control ItemControl) DeleteItem(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	result := control.itemService.DeleteItem(id)

	return c.JSON(http.StatusNoContent, result)
}

func NewItemControl(db *gorm.DB) ItemControl {
	service := services.NewItemService(db)
	control := ItemControl{
		itemService: service,
		validate:    *validator.New(),
	}

	return control
}
