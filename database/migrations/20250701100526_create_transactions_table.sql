-- +goose Up
-- +goose StatementBegin
CREATE TYPE sync_status AS ENUM ('pending', 'synced', 'failed', 'processing');
CREATE TYPE transaction_type AS ENUM ('checkin', 'checkout', 'topup', 'refund');

CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    card_id INT NOT NULL REFERENCES cards(id),
    origin_terminal_id INTEGER REFERENCES terminals(id),
    destination_terminal_id INTEGER REFERENCES terminals(id),
    checkin_gate_id INTEGER REFERENCES gates(id),
    checkout_gate_id INTEGER REFERENCES gates(id),
    transaction_type transaction_type ,
    amount DECIMAL(10,2) NOT NULL DEFAULT 0.00,
    balance_before DECIMAL(12,2) NOT NULL,
    balance_after DECIMAL(12,2) NOT NULL,
    checkin_time TIMESTAMP WITH TIME ZONE,
    checkout_time TIMESTAMP WITH TIME ZONE,
    journey_duration INTERVAL GENERATED ALWAYS AS (checkout_time - checkin_time) STORED,
    fare_calculation JSONB, -- Store calculation details
    status sync_status NOT NULL DEFAULT 'pending',
    sync_status sync_status NOT NULL DEFAULT 'pending',
    error_message TEXT,
    processed_offline BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    synced_at TIMESTAMP WITH TIME ZONE,
    
    CONSTRAINT chk_transaction_amount CHECK (
        (transaction_type IN ('checkin', 'checkout') AND amount >= 0) OR
        (transaction_type = 'topup' AND amount > 0) OR
        (transaction_type = 'refund' AND amount >= 0)
    ),
    CONSTRAINT chk_checkout_logic CHECK (
        (transaction_type = 'checkout' AND checkin_time IS NOT NULL AND checkout_time IS NOT NULL) OR
        (transaction_type != 'checkout')
    )
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS transactions;
DROP TYPE IF EXISTS sync_status;
DROP TYPE IF EXISTS transaction_type;
-- +goose StatementEnd
