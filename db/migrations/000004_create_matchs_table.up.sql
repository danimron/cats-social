CREATE TABLE IF NOT EXISTS matchs(
    id serial PRIMARY KEY,
    issuer_cat_id INT NOT NULL,
    receiver_cat_id INT NOT NULL,
    user_id INT NOT NULL,
    status VARCHAR (50) NOT NULL,
    message TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    CONSTRAINT fk_issuer_cat_id
        FOREIGN KEY(issuer_cat_id)
            REFERENCES cats(id)
            ON DELETE CASCADE,
    CONSTRAINT fk_receiver_cat_id
        FOREIGN KEY(receiver_cat_id)
            REFERENCES cats(id)
            ON DELETE CASCADE, 
    CONSTRAINT fk_user_id
        FOREIGN KEY(user_id)
            REFERENCES users(id)
            ON DELETE CASCADE
);

