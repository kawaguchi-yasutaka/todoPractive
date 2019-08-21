package repository

import "github.com/jinzhu/gorm"

type UserRepository interface {
	Create()
	Update()
}

type DbUserRepository struct {
	Db *gorm.DB
}

func (repository DbUserRepository) Create() {

}

func (repository DbUserRepository) Update() {

}
