-- init.sql

CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    buyer_address VARCHAR(200) NOT NULL,
    seller_address VARCHAR(200) NOT NULL,
    transaction_token VARCHAR(20) NOT NULL
);
