package repositories

import (
	"go-todo/models"

	"gorm.io/gorm"
)

type subItemRepo struct {
	Conn *gorm.DB
}

func (db *subItemRepo) ReadSubItems(pageSize, pageNumber int, filterDesc string) ([]models.SubItem, error) {
	var data []models.SubItem
	query := db.Conn.Model(&models.SubItem{})
	if filterDesc != "" {
		query = query.Where("description LIKE ?", "%"+filterDesc+"%")
	}
	query = query.Limit(pageSize).Offset((pageNumber - 1) * pageSize)
	result := query.Find(&data)
	return data, result.Error
}

func (db *subItemRepo) ReadSubItem(id int) (models.SubItem, error) {
	var data models.SubItem
	result := db.Conn.First(&data, id)
	return data, result.Error
}

func (db *subItemRepo) InsertSubItem(subItem models.SubItem) error {
	result := db.Conn.Create(&subItem)
	return result.Error
}

func (db *subItemRepo) UpdateSubItem(id int, subItem models.SubItem) error {
	result := db.Conn.Model(&models.SubItem{}).Where("id = ?", id).Updates(subItem)
	return result.Error
}

func (db *subItemRepo) DeleteSubItem(id int) error {
	result := db.Conn.Unscoped().Delete(&models.SubItem{}, id)
	return result.Error
}

// ini di export ke service, ingat, capital letter untuk export
type SubItemRepo interface {
	ReadSubItems(pageSize, pageNumber int, filterDesc string) ([]models.SubItem, error)
	ReadSubItem(id int) (models.SubItem, error)
	InsertSubItem(subItem models.SubItem) error
	UpdateSubItem(id int, subItem models.SubItem) error
	DeleteSubItem(id int) error
}

func NewSubItemRepo(conn *gorm.DB) SubItemRepo {
	return &subItemRepo{Conn: conn}
}
