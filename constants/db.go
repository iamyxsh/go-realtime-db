package constants

const UserSchema = `
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    apiKey TEXT
);`

const ProjectSchema = `
CREATE TABLE IF NOT EXISTS project (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    user_id INTEGER REFERENCES users(id),
    json_fields TEXT,
    db_name VARCHAR(255) NOT NULL UNIQUE
);`

const AuthUserSchema = `
CREATE TABLE IF NOT EXISTS users9999 (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL
);`
