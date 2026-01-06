package postgres

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/matlimaaa/fx-ledger-platform/services/api-gateway/internal/domain"
)

type LedgerRepository struct {
	tx *sql.Tx
}

func NewLedgerRepository(tx *sql.Tx) *LedgerRepository {
	return &LedgerRepository{tx: tx}
}

func (r *LedgerRepository) Save(entry domain.LedgerEntry) error {
	_, err := r.tx.Exec(
		`INSERT INTO ledger_entries
		(id, wallet_id, amount, currency, direction, reference_id, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		uuid.New(),
		entry.WalletID,
		entry.Amount,
		entry.Currency,
		entry.Direction,
		entry.ReferenceID,
		entry.CreatedAt,
	)

	return err
}
