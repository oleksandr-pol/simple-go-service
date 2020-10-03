package models

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type DataStore interface {
	AllMaterials() ([]*Material, error)
	CreateMaterial(*Material) (int, error)
	UpdateMaterial(*Material) error
	DeleteMaterial(int) error
	GetMaterial(int) (*Material, error)
}

type DB struct {
	*sql.DB
}

func NewDB(dataSourceName string) (*DB, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &DB{db}, nil
}
