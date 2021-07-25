package mysql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/ilyavovnenko/shops_categories_ms/internal/category"
	"github.com/sirupsen/logrus"
)

type CategoryRepo struct {
	DB  *sql.DB
	log *logrus.Logger
}

func New(db *sql.DB, log *logrus.Logger) category.CategoryRepository {
	return &CategoryRepo{db, log}
}

func (cr *CategoryRepo) StoreCategory(cat category.Category) (category.Category, error) {
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

func (cr *CategoryRepo) execQuery(ctx context.Context, query string, args ...interface{}) ([]category.Category, error) {
	rows, err := cr.DB.QueryContext(ctx, query, args...)
	if err != nil {
		cr.log.Error(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			logrus.Error(errRow)
		}
	}()

	res := []category.Category{}

	for rows.Next() {
		tmp := category.Category{}
		err = rows.Scan(
			&tmp.ID,
			&tmp.Active,
			&tmp.ShopID,
			&tmp.ShopExternalId,
			&tmp.Name,
			&tmp.CreatedAt,
			&tmp.UpdatedAt,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		res = append(res, tmp)
	}
	return res, err
}

func (cr *CategoryRepo) GetAll(ctx context.Context, perPage uint16, page int) ([]category.Category, error) {
	query := "SELECT id, active, shop_id, shop_external_id, name, created_at, updated_at FROM categories LIMIT ? OFFSET ?"

	if page > 0 {
		page = page - 1
	}

	offset := page * int(perPage)

	return cr.execQuery(ctx, query, perPage, offset)
}

func (cr *CategoryRepo) GetByID(ctx context.Context, id int64) (category.Category, error) {
	query := "SELECT id, active, shop_id, shop_external_id, name, created_at, updated_at FROM categories WHERE id=?"

	categories, err := cr.execQuery(ctx, query, id)
	if err != nil {
		return category.Category{}, err
	}

	if len(categories) > 0 {
		return categories[0], nil
	}

	return category.Category{}, errors.New("your item is not found")
}
