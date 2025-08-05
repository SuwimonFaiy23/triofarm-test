package model

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	ID         int64    `gorm:"primaryKey"`
	MenuID     int64    `gorm:"not null"`
	Name       string   `gorm:"size:255;not null"`
	IndexOrder *float64 `gorm:"not null;index"`
	CreatedAt  time.Time
	UpdatedAt  *time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"` // สำคัญสำหรับ Soft Delete
	Menu       Menu           `gorm:"foreignKey:MenuID;constraint;"`
}

type ItemRequest struct {
	ID     int64  `json:"id"`
	MenuID int64  `json:"menu_id"`
	Index  int64  `json:"index"`
	Name   string `json:"name"`
}

type ItemResponse struct {
	MenuID   int64      `json:"menu_id"`
	MenuName string     `json:"menu_name"`
	ItemList []ItemList `json:"item_list"`
}

type ItemList struct {
	ID    int64  `json:"id"`
	Index int    `json:"index"`
	Name  string `json:"name"`
}

type UpdateItemRequest struct {
	ItemID       int64  `json:"item_id"`
	BeforeItemID *int64 `json:"before_item_id"`
	AfterItemID  *int64 `json:"after_item_id"`
}
