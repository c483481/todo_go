package repository

import (
	"github.com/c483481/todo_go/internal/contract"
	"github.com/c483481/todo_go/internal/models"
	"github.com/c483481/todo_go/pkg/handler"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type todosRepository struct {
	db *gorm.DB
}

func implTodosRepository(db *gorm.DB) contract.TodosRepository {
	return &todosRepository{
		db: db,
	}
}

func (t *todosRepository) Create(todo *models.Todos) error {
	result := t.db.Create(&todo)
	return result.Error
}

func (t *todosRepository) FindByXid(xid string) (*models.Todos, error) {
	todo := &models.Todos{}

	err := t.db.Where("xid = ?", xid).First(&todo).Error

	return todo, err
}

func (t *todosRepository) FindList(payload *handler.ListPayload) (*handler.FindResult[*models.Todos], error) {
	var result []*models.Todos
	var count int64

	// parse sort by
	order := t.parseSortBy(payload.SortBy)
	query := t.db.Order(order)

	// add limit and offset if show all is false
	if !payload.ShowAll {
		query.Limit(payload.Limit).Offset(payload.Skip)
	}

	// check if filter title has value
	if payload.Filters["title"] != "" {
		query = query.Clauses(clause.Like{
			Column: "title",
			Value:  "%" + payload.Filters["title"] + "%",
		})
	}

	err := query.Find(&result).Count(&count).Error

	return &handler.FindResult[*models.Todos]{
		Result: result,
		Count:  count,
	}, err
}

func (t *todosRepository) Update(id int64, payload *models.Todos, version int) (int64, error) {
	where := make(map[string]interface{})

	where["id"] = id
	where["version"] = version

	result := t.db.Where(where).Updates(&payload)

	return result.RowsAffected, result.Error
}

func (t *todosRepository) Delete(id int64) (int64, error) {
	result := t.db.Delete(&models.Todos{}, id)
	return result.RowsAffected, result.Error
}

func (t *todosRepository) parseSortBy(order string) string {
	var sortBy string
	switch order {
	case "createdAt-asc":
		sortBy = "\"created_at\" ASC"
	case "createdAt-desc":
		sortBy = "\"created_at\" DESC"
	case "updatedAt-asc":
		sortBy = "\"updated_at\" ASC"
	case "updatedAt-desc":
		sortBy = "\"updated_at\" DESC"
	default:
		sortBy = "\"created_at\" DESC"
	}

	return sortBy
}
