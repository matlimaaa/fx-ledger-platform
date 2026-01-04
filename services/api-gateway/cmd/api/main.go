package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	httpadapter "github.com/matlimaaa/fx-ledger-platform/services/api-gateway/internal/adapters/http"
	"github.com/matlimaaa/fx-ledger-platform/services/api-gateway/internal/adapters/postgres"
	"github.com/matlimaaa/fx-ledger-platform/services/api-gateway/internal/usecase"
)

func main() {
	db, err := sql.Open(
		"postgres",
		"postgres://fx:fx@localhost:5432/fx_ledger?sslmode=disable",
	)

	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	depositRepo := postgres.NewDepositRepository(db)
	createDepositUC := usecase.NewCreateDepositUseCase(depositRepo)
	depositHandler := httpadapter.NewDepositHandler(createDepositUC)

	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	mux.HandleFunc("/deposits", depositHandler.Create)

	log.Println("api-gateway running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
