package comerprofile

import (
	"time"
)

type ComerProfile struct {
	Id         int64     `gorm:"id,primary_key" db:"id"`
	ComerId    int64     `gorm:"comer_id" db:"comer_id"`
	Name       string    `gorm:"name" db:"name"`     // name
	Avatar     string    `gorm:"avatar" db:"avatar"` // avatar
	Cover      string    `gorm:"cover" db:"cover"`
	Location   string    `gorm:"location" db:"location"`   // location city
	TimeZone   string    `gorm:"time_zone" db:"time_zone"` // time zone: UTC-09:30
	Website    string    `gorm:"website" db:"website"`     // website
	Email      string    `gorm:"email" db:"email"`         // email
	Twitter    string    `gorm:"twitter" db:"twitter"`     // twitter
	Discord    string    `gorm:"discord" db:"discord"`     // discord
	Telegram   string    `gorm:"telegram" db:"telegram"`   // telegram
	Medium     string    `gorm:"medium" db:"medium"`       // medium
	Facebook   string    `gorm:"facebook" db:"facebook"`
	Linktree   string    `gorm:"linktree" db:"linktree"`
	Bio        string    `gorm:"bio" db:"bio"` // bio
	Languages  string    `gorm:"languages" db:"languages"`
	Educations string    `gorm:"educations" db:"educations"`
	CreatedAt  time.Time `gorm:"created_at" db:"created_at"`
	UpdatedAt  time.Time `gorm:"updated_at" db:"updated_at"`
	IsDeleted  int64     `gorm:"is_deleted" db:"is_deleted"` // Is Deleted
}

// TableName Startup table name for gorm
func (ComerProfile) TableName() string {
	return "comer_profile"
}
