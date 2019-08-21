package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"model"
)

func initDb() *gorm.DB {
	db, err := gorm.Open("mysql", "master:gagagigu123@tcp(localhost:3306)/todoPractice?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	db.AutoMigrate(&model.User{}, &model.Todo{})
	return db
}

func Init(db *gorm.DB) {
	//todoRepository := repository.DbTodoRepository{Db: db}
	//userRepository := repository.DbUserRepository{Db: db}
}

func main() {
	db := initDb()
	Init(db)
	defer db.Close()
	fmt.Println("hello world")
}
