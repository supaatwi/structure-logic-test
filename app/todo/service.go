package todo

import (
	"logictest/database"
	"logictest/model"
)

type InterfaceService interface {
	GetTodoByID(id int) (*model.Todo, error)
	GetTodoList() ([]*model.Todo, error)
}
type service struct {
	db database.Database
}

func NewService(db database.Database) InterfaceService {
	return &service{db}
}

func (s *service) GetTodoByID(id int) (*model.Todo, error) {

	return s.db.GetTodoByID(id)
}

func (s *service) GetTodoList() ([]*model.Todo, error) {
	return s.db.GetTodoList()
}
