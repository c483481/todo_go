package models

import "time"

type Todos struct {
	ID          int64     `gorm:"column:id;primaryKey;autoIncrement;not null;<-create"`
	Xid         string    `gorm:"column:xid;not null;unique;<-create"`
	Version     int       `gorm:"column:version;not null;default:1"`
	UpdatedAt   time.Time `gorm:"column:updated_at;not null;autoCreateTime;autoUpdateTime"`
	CreatedAt   time.Time `gorm:"column:created_at;not null;autoCreateTime;<-create"`
	Title       string    `gorm:"column:title;not null"`
	Description string    `gorm:"column:description;not null"`
}

func (t *Todos) TableName() string {
	return "todos"
}
