-- +goose Up
-- +goose StatementBegin
CREATE TABLE sync_queue (
    id SERIAL PRIMARY KEY,
    gate_id INTEGER NOT NULL REFERENCES gates(id),
    data_type VARCHAR(50) NOT NULL, -- 'TRANSACTION', 'CARD_UPDATE', 'BLACKLIST_UPDATE'
    data_payload JSONB NOT NULL,
    priority INTEGER NOT NULL DEFAULT 5, -- 1 (highest) to 10 (lowest)
    sync_status sync_status NOT NULL DEFAULT 'pending',
    retry_count INTEGER NOT NULL DEFAULT 0,
    max_retries INTEGER NOT NULL DEFAULT 3,
    error_message TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    processed_at TIMESTAMP WITH TIME ZONE,
    next_retry TIMESTAMP WITH TIME ZONE,
    
    CONSTRAINT chk_priority CHECK (priority BETWEEN 1 AND 10),
    CONSTRAINT chk_retry_count CHECK (retry_count >= 0)
);;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS sync_queue;
-- +goose StatementEnd
