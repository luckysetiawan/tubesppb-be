-- name: create-users-table
CREATE TABLE users (
    id  INTEGER PRIMARY KEY,
    username VARCHAR(255),
    email VARCHAR(255),
    password VARCHAR(255)
)