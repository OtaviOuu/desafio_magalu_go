-- +goose Up
-- +goose StatementBegin
CREATE TABLE notifications (
    id UUID PRIMARY KEY,
    message TEXT NOT NULL,
    send_at TIMESTAMP NOT NULL,
    status status NOT NULL DEFAULT 'pending'
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE notifications;
DROP EXTENSION IF EXISTS "uuid-ossp";
-- +goose StatementEnd
