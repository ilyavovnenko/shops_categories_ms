-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE categories_parent_categories (
    parent_category_id bigint unsigned not null,
    category_id        bigint unsigned not null,

    FOREIGN KEY (parent_category_id) REFERENCES categories(id),
    FOREIGN KEY (category_id) REFERENCES categories(id),

    UNIQUE KEY (parent_category_id, category_id)
)
    collate = utf8mb4_unicode_ci;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE categories_parent_categories;
