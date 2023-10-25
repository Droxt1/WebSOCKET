CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       username TEXT UNIQUE NOT NULL,
                       password TEXT NOT NULL,
                       email TEXT UNIQUE NOT NULL
--                        room_id INTEGER REFERENCES rooms(id) ON DELETE CASCADE
);

CREATE INDEX idx_username ON users (username);
CREATE INDEX idx_email ON users (email);

-- CREATE TABLE rooms (
--                        id SERIAL PRIMARY KEY,
--                        name TEXT NOT NULL
--
--                    );