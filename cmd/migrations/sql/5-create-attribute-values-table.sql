-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE attribute_values (
    id            bigint unsigned auto_increment primary key,
    attribute_id  bigint unsigned         not null,
    tech_name     varchar(255)            not null,
    name          varchar(255)            not null,
    created_at    timestamp default NOW() not null,
    updated_at    timestamp default NOW() not null,

    FOREIGN KEY (attribute_id) REFERENCES attributes(id),

    UNIQUE KEY (attribute_id,tech_name)
)
    collate = utf8mb4_unicode_ci;

create index idxName on attribute_values (name);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE attribute_values;
