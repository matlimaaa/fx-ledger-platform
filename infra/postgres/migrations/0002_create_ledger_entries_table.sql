CREATE TABLE ledger_entries (
    id           UUID PRIMARY KEY,
    wallet_id    TEXT NOT NULL,
    amount       BIGINT NOT NULL,
    currency     TEXT NOT NULL,
    direction    TEXT NOT NULL, -- credit | debit
    reference_id UUID NOT NULL, -- deposit_id
    created_at   TIMESTAMP NOT NULL
);