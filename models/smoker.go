package models

import (
	"time"

	"gorm.io/gorm"
)

type Smoker struct {
	ID           uint64 `gorm:"primaryKey"`
	Name         string `gorm:"size:255"`
	PatientID    string `gorm:"size:15"`
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

func SmokersCreate(name string, patientId string, healthUpdate string) *Smoker {

	entry := Smoker{Name: name, PatientID: patientId, HealthUpdate: healthUpdate}
	DB.Create(&entry)
	return &entry
}

func SmokersFind(id uint64) *Smoker {
	var smokers Smoker
	DB.Where("id = ?", id).First(&smokers)
	return &smokers
}
