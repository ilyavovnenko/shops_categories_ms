-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE attributes (
    id            bigint unsigned auto_increment primary key,
    category_id   bigint unsigned                                         not null,
    type          enum ('string', 'text', 'integer' , 'float', 'boolean') not null,
    level         enum ('product', 'variation')                           null,
    active        tinyint default 1                                       not null,
    mandatory     tinyint default 0                                       not null,
    multivalue    tinyint default 0                                       not null,
    priority      tinyint                                                 null,
    tech_name     varchar(255)                                            not null,
    name          varchar(255)                                            not null,
    default_value varchar(255)                                            null,
    validation    json                                                    null, 
    created_at    timestamp default NOW()                                 not null,
    updated_at    timestamp default NOW()                                 not null,

    FOREIGN KEY (category_id) REFERENCES categories(id)
)
    collate = utf8mb4_unicode_ci;

-- todo: add foreifn key for shop_id

create index idxName on attributes (name);
create index idxTechName on attributes (tech_name);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE attributes;
