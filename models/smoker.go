package models

import (
	"time"

	"gorm.io/gorm"
)

type Smoker struct {
	ID           uint64 `gorm:"primaryKey"`
	Name         string `gorm:"size:255"`
	HealthUpdate string `gorm:"type:text"`
	CreatedAt    time.Time
	UpdatedAt    time.Time      `gorm:"index"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

func SmokerAll() *[]Smoker {
	var smokers []Smoker
	DB.Where("deleted_at is NULL").Order("updated_at desc").Find(&smokers)
	return &smokers
}
