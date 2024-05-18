-- init.sql

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE transactions (
    transaction_id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    buyer_address VARCHAR(200) DEFAULT NULL,
    seller_address VARCHAR(200) DEFAULT NULL,
    post_token VARCHAR(20) NOT NULL,
    callback_url TEXT NOT NULL,
    supplier_id VARCHAR(20) NOT NULL,
    demand_id VARCHAR(20) NOT NULL,
    UNIQUE (post_token, supplier_id, demand_id)
);
