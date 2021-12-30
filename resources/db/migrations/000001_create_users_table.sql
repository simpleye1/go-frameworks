-- +migrate Up
CREATE TABLE IF NOT EXISTS user_records
(
    id       serial PRIMARY KEY,
    username VARCHAR(50) UNIQUE  NOT NULL,
    password VARCHAR(50)         NOT NULL,
    email    VARCHAR(300) UNIQUE NOT NULL
);
-- +migrate Down
DROP TABLE user_records;