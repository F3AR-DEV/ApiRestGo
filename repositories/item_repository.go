package repositories

import (
	"github.com/F3AR-DEV/ApiRestGO/config/db"
	"github.com/F3AR-DEV/ApiRestGO/models"
)

func CreateItem(item *models.Item) error {
	return db.DB.Create(item).Error
}

func GetItems() ([]models.Item, error) {
	var items []models.Item
	err := db.DB.Find(&items).Error
	return items, err
}

func GetItemByID(id uint) (*models.Item, error) {
	var item models.Item
	err := db.DB.First(&item, id).Error
	return &item, err
}

func UpdateItem(item *models.Item) error {
	return db.DB.Save(item).Error
}

func DeleteItem(id uint) error {
	return db.DB.Delete(&models.Item{}, id).Error
}
