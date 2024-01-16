package services

import (
	"fmt"
	"go-todo/helpers"
	"go-todo/models"
	"go-todo/repositories"

	"gorm.io/gorm"
)

type subItemService struct {
	subItemRepo repositories.SubItemRepo
}

func (service *subItemService) ReadSubItems(pageSize, pageNumber int, filterDesc string) helpers.Response {
	var response helpers.Response
	data, err := service.subItemRepo.ReadSubItems(pageSize, pageNumber, filterDesc)
	if err != nil {
		response.Status = 500
		response.Message = fmt.Sprint("Failed to read sub items: ", err.Error())
	} else {
		response.Status = 200
		response.Message = "Sub items read"
		response.Data = data
	}
	return response
}

func (service *subItemService) ReadSubItem(id int) helpers.Response {
	var response helpers.Response
	data, err := service.subItemRepo.ReadSubItem(id)
	if err != nil {
		response.Status = 500
		response.Message = fmt.Sprint("Failed to read sub item ", id, ": ", err.Error())
	} else {
		response.Status = 200
		response.Message = fmt.Sprint("Sub item", id, "read")
		response.Data = data
	}
	return response
}

func (service *subItemService) InsertSubItem(subItem models.SubItem) helpers.Response {
	var response helpers.Response
	if err := service.subItemRepo.InsertSubItem(subItem); err != nil {
		response.Status = 500
		response.Message = fmt.Sprint("Failed to insert sub item: ", err.Error())
	} else {
		response.Status = 201
		response.Message = "Sub item inserted"
	}
	return response
}

func (service *subItemService) UpdateSubItem(id int, subItem models.SubItem) helpers.Response {
	var response helpers.Response
	if err := service.subItemRepo.UpdateSubItem(id, subItem); err != nil {
		response.Status = 500
		response.Message = fmt.Sprint("Failed to update sub item ", id, ": ", err.Error())
	} else {
		response.Status = 200
		response.Message = "Sub item updated"
	}
	return response
}

func (service *subItemService) DeleteSubItem(id int) helpers.Response {
	var response helpers.Response
	if err := service.subItemRepo.DeleteSubItem(id); err != nil {
		response.Status = 500
		response.Message = fmt.Sprint("Failed to delete sub item ", id, ": ", err.Error())
	} else {
		response.Status = 204
		response.Message = fmt.Sprint("Sub item", id, "deleted")
	}
	return response
}

type SubItemService interface {
	ReadSubItems(pageSize, pageNumber int, filterDesc string) helpers.Response
	ReadSubItem(id int) helpers.Response
	InsertSubItem(subItem models.SubItem) helpers.Response
	UpdateSubItem(id int, subItem models.SubItem) helpers.Response
	DeleteSubItem(id int) helpers.Response
}

func NewSubItemService(db *gorm.DB) SubItemService {
	return &subItemService{subItemRepo: repositories.NewSubItemRepo(db)}
}
