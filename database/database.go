package database

import "logictest/model"

type Database interface {
	GetTodoByID(id int) (*model.Todo, error)
	GetTodoList() ([]*model.Todo, error)
}
