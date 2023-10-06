CREATE TABLE "users"
(
    "id"       bigserial PRIMARY KEY,
    "username" varchar(255) NOT NULL UNIQUE,
    "email"    varchar(255) NOT NULL,
    "password" varchar(255) NOT NULL
)
