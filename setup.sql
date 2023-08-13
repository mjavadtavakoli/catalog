CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    price INT,
    category_id INT,
    weight FLOAT,
    title VARCHAR(255),
    image VARCHAR(255),
    pdf VARCHAR(255),
    description TEXT
);

CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY, 
    title VARCHAR(255),
    image VARCHAR(255)
);
 