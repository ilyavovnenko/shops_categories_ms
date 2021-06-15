-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE categories (
    id               bigint unsigned auto_increment primary key,
    shop_id          int unsigned            not null,
    shop_external_id varchar(255)            not null,
    parent_id        bigint unsigned         null,
    active           tinyint default 1       not null,
    name             varchar(255)            not null,
    created_at       timestamp default NOW() not null,
    updated_at       timestamp default NOW() not null,

    FOREIGN KEY (shop_id) REFERENCES shops(id)
)
    collate = utf8mb4_unicode_ci;

-- todo: add foreifn key for shop_id

create index idxName on categories (name);
create index idxShopExternalId on categories (shop_external_id);
create index idxParentId on categories (parent_id);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE categories;
