package repository

import (
	"github.com/c483481/todo_go/internal/contract"
	"gorm.io/gorm"
)

func InitRepository(db *gorm.DB) *contract.Repository {
	return &contract.Repository{
	}
}
