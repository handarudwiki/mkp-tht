-- +goose Up
-- +goose StatementBegin
CREATE TYPE gate_type_enum AS ENUM ('entry', 'exit', 'both');
CREATE TABLE gates (
    id SERIAL PRIMARY KEY,
    number INT NOT NULL,
    terminal_id INT REFERENCES terminals(id) ON DELETE CASCADE,
    device_serial VARCHAR(255) NOT NULL,
    ip_address VARCHAR(255) NOT NULL,
    last_sync TIMESTAMP WITH TIME ZONE,
    type gate_type_enum NOT NULL DEFAULT 'entry',
    status status NOT NULL  DEFAULT 'active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    UNIQUE (terminal_id, number)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS gates;
DROP TYPE IF EXISTS gate_type_enum;
-- +goose StatementEnd
