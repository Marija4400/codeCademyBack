package models

import (
	"gorm.io/gorm"
	"time"
)

type Common struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"createdAt" gorm:"default:now()"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"default:now()"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"default:now()"`
}

