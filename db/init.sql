-- init.sql

CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    buyer_token TEXT NOT NULL,
    seller_token TEXT NOT NULL,
    buyer_address TEXT NOT NULL,
    seller_address TEXT NOT NULL,
    transaction_token TEXT NOT NULL
);
