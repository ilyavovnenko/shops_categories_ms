package attribute

import (
	"database/sql"
)

type AttributeRepo struct {
	DB *sql.DB
}

func New(db *sql.DB) AttributeRepository {
	return &AttributeRepo{db}
}

func (ar *AttributeRepo) StoreAttribute(attr Attribute) (Attribute, error) {
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

	validation := NewNullString(attr.Validation)
	defaultValue := NewNullString(attr.DefaultValue)

	result, err := stmt.Exec(
		attr.CategoryID,
		attr.Type,
		attr.Mandatory,
		attr.Multivalue,
		attr.Priority,
		attr.TechName,
		attr.Name,
		defaultValue,
		validation,
		attr.Type,
		attr.Active,
		attr.Mandatory,
		attr.Multivalue,
		attr.Priority,
		attr.Name,
		defaultValue,
		validation,
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

func (ar *AttributeRepo) StoreAttributeValue(attrVal AttributeValue) (AttributeValue, error) {
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

func NewNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}
