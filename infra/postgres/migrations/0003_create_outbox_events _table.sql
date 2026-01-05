CREATE TABLE outbox_events (
    id          UUID PRIMARY KEY,
    aggregate   TEXT NOT NULL,
    aggregate_id UUID NOT NULL,
    event_type  TEXT NOT NULL,
    payload     JSONB NOT NULL,
    created_at  TIMESTAMP NOT NULL,
    processed   BOOLEAN DEFAULT FALSE
);