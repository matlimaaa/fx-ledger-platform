# FX Ledger Platform

FX Ledger Platform is a simplified financial backend platform designed to demonstrate
production-grade backend engineering skills using Go.

This project simulates core components of an international brokerage or fintech system,
focusing on correctness, consistency, and architectural clarity rather than feature volume.

## Core Concepts

- Multi-currency wallet
- Financial transactions and deposits
- Immutable ledger (double-entry accounting)
- Asynchronous processing via message queues
- Clean Architecture and explicit domain boundaries
- Observability and operational concerns

## Non-goals

This project intentionally avoids:
- Full matching engines
- Real banking integrations
- KYC/AML implementations
- Premature scalability optimizations

The focus is correctness, clarity, and explainability.

## Tech Stack

- Go
- PostgreSQL
- RabbitMQ (later stages)
- Docker / Docker Compose

## Status

🚧 Work in progress  
Currently implementing core domain and deposit flow.
