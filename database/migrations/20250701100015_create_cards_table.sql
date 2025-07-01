-- +goose Up
-- +goose StatementBegin
CREATE TYPE card_status AS ENUM ('active', 'blocked', 'expired', 'suspended');

CREATE TABLE cards (
    id SERIAL PRIMARY KEY,
    card_number VARCHAR(16) UNIQUE NOT NULL,
    card_physical_id VARCHAR(32) UNIQUE NOT NULL, -- RFID/NFC physical ID
    balance DECIMAL(12,2) NOT NULL DEFAULT 0.00,
    status card_status NOT NULL DEFAULT 'active',
    issued_date TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    expiry_date TIMESTAMP WITH TIME ZONE,
    last_transaction_id INT,
    last_transaction_time TIMESTAMP WITH TIME ZONE,
    security_key VARCHAR(64) NOT NULL ,
    daily_limit DECIMAL(10,2) DEFAULT 500000.00, -- Daily spending limit
    monthly_limit DECIMAL(12,2) DEFAULT 5000000.00, -- Monthly spending limit
    owner_name VARCHAR(100),
    owner_phone VARCHAR(20),
    owner_email VARCHAR(100),
    created_at TIMESTAMP  NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP  NOT NULL DEFAULT CURRENT_TIMESTAMP,
    
    CONSTRAINT chk_balance CHECK (balance >= 0),
    CONSTRAINT chk_limits CHECK (daily_limit > 0 AND monthly_limit > 0)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS cards;
DROP TYPE IF EXISTS card_status;
-- +goose StatementEnd
