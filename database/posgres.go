package database

import (
	"database/sql"
	"logictest/model"
	"time"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Pg struct {
	db *sql.DB
}

func NewProgressSql(datasource string) (Database, error) {
	var db *sql.DB
	// db, err := sql.Open("postgres", datasource)
	// if err != nil {
	// 	return nil, err
	// }
	// err = db.Ping()
	// if err != nil {
	// 	return nil, err
	// }
	return &Pg{db: db}, nil
}

func (pg *Pg) GetTodoByID(id int) (*model.Todo, error) {

	result := model.Todo{
		ID:          1,
		Description: "todo 1",
		CreatedAt:   time.Now(),
	}
	return &result, nil
}

func (pg *Pg) GetTodoList() ([]*model.Todo, error) {
	result := []*model.Todo{
		{
			ID:          1,
			Description: "todo 1",
			CreatedAt:   time.Now(),
		},
	}
	return result, nil
}
