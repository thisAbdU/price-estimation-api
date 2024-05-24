CREATE TABLE IF NOT EXISTS estimates (
    id SERIAL PRIMARY KEY,
    product_name TEXT NOT NULL,
    price NUMERIC NOT NULL,
    longitude NUMERIC,
    latitude NUMERIC
);
