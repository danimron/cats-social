CREATE TABLE IF NOT EXISTS sessions(
    id serial PRIMARY KEY,
    token VARCHAR (300) NOT NULL,
    created_at TIMESTAMP NOT NULL
);

