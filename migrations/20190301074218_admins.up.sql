CREATE TABLE admins(
    "id" character varying(255) PRIMARY KEY NOT NULL,
    "username" character varying(255) UNIQUE NOT NULL,
    "password_hash" character varying(255) NOT NULL
);