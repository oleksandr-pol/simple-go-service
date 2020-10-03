package models

type Material struct {
	Id    int
	Url   string
	Title string
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
