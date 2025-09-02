package services

import (
	"errors"

	"github.com/F3AR-DEV/ApiRestGO/data/models"
	"github.com/F3AR-DEV/ApiRestGO/data/repositories"
)

func CreateItem(item *models.Item) error {
	if item.Item == "" {
		return errors.New("el nombre del item no puede estar vacío")
	}
	return repositories.CreateItem(item)
}

func GetItems() ([]models.Item, error) {
	return repositories.GetItems()
}

func GetItemByID(id uint) (*models.Item, error) {
	return repositories.GetItemByID(id)
}

func UpdateItem(id uint, newData *models.Item) (*models.Item, error) {
	item, err := repositories.GetItemByID(id)
	if err != nil {
		return nil, err
	}

	item.Item = newData.Item
	err = repositories.UpdateItem(item)
	return item, err
}

func DeleteItem(id uint) error {
	return repositories.DeleteItem(id)
}
