package repository

import (
	"github.com/jinzhu/gorm"
)

type TodoRepository interface {
	Create()
	Update()
	Delete()
}

type DbTodoRepository struct {
	Db *gorm.DB
}

func (repository DbTodoRepository) Create() {
}

func (repository DbTodoRepository) Update() {

}

func (repository DbTodoRepository) Delete() {

}
