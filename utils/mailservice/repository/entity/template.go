package entity

import (
	"time"

	"gorm.io/gorm"
)

type Template struct {
	ID        uint32         `json:"id" gorm:"Column:id; Type:int; PRIMARY KEY"`
	Name      string         `json:"name" gorm:"Column:name; Type:text"`
	Content   string         `json:"content" gorm:"Column:content; Type:text"`
	Variables []byte         `json:"-" gorm:"Column:variables; Type:jsonb"`
	Mails     []Mail         `json:"-" gorm:"foreignKey:TemplateID"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

type Variable struct {
	Name    string `json:"name"`
	Value   string `json:"value"`
	Type    string `json:"type"`
	Require bool   `json:"require"`
	Default string `json:"default"`
}
