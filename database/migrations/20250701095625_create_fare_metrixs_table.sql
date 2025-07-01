-- +goose Up
-- +goose StatementBegin
CREATE TABLE fare_matrixs (
    id SERIAL PRIMARY KEY,
    origin_terminal_id INTEGER NOT NULL REFERENCES terminals(id),
    destination_terminal_id INTEGER NOT NULL REFERENCES terminals(id),
    base_fare DECIMAL(8,2) NOT NULL,
    distance_km DECIMAL(6,2) NOT NULL,
    travel_time_minutes INTEGER,
    peak_hour_multiplier DECIMAL(3,2) DEFAULT 1.00,
    weekend_multiplier DECIMAL(3,2) DEFAULT 1.00,
    effective_date TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    expiry_date TIMESTAMP WITH TIME ZONE,
    status status NOT NULL DEFAULT 'active',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    
    CONSTRAINT chk_base_fare CHECK (base_fare > 0),
    CONSTRAINT chk_distance CHECK (distance_km >= 0),
    CONSTRAINT chk_different_terminals CHECK (origin_terminal_id != destination_terminal_id),
    UNIQUE(origin_terminal_id, destination_terminal_id, effective_date)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS fare_matrixs;
-- +goose StatementEnd
