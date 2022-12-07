package models

import "time"

type Product struct {
	ID       uint   `gorm:"primarykey"`
	Name     string `gorm:"not null;unique;type:varchar(191)"`
	Brand    string `gorm:"not null;unique;type:varchar(191)"`
	UserID   uint
	CreateAt time.Time
	UpdateAt time.Time
}
