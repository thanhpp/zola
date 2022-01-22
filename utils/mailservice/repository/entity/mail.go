package entity

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Mail struct {
	ID         uint32         `json:"id" gorm:"Column:id; Type:int; PRIMARY KEY"`
	AccountID  uint32         `json:"account_id" gorm:"Column:account_id; Type:int"`
	SendTo     pq.StringArray `json:"send_to" gorm:"Column:send_to; Type:text[]"`
	Subject    string         `json:"subject" gorm:"Column:subject; Type:text"`
	TemplateID uint32         `json:"templateID" gorm:"Column:template_id; Type:int"`
	Varibles   []byte         `json:"variables" gorm:"Column:variables; Type:jsonb"`
	Status     string         `json:"status" gorm:"Column:status; Type:text"`
	CreatedAt  time.Time      `json:"created_at"`
	SentAt     time.Time      `json:"send_at"`
	DeletedAt  gorm.DeletedAt `json:"-"`
}
