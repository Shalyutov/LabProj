CREATE TABLE orders (
    id Uuid not null,
    created_at Datetime not null,
    deleted_at Datetime null,
    PRIMARY KEY(id)
)