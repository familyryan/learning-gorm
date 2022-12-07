package models

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model        // adds ID, created_at etc.
	Username   string `gorm:"type:varchar(100);unique" json:"username"`
	Password   string `gorm:"type:varchar(100)" json:"password"`
	Previlage  string `json:"previlage,omitempty"`
}

type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
