CREATE TABLE IF NOT EXISTS deposits (
    id          UUID PRIMARY KEY,
    wallet_id   TEXT NOT NULL,
    amount      BIGINT NOT NULL,
    currency    TEXT NOT NULL,
    created_at  TIMESTAMP NOT NULL
);