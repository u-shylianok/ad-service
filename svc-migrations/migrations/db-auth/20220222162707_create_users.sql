-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE users
(
    id          SERIAL PRIMARY KEY,
    name        TEXT NOT NULL,
    username    TEXT NOT NULL,
    password    TEXT NOT NULL
);
INSERT INTO users(name, username, password) VALUES
    ('Тестовый пользователь', 'test', '$2a$10$1hN6TfPRPS9usxbx9DVoY.ix6a8o.kxsednj6CPTkHujR2JGbvLXG'); -- u: test, p: test
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE users;
-- +goose StatementEnd
