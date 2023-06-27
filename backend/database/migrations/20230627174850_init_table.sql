-- +goose Up
CREATE TABLE facts (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    fact TEXT NOT NULL,
    length INTEGER NOT NULL
);


-- +goose Down
DROP TABLE facts;
