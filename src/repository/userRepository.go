package repository

import (
	"github.com/jinzhu/gorm"
	"model"
)

type UserRepository interface {
	Create()
	Update()
}

type DbUserRepository struct {
	Db *gorm.DB
}

func (repository DbUserRepository) Create(user *model.User) {
	repository.Db.Create(user)
}

func (repository DbUserRepository) Update() {

}
