-- +gooseUp

-- gooseStatementBegin
CREATE TABLE IF NOT EXISTS users (
    id VARCHAR NOT NULL,
    ip VARCHAR NOT NULL,
    refresh_token_hash VARCHAR NOT NULL
);
-- gooseStatementEnd

-- +gooseDown
DROP TABLE IF EXISTS users;