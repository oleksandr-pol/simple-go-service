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

var m = &Material{3, "https://github.com/oleksandr-pol/simple-go-service", "test"}

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

func TestUpdateMaterial(t *testing.T) {
	db, mock := NewFakeDb()
	repo := &DB{db}
	defer db.Close()

	sqlUpdate := `UPDATE materials SET url=$1, name=$2 WHERE id=$3`
	mock.ExpectExec(sqlUpdate).WithArgs(m.Url, m.Title, m.Id).WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.UpdateMaterial(m)

	if err != nil {
		t.Errorf("Repository does not update material: %v", err.Error())
	}
}

func TestDeleteMaterial(t *testing.T) {
	db, mock := NewFakeDb()
	repo := &DB{db}
	defer db.Close()

	mock.ExpectExec("DELETE FROM materials WHERE id=$1").WithArgs(m.Id).WillReturnResult(sqlmock.NewResult(1, 1))
	err := repo.DeleteMaterial(m.Id)

	if err != nil {
		t.Errorf("Repository does not delete material: %v", err.Error())
	}
}

func TestGetMaterial(t *testing.T) {
	db, mock := NewFakeDb()
	repo := &DB{db}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "url", "title"}).
		AddRow(m.Id, m.Url, m.Title)

	mock.ExpectQuery("SELECT url, name FROM materials WHERE id=$1").WithArgs(m.Id).WillReturnRows(rows)

	res, err := repo.GetMaterial(m.Id)

	if res == nil {
		t.Error("Repository does not return material")
	}

	if err != nil {
		t.Errorf("Error while processing GetMaterial: %v", err.Error())
	}
}
