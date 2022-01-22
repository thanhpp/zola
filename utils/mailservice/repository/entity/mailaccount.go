package entity

import (
	"time"

	"gorm.io/gorm"
)

type MailAccount struct {
	ID        uint32         `json:"id" gorm:"Column:id; Type:int; PRIMARY KEY"`
	Username  string         `json:"username" gorm:"Column:username; Type:text"`
	Password  string         `json:"password" gorm:"Column:password; Type:text"`
	SMTPHost  string         `json:"smtp_host" gorm:"Column:smtp_host; Type:text"`
	SMTPPort  string         `json:"smtp_port" gorm:"Column:smtp_port; Type:text"`
	Mails     []Mail         `json:"-" gorm:"foreignKey:AccountID"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
