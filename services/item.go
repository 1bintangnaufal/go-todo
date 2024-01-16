package services

import (
	"fmt"
	"go-todo/helpers"
	"go-todo/models"
	"go-todo/repositories"

	"gorm.io/gorm"
)

type itemService struct {
	itemRepo repositories.ItemRepo
}

func (service *itemService) ReadItems(pageSize, pageNumber int, filterTitle, filterDesc string) helpers.Response {
	var response helpers.Response
	data, err := service.itemRepo.ReadItems(pageSize, pageNumber, filterTitle, filterDesc)
	if err != nil {
		response.Status = 500
		response.Message = fmt.Sprint("Failed to read items: ", err.Error())
	} else {
		response.Status = 200
		response.Message = "Items read"
		response.Data = data
	}
	return response
}

func (service *itemService) ReadItem(id int) helpers.Response {
	var response helpers.Response
	data, err := service.itemRepo.ReadItem(id)
	if err != nil {
		response.Status = 500
		response.Message = fmt.Sprint("Failed to read item ", id, ": ", err.Error())
	} else {
		response.Status = 200
		response.Message = fmt.Sprint("Item", id, "read")
		response.Data = data
	}
	return response
}

func (service *itemService) InsertItem(item models.Item) helpers.Response {
	var response helpers.Response
	if err := service.itemRepo.InsertItem(item); err != nil {
		response.Status = 500
		response.Message = fmt.Sprint("Failed to insert item: ", err.Error())
	} else {
		response.Status = 201
		response.Message = "Item inserted"
		response.Data = item
	}
	return response
}

func (service *itemService) UpdateItem(id int, item models.Item) helpers.Response {
	var response helpers.Response
	if err := service.itemRepo.UpdateItem(id, item); err != nil {
		response.Status = 500
		response.Message = fmt.Sprint("Failed to update item ", id, ": ", err.Error())
	} else {
		response.Status = 200
		response.Message = "Item updated"
	}
	return response
}

func (service *itemService) DeleteItem(id int) helpers.Response {
	var response helpers.Response
	if err := service.itemRepo.DeleteItem(id); err != nil {
		response.Status = 500
		response.Message = fmt.Sprint("Failed to delete item ", id, ": ", err.Error())
	} else {
		response.Status = 204
		response.Message = fmt.Sprint("Item", id, "deleted")
	}
	return response
}

type ItemService interface {
	ReadItems(pageSize, pageNumber int, filterTitle, filterDesc string) helpers.Response
	ReadItem(id int) helpers.Response
	InsertItem(item models.Item) helpers.Response
	UpdateItem(id int, item models.Item) helpers.Response
	DeleteItem(id int) helpers.Response
}

func NewItemService(db *gorm.DB) ItemService {
	return &itemService{itemRepo: repositories.NewItemRepo(db)}
}
