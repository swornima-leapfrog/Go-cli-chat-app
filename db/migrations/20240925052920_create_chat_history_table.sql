-- +goose Up
-- +goose StatementBegin

CREATE TABLE chat_histories (
    id SERIAL PRIMARY KEY,
    sender_id INTEGER NOT NULL,
    message TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_sender_id FOREIGN KEY (sender_id) REFERENCES users(id) ON DELETE CASCADE
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS chat_histories;

-- +goose StatementEnd
