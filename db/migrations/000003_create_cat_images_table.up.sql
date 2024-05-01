CREATE TABLE IF NOT EXISTS cat_images(
    id serial PRIMARY KEY,
    cat_id INT NOT NULL,
    image_url VARCHAR (255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    CONSTRAINT fk_cat_id
        FOREIGN KEY(cat_id)
            REFERENCES cats(id)
            ON DELETE CASCADE
);

