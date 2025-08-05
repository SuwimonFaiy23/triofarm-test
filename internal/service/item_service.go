package service

import (
	"log"
	"time"

	"github.com/SuwimonFaiy23/triofarm-test/internal/model"
	"github.com/SuwimonFaiy23/triofarm-test/internal/repository"
)

type ItemService interface {
	CreateItem(req model.ItemRequest) error
	UpdateItem(req model.ItemRequest) error
	DeleteItem(id int64) error
	UpdateIndexItemMenu(req model.UpdateItemRequest) error
	GetItemListByMenuID(menuID int64) (model.ItemResponse, error)
}

type itemService struct {
	itemRepo repository.ItemRepository
	menuRepo repository.MenuRepository
}

func NewItemService(itemRepo repository.ItemRepository, menuRepo repository.MenuRepository) ItemService {
	return &itemService{itemRepo, menuRepo}
}

func (s *itemService) CreateItem(req model.ItemRequest) error {
	// check menu id
	if _, err := s.menuRepo.GetByID(req.MenuID); err != nil {
		log.Printf("get menu ID error : %s", err)
		return err
	}

	// check last index
	lastIndex, err := s.itemRepo.GetLastIndexOrder(req.MenuID)
	if err != nil {
		log.Printf("create item menu error : %s", err)
		return err
	}
	index := lastIndex + 1
	setData := model.Item{
		MenuID:     req.MenuID,
		Name:       req.Name,
		IndexOrder: &index,
	}
	if err := s.itemRepo.Create(setData); err != nil {
		log.Printf("create item menu error : %s", err)
		return err
	}
	return nil
}

func (s *itemService) UpdateItem(req model.ItemRequest) error {
	setData := model.Item{
		ID:   req.ID,
		Name: req.Name,
	}
	if err := s.itemRepo.Update(setData); err != nil {
		log.Printf("update item menu error : %s", err)
		return err
	}
	return nil
}

func (s *itemService) DeleteItem(id int64) error {
	if err := s.itemRepo.Delete(id); err != nil {
		log.Printf("delete item menu error : %s", err)
		return err
	}
	return nil
}

func (s *itemService) UpdateIndexItemMenu(req model.UpdateItemRequest) error {
	var beforeIndex, afterIndex float64
	newIndex := 0.0
	if req.BeforeItemID != nil {
		resBeforeData, err := s.itemRepo.GetByID(*req.BeforeItemID)
		if err != nil {
			log.Printf("update index item menu error : %s", err)
			return err
		}
		beforeIndex = float64(*resBeforeData.IndexOrder)
	}

	if req.AfterItemID != nil {
		resAfterData, err := s.itemRepo.GetByID(*req.AfterItemID)
		if err != nil {
			log.Printf("update index item menu error : %s", err)
			return err
		}
		afterIndex = float64(*resAfterData.IndexOrder)
	}

	switch {
	case req.BeforeItemID != nil && req.AfterItemID != nil:
		newIndex = (beforeIndex + afterIndex) / 2
	case req.BeforeItemID != nil:
		newIndex = beforeIndex - 1
	case req.AfterItemID != nil:
		newIndex = afterIndex + 1
	default:
		newIndex = float64(time.Now().UnixNano())
	}

	setData := model.Item{
		ID:         req.ItemID,
		IndexOrder: &newIndex,
	}
	if err := s.itemRepo.UpdateIndex(setData); err != nil {
		log.Printf("update index item menu error : %s", err)
		return err
	}
	return nil
}

func (s *itemService) GetItemListByMenuID(menuID int64) (model.ItemResponse, error) {
	var itemList []model.ItemList
	var menuName string
	var result model.ItemResponse

	data, err := s.itemRepo.GetByMenuID(menuID)
	if err != nil {
		log.Printf("get menu list error : %s", err)
		return result, err
	}

	for i, v := range data {
		menuName = v.Menu.Name
		item := model.ItemList{
			ID:    v.ID,
			Index: i + 1,
			Name:  v.Name,
		}
		itemList = append(itemList, item)
	}
	result = model.ItemResponse{
		MenuID:   menuID,
		MenuName: menuName,
		ItemList: itemList,
	}

	return result, nil
}
