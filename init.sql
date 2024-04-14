CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(64) UNIQUE NOT NULL,
    firstname VARCHAR(50) NOT NULL,
    lastname VARCHAR(50) NOT NULL,
    password VARCHAR(128) NOT NULL,
    activationtoken VARCHAR(128) NOT NULL,
    activationtokenexpiration TIMESTAMP NOT NULL,
    passwordchangetoken VARCHAR(128),
    passwordchangetokenexpiration TIMESTAMP NOT NULL,
    pinnumber INTEGER,
    registerip VARCHAR(64),
    enabled INTEGER NOT NULL DEFAULT '0'
);