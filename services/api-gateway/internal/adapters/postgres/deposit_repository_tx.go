package postgres

import (
	"database/sql"

	"github.com/matlimaaa/fx-ledger-platform/services/api-gateway/internal/domain"
)

type DepositRepositoryTx struct {
	tx *sql.Tx
}

func NewDepositRepositoryTx(tx *sql.Tx) *DepositRepositoryTx {
	return &DepositRepositoryTx{tx: tx}
}

func (r *DepositRepositoryTx) Save(d domain.Deposit) error {
	_, err := r.tx.Exec(
		`INSERT INTO deposits (id, wallet_id, amount, currency, created_at)
		 VALUES ($1, $2, $3, $4, $5)`,
		d.ID,
		d.WalletID,
		d.Money.Amount,
		d.Money.Currency,
		d.CreatedAt,
	)

	return err
}
