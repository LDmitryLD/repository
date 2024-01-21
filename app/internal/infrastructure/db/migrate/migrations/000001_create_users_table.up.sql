CREATE TABLE IF NOT EXISTS users(
    id BIGSERIAL PRIMARY KEY,
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    username VARCHAR(100),
    email VARCHAR(100),
    address VARCHAR(100),
    deleted_at VARCHAR(100)
);