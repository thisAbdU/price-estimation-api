CREATE TABLE estimates (
    id SERIAL PRIMARY KEY,
    product_name TEXT NOT NULL,
    price NUMERIC(10, 2) NOT NULL,
    longitude TEXT NOT NULL,
    latitude TEXT NOT NULL
);
