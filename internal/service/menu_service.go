package service

import (
	"log"

	"github.com/SuwimonFaiy23/triofarm-test/internal/model"
	"github.com/SuwimonFaiy23/triofarm-test/internal/repository"
)

type MenuService interface {
	CreateMenu(req model.MenuRequest) error
	UpdateMenu(req model.MenuRequest) error
	DeleteMenu(id int64) error
	GetMenuList() ([]model.MenuResponse, error)
}

type menuService struct {
	menuRepo repository.MenuRepository
}

func NewMenuService(menuRepo repository.MenuRepository) MenuService {
	return &menuService{menuRepo}
}

func (s *menuService) CreateMenu(req model.MenuRequest) error {
	setData := model.Menu{
		Name: req.Name,
	}
	if err := s.menuRepo.Create(setData); err != nil {
		log.Printf("create menu error : %s", err)
		return err
	}
	return nil
}

func (s *menuService) UpdateMenu(req model.MenuRequest) error {
	setData := model.Menu{
		ID:   req.ID,
		Name: req.Name,
	}
	if err := s.menuRepo.Update(setData); err != nil {
		log.Printf("update menu error : %s", err)
		return err
	}
	return nil
}

func (s *menuService) DeleteMenu(id int64) error {
	if err := s.menuRepo.Delete(id); err != nil {
		log.Printf("delete menu error : %s", err)
		return err
	}
	return nil
}

func (s *menuService) GetMenuList() ([]model.MenuResponse, error) {
	var result []model.MenuResponse
	data, err := s.menuRepo.GetAll()
	if err != nil {
		log.Printf("get menu list error : %s", err)
		return nil, err
	}

	for _, v := range data {
		item := model.MenuResponse{
			ID:   v.ID,
			Name: v.Name,
		}
		result = append(result, item)
	}

	return result, nil
}
