-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TYPE status AS ENUM ('pending', 'sent', 'failed');
CREATE TABLE notifications (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
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
