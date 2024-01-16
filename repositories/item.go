package repositories

import (
	"go-todo/models"

	"gorm.io/gorm"
)

type itemRepo struct {
	Conn *gorm.DB
}

func (db *itemRepo) ReadItems(pageSize, pageNumber int, filterTitle, filterDesc string) ([]models.Item, error) {
	var data []models.Item
	query := db.Conn.Model(&models.Item{})
	if filterTitle != "" {
		query = query.Where("title LIKE ?", "%"+filterTitle+"%")
	}
	if filterDesc != "" {
		query = query.Where("description LIKE ?", "%"+filterDesc+"%")
	}
	query = query.Limit(pageSize).Offset((pageNumber - 1) * pageSize)
	result := query.Preload("SubItems").Find(&data)
	return data, result.Error
}

func (db *itemRepo) ReadItem(id int) (models.Item, error) {
	var data models.Item
	result := db.Conn.Preload("SubItems").First(&data, id)
	return data, result.Error
}

func (db *itemRepo) InsertItem(item models.Item) error {
	result := db.Conn.Create(&item)
	return result.Error
}

func (db *itemRepo) UpdateItem(id int, item models.Item) error {
	result := db.Conn.Model(&models.Item{}).Where("id = ?", id).Updates(item)
	return result.Error
}

func (db *itemRepo) DeleteItem(id int) error {
	result := db.Conn.Unscoped().Delete(&models.Item{}, id)
	return result.Error
}

// ini di export ke service, ingat, capital letter untuk export
type ItemRepo interface {
	ReadItems(pageSize, pageNumber int, filterTitle string, filterDesc string) ([]models.Item, error)
	ReadItem(id int) (models.Item, error)
	InsertItem(item models.Item) error
	UpdateItem(id int, item models.Item) error
	DeleteItem(id int) error
}

func NewItemRepo(conn *gorm.DB) ItemRepo {
	return &itemRepo{Conn: conn}
}
