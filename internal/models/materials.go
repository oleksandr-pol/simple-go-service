package models

type Material struct {
	Id    int    `json:"id,omitempty"`
	Url   string `json:"url"`
	Title string `json:"title"`
}

func (db *DB) AllMaterials() ([]*Material, error) {
	rows, err := db.Query("SELECT * FROM materials")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	materials := make([]*Material, 0)
	for rows.Next() {
		material := new(Material)
		err := rows.Scan(&material.Id, &material.Url, &material.Title)
		if err != nil {
			return nil, err
		}
		materials = append(materials, material)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return materials, nil
}

func (db *DB) CreateMaterial(m *Material) (int, error) {
	var id int

	sqlInsert := `
	INSERT INTO materials (url, name)
	VALUES ($1, $2)
	RETURNING id`

	row := db.QueryRow(sqlInsert, m.Url, m.Title)
	err := row.Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (db *DB) UpdateMaterial(m *Material) error {
	sqlUpdate := `UPDATE materials SET url=$1, name=$2 WHERE id=$3`

	_, err :=
		db.Exec(sqlUpdate,
			m.Url, m.Title, m.Id)

	return err
}

func (db *DB) DeleteMaterial(id int) error {
	_, err := db.Exec("DELETE FROM materials WHERE id=$1", id)

	return err
}

func (db *DB) GetMaterial(id int) (*Material, error) {
	var m Material
	err := db.QueryRow("SELECT url, name FROM materials WHERE id=$1",
		id).Scan(&m.Id, &m.Url, &m.Title)
	if err != nil {
		return nil, err
	}

	return &m, nil
}
