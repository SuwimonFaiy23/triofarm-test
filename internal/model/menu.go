package model

import (
	"time"

	"gorm.io/gorm"
)

type Menu struct {
	ID        int64  `gorm:"primaryKey"`
	Name      string `gorm:"size:255;not null"`
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"` // สำคัญสำหรับ Soft Delete
}

type MenuRequest struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type MenuResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
