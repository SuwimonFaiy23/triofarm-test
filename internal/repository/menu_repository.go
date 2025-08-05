package repository

import (
	"log"

	"github.com/SuwimonFaiy23/triofarm-test/internal/model"
	"gorm.io/gorm"
)

type MenuRepository interface {
	Create(req model.Menu) error
	Update(req model.Menu) error
	Delete(id int64) error
	GetAll() ([]model.Menu, error) 
	GetByID(id int64) (model.Menu, error)
}

type menuRepository struct {
	db *gorm.DB
}

func NewMenuRepository(db *gorm.DB) MenuRepository {
	return &menuRepository{db: db}
}

func (r *menuRepository) Create(req model.Menu) error {
	if err := r.db.Create(&req).Error; err != nil {
		log.Printf("create menu error : %s", err)
		return err
	}
	return nil
}

func (r *menuRepository) Update(req model.Menu) error {
	result := r.db.Model(&model.Menu{}).Where("id = ?", req.ID).Updates(model.Menu{Name: req.Name})
	if result.Error != nil {
		log.Println("Error:", result.Error)
		return result.Error
	}
	return nil
}

func (r *menuRepository) Delete(id int64) error {
	result := r.db.Delete(&model.Menu{}, &id)
	if result.Error != nil {
		log.Println("Error:", result.Error)
		return result.Error
	}
	return nil
}

func (r *menuRepository) GetAll() ([]model.Menu, error) {
	var menus []model.Menu
	if err := r.db.Model(&model.Menu{}).Find(&menus).Error; err != nil {
		log.Println("Error:", err)
		return menus, err
	}
	return menus, nil
}

func (r *menuRepository) GetByID(id int64) (model.Menu, error) {
	var menu model.Menu
	if err := r.db.First(&menu, id).Error; err != nil {
		log.Println("Error:", err)
		return menu, err
	}
	return menu, nil
}