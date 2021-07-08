package shop

import (
	"database/sql"
)

type ShopRepo struct {
	DB *sql.DB
}

func New(db *sql.DB) ShopRepository {
	return &ShopRepo{db}
}

func (sr *ShopRepo) GetIdByName(name string) (uint, error) {
	var shopID uint
	query := `SELECT id FROM shops WHERE name = ?;`

	stmt, err := sr.DB.Prepare(query)
	if err != nil {
		return shopID, err
	}

	// Prepared statements take up server resources and should be closed after use.
	defer stmt.Close()

	result := stmt.QueryRow(name)
	if err != nil {
		return shopID, err
	}

	err = result.Scan(&shopID)

	return shopID, err
}
