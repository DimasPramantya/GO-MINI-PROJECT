-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE bioskop (
    id SERIAL PRIMARY KEY,
    nama VARCHAR(100) NOT NULL,
    lokasi VARCHAR(255) NOT NULL,
    rating REAL
);

-- +migrate StatementEnd