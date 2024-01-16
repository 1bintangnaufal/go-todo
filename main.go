package main

import (
	"go-todo/config"
	"go-todo/controllers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	db := config.Init()
	config.Migrate(db)

	api := e.Group("api/v1/")

	itemControl := controllers.NewItemControl(db)
	api.GET("items", itemControl.ReadItems)
	api.GET("item", itemControl.ReadItem)
	api.POST("item/insert", itemControl.InsertItem)
	api.PUT("item/update/:id", itemControl.UpdateItem)
	api.DELETE("item/delete/:id", itemControl.DeleteItem)

	subItemControl := controllers.NewSubItemControl(db)
	api.GET("sub_items", subItemControl.ReadSubItems)
	api.GET("sub_item", subItemControl.ReadSubItem)
	api.POST("sub_item/insert", subItemControl.InsertSubItem)
	api.PUT("sub_item/update/:id", subItemControl.UpdateSubItem)
	api.DELETE("sub_item/delete/:id", subItemControl.DeleteSubItem)

	e.Logger.Fatal(e.Start(":1323"))
}
