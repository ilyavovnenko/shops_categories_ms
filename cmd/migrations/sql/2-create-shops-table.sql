-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE shops (
    id             int unsigned auto_increment primary key,
    shop_type_id   smallint unsigned       not null,
    name           varchar(100)            not null unique,
    created_at     timestamp default NOW() not null,
    updated_at     timestamp default NOW() not null
)
    collate = utf8mb4_unicode_ci;


-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE shops;
