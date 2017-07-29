-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE users_kano (
  user_id integer unique,
  name    varchar(40),
  email   varchar(40)
);


-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS users_kano;
