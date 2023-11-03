CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       username TEXT UNIQUE NOT NULL,
                       password TEXT NOT NULL,
                       email TEXT UNIQUE NOT NULL
);

CREATE INDEX idx_username ON users (username);
CREATE INDEX idx_email ON users (email);

