CREATE TABLE IF NOT EXISTS cats(
    id serial PRIMARY KEY,
    user_id INT NOT NULL,
    name VARCHAR (50) NOT NULL,
    race VARCHAR (50) NOT NULL,
    sex VARCHAR (10) NOT NULL,
    age_in_months INT NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    CONSTRAINT fk_user_id
        FOREIGN KEY(user_id)
            REFERENCES users(id)
            ON DELETE CASCADE
);

