package database

import (
	"database/sql"
	"logictest/model"
	"time"
)

type mySql struct {
	db *sql.DB
}

func NewMySql(datasource string) (Database, error) {
	db, err := sql.Open("postgres", datasource)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &Pg{db: db}, nil
}

func (ms *mySql) GetTodoByID(id int) (*model.Todo, error) {

	result := model.Todo{
		ID:          1,
		Description: "todo 1",
		CreatedAt:   time.Now(),
	}
	return &result, nil
}

func (ms *mySql) GetTodoList() ([]*model.Todo, error) {
	result := []*model.Todo{
		{
			ID:          1,
			Description: "todo 1",
			CreatedAt:   time.Now(),
		},
	}
	return result, nil
}
