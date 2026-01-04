package postgres

import (
	"database/sql"

	"github.com/matlimaaa/fx-ledger-platform/services/api-gateway/internal/domain"
)

type DepositRepository struct {
	db *sql.DB
}

func NewDepositRepository(db *sql.DB) *DepositRepository {
	return &DepositRepository{db: db}
}

func (r *DepositRepository) Save(d domain.Deposit) error {
	_, error := r.db.Exec(
		`INSERT INTO deposits (id, wallet_id, amount, currency, created_at)
		VALUES ($1, $2, $3, $4, $5)`,
		d.ID,
		d.WalletID,
		d.Money.Amount,
		d.Money.Currency,
		d.CreatedAt,
	)

	return error
}
