package postgres

import (
	"database/sql"
	"encoding/json"

	"github.com/google/uuid"
)

type OutboxRepository struct {
	tx *sql.Tx
}

func NewOutboxRepository(tx *sql.Tx) *OutboxRepository {
	return &OutboxRepository{tx: tx}
}

func (r *OutboxRepository) Save(
	aggregate string,
	aggregateID string,
	eventType string,
	payload any,
) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	_, err = r.tx.Exec(
		`INSERT INTO outbox_events
		(id, aggregate, aggregate_id, event_type, payload, created_at)
		VALUES ($1, $2, $3, $4, $5, now())`,
		uuid.New(),
		aggregate,
		aggregateID,
		eventType,
		data,
	)

	return err
}
