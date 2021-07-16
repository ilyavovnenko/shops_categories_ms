package category

import (
	"database/sql"
)

type CategoryRepo struct {
	DB *sql.DB
}

func New(db *sql.DB) CategoryRepository {
	return &CategoryRepo{db}
}

func (cr *CategoryRepo) StoreCategory(cat Category) (Category, error) {
	query := `INSERT INTO categories (shop_id, shop_external_id, name)
			  VALUES(?,?,?)
			  ON DUPLICATE KEY UPDATE
			  active = ?,
			  name = ?,
			  updated_at = ?;`

	stmt, err := cr.DB.Prepare(query)
	if err != nil {
		return cat, err
	}

	// Prepared statements take up server resources and should be closed after use.
	defer stmt.Close()

	result, err := stmt.Exec(cat.ShopID, cat.ShopExternalId, cat.Name, cat.Active, cat.Name, cat.UpdatedAt)
	if err != nil {
		return cat, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return cat, err
	}

	cat.ID = uint64(lastID)

	return cat, nil
}

func (cr *CategoryRepo) StoreParentCategoryConnection(catId uint64, parentCatId uint64) error {
	return nil
}

func (ar *CategoryRepo) DeactivateShopCategories(ShopID uint) {

}
