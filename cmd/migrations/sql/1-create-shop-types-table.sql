-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE shop_types (
    id   smallint unsigned auto_increment primary key,
    type varchar(100) not null unique
)
    collate = utf8mb4_unicode_ci;


-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE shop_types;
