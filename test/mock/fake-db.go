package mock

import "github.com/oleksandr-pol/simple-go-service/internal/models"

type FakeDB struct{}

func (db *FakeDB) AllMaterials() ([]*models.Material, error) {
	return []*models.Material{
		&models.Material{Id: 1, Url: "url", Title: "title"},
		&models.Material{Id: 1, Url: "url1", Title: "title1"},
	}, nil
}

func (db *FakeDB) CreateMaterial(m *models.Material) (int, error) {
	return 1, nil
}

func (db *FakeDB) UpdateMaterial(m *models.Material) error {
	return nil
}

func (db *FakeDB) DeleteMaterial(id int) error {
	return nil
}

func (db *FakeDB) GetMaterial(id int) (*models.Material, error) {
	return &models.Material{Id: 1, Url: "url", Title: "title"}, nil
}
