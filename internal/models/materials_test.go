package models

import (
	"database/sql"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func NewFakeDb() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatalf("error while creating fake db: %s", err)
	}

	return db, mock
}

var m = &Material{1, "https://github.com/oleksandr-pol/simple-go-service", "test"}

func TestAllMaterials(t *testing.T) {
	db, mock := NewFakeDb()
	repo := &DB{db}

	defer db.Close()

	query := "SELECT * FROM materials"

	rows := sqlmock.NewRows([]string{"id", "url", "title"}).
		AddRow(m.Id, m.Url, m.Title)

	mock.ExpectQuery(query).WillReturnRows(rows)

	res, err := repo.AllMaterials()

	if res == nil {
		t.Error("Repository does not return materials")
	}

	if err != nil {
		t.Errorf("Error while processing method: %v", err.Error())
	}
}

func TestCreateMaterial(t *testing.T) {
	db, mock := NewFakeDb()
	repo := &DB{db}

	defer db.Close()

	sqlInsert := `
	INSERT INTO materials (url, name)
	VALUES ($1, $2)
	RETURNING id`

	rows := sqlmock.NewRows([]string{"id"}).
		AddRow(m.Id)

	mock.ExpectQuery(sqlInsert).WithArgs(m.Url, m.Title).WillReturnRows(rows)

	res, err := repo.CreateMaterial(m)

	if res == 0 {
		t.Error("Repository does not create new material")
	}

	if err != nil {
		t.Errorf("Error while processing CreateMaterial method: %v", err.Error())
	}
}
