-- migrate:up
CREATE TABLE users (id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, name TEXT, email TEXT);
DELETE FROM users;

-- migrate:down
DROP TABLE IF EXISTS users;

