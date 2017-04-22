CREATE TABLE users(
    id SERIAL,
    username VARCHAR(25) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    password VARCHAR(60) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,

    PRIMARY KEY (id)
);

INSERT INTO users (username, email, first_name, last_name, created_at, updated_at) VALUES
    ('system_user', 'system_user@localhost', 'System', 'User', DATE_TRUNC('second', CURRENT_TIMESTAMP), DATE_TRUNC('second', CURRENT_TIMESTAMP));
