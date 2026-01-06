package http

import (
	"encoding/json"
	"net/http"

	"github.com/matlimaaa/fx-ledger-platform/services/api-gateway/internal/usecase"
)

type DepositHandler struct {
	uc *usecase.DepositFundsUseCase
}

func NewDepositHandler(
	uc *usecase.DepositFundsUseCase,
) *DepositHandler {
	return &DepositHandler{uc: uc}
}

func (h *DepositHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req struct {
		WalletID string `json:"wallet_id"`
		Amount   int64  `json:"amount"`
		Currency string `json:"currency"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	output, err := h.uc.Execute(usecase.DepositFundsInput(req))
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}
