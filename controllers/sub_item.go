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

type SubItemControl struct {
	subItemService services.SubItemService
	validate       validator.Validate
}

func (control SubItemControl) ReadSubItems(c echo.Context) error {
	pageSize, _ := strconv.Atoi(c.QueryParam("page_size"))
	pageNumber, _ := strconv.Atoi(c.QueryParam("page_number"))
	filterDesc := c.QueryParam("filter_desc")
	result := control.subItemService.ReadSubItems(pageSize, pageNumber, filterDesc)
	return c.JSON(http.StatusOK, result)
}
func (control SubItemControl) ReadSubItem(c echo.Context) error {
	id, _ := strconv.Atoi(c.QueryParam("id"))
	result := control.subItemService.ReadSubItem(id)

	return c.JSON(http.StatusOK, result)
}

func (control SubItemControl) InsertSubItem(c echo.Context) error {
	type payload struct {
		Description string `json:"description" validate:"required,min=1,max=1000"`
		ItemId      uint   `json:"item_id"`
	}
	validator := new(payload)

	if err := c.Bind(validator); err != nil {
		return err
	}

	result := control.subItemService.InsertSubItem(models.SubItem{
		Description: validator.Description,
		ItemId:      validator.ItemId,
	})

	return c.JSON(http.StatusOK, result)
}
func (control SubItemControl) UpdateSubItem(c echo.Context) error {
	type payload struct {
		Description string `json:"description" validate:"required,min=1,max=1000"`
		ItemId      uint   `json:"item_id"`
	}
	validator := new(payload)

	if err := c.Bind(validator); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Param("id"))
	result := control.subItemService.UpdateSubItem(id, models.SubItem{
		Description: validator.Description,
		ItemId:      validator.ItemId,
	})

	return c.JSON(http.StatusOK, result)
}
func (control SubItemControl) DeleteSubItem(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	result := control.subItemService.DeleteSubItem(id)

	return c.JSON(http.StatusNoContent, result)
}

func NewSubItemControl(db *gorm.DB) SubItemControl {
	service := services.NewSubItemService(db)
	control := SubItemControl{
		subItemService: service,
		validate:       *validator.New(),
	}

	return control
}
