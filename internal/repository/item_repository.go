package repository

import (
	"errors"
	"fmt"
	"log"

	"github.com/SuwimonFaiy23/triofarm-test/internal/model"
	"gorm.io/gorm"
)

type ItemRepository interface {
	Create(req model.Item) error
	GetLastIndexOrder(menuID int64) (float64, error)
	Update(req model.Item) error
	Delete(id int64) error
	UpdateIndex(req model.Item) error
	GetByMenuID(menuID int64) ([]model.Item, error)
	GetByID(id int64) (model.Item, error)
}

type itemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) ItemRepository {
	return &itemRepository{db: db}
}

func (r *itemRepository) Create(req model.Item) error {
	if err := r.db.Create(&req).Error; err != nil {
		log.Printf("create item error : %s", err)
		return err
	}
	return nil
}

func (r *itemRepository) GetLastIndexOrder(menuID int64) (float64, error) {
	var lastNumber float64
	if err := r.db.Model(&model.Item{}).Where("menu_id = ?", menuID).Select("index_order").Order("index_order DESC").Limit(1).Pluck("index_order", &lastNumber).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, nil
		}
		log.Println("Error:", err)
		return 0, err
	}
	return lastNumber, nil
}

func (r *itemRepository) Update(req model.Item) error {
	result := r.db.Model(&model.Item{}).Where("id = ?", req.ID).Updates(model.Item{Name: req.Name})
	if result.Error != nil {
		log.Println("Error:", result.Error)
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("update item menu with ID %d not found", req.ID)
	}
	return nil
}

func (r *itemRepository) Delete(id int64) error {
	result := r.db.Delete(&model.Item{}, &id)
	if result.Error != nil {
		log.Println("Error:", result.Error)
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("delete item menu with ID %d not found", id)
	}
	return nil
}

func (r *itemRepository) UpdateIndex(req model.Item) error {
	result := r.db.Model(&model.Item{}).Where("id = ?", req.ID).Updates(model.Item{IndexOrder: req.IndexOrder})
	if result.Error != nil {
		log.Println("Error:", result.Error)
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("update index item menu with ID %d not found", req.ID)
	}
	return nil
}

func (r *itemRepository) GetByMenuID(menuID int64) ([]model.Item, error) {
	var items []model.Item
	if err := r.db.Model(&model.Item{}).Preload("Menu").Where("menu_id = ?", menuID).Order("index_order ASC").Find(&items).Error; err != nil {
		log.Println("Error:", err)
		return items, err
	}
	return items, nil
}

func (r *itemRepository) GetByID(id int64) (model.Item, error) {
	var item model.Item
	if err := r.db.First(&item, id).Error; err != nil {
		log.Println("Error:", err)
		return item, err
	}
	return item, nil
}
