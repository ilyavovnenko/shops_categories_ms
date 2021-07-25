package mysql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/ilyavovnenko/shops_categories_ms/internal/attribute"
	"github.com/sirupsen/logrus"
)

type AttributeRepo struct {
	DB  *sql.DB
	log *logrus.Logger
}

func New(db *sql.DB, log *logrus.Logger) attribute.AttributeRepository {
	return &AttributeRepo{db, log}
}

func (ar *AttributeRepo) StoreAttribute(attr attribute.Attribute) (attribute.Attribute, error) {
	query := `INSERT INTO attributes (category_id, type, mandatory, multivalue, priority, tech_name, name, default_value, validation)
			  VALUES(?,?,?,?,?,?,?,?,?)
			  ON DUPLICATE KEY UPDATE
			  type = ?,
			  active = ?,
			  mandatory = ?,
			  multivalue = ?,
			  priority = ?,
			  name = ?,
			  default_value = ?,
			  validation = ?,
			  updated_at = ?;`

	stmt, err := ar.DB.Prepare(query)
	if err != nil {
		return attr, err
	}

	// Prepared statements take up server resources and should be closed after use.
	defer stmt.Close()

	result, err := stmt.Exec(
		attr.CategoryID,
		attr.Type,
		attr.Mandatory,
		attr.Multivalue,
		attr.Priority,
		attr.TechName,
		attr.Name,
		attr.DefaultValue,
		attr.Validation,
		attr.Type,
		attr.Active,
		attr.Mandatory,
		attr.Multivalue,
		attr.Priority,
		attr.Name,
		attr.DefaultValue,
		attr.Validation,
		attr.UpdatedAt,
	)

	if err != nil {
		return attr, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return attr, err
	}

	attr.ID = uint64(lastID)

	return attr, nil
}

func (ar *AttributeRepo) StoreAttributeValue(attrVal attribute.AttributeValue) (attribute.AttributeValue, error) {
	query := `INSERT INTO attribute_values (attribute_id, tech_name, name)
			  VALUES(?,?,?)
			  ON DUPLICATE KEY UPDATE
			  name = ?,
			  updated_at = ?;`

	stmt, err := ar.DB.Prepare(query)
	if err != nil {
		return attrVal, err
	}

	// Prepared statements take up server resources and should be closed after use.
	defer stmt.Close()

	result, err := stmt.Exec(
		attrVal.AttributeID,
		attrVal.TechName,
		attrVal.Name,
		attrVal.Name,
		attrVal.UpdatedAt,
	)

	if err != nil {
		return attrVal, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return attrVal, err
	}

	attrVal.ID = uint64(lastID)

	return attrVal, nil
}

func (ar *AttributeRepo) DeactivateShopAttributes(ShopID uint) {

}

func (ar *AttributeRepo) execQuery(ctx context.Context, query string, args ...interface{}) ([]attribute.Attribute, error) {
	rows, err := ar.DB.QueryContext(ctx, query, args...)
	if err != nil {
		ar.log.Error(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			logrus.Error(errRow)
		}
	}()

	res := []attribute.Attribute{}

	for rows.Next() {
		tmp := attribute.Attribute{}
		err = rows.Scan(
			&tmp.ID,
			&tmp.Active,
			&tmp.CategoryID,
			&tmp.Type,
			&tmp.Level,
			&tmp.Mandatory,
			&tmp.Multivalue,
			&tmp.Priority,
			&tmp.TechName,
			&tmp.Name,
			&tmp.DefaultValue,
			&tmp.Validation,
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

func (ar *AttributeRepo) GetAllAttributes(ctx context.Context, perPage uint16, page int) ([]attribute.Attribute, error) {
	query := "SELECT id, active, category_id, type, level, mandatory, multivalue, priority, tech_name, name, default_value, validation, created_at, updated_at FROM attributes LIMIT ? OFFSET ?"

	if page > 0 {
		page = page - 1
	}

	offset := page * int(perPage)

	return ar.execQuery(ctx, query, perPage, offset)
}

func (ar *AttributeRepo) GetAttributeByID(ctx context.Context, id int64) (attribute.Attribute, error) {
	query := "SELECT id, active, category_id, type, level, mandatory, multivalue, priority, tech_name, name, default_value, validation, created_at, updated_at FROM attributes WHERE id=?"

	categories, err := ar.execQuery(ctx, query, id)
	if err != nil {
		return attribute.Attribute{}, err
	}

	if len(categories) > 0 {
		return categories[0], nil
	}

	return attribute.Attribute{}, errors.New("your item is not found")
}
