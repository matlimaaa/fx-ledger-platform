# Architecture Decision Records

## 1. Why Go?

Go was chosen due to:
- Strong support for concurrency and parallelism
- Simplicity and readability for backend systems
- First-class tooling (testing, formatting, profiling)
- Proven adoption in fintech and distributed systems

## 2. Why a Ledger-based Financial Domain?

Financial systems require:
- Deterministic behavior
- Strong data consistency
- Auditable and immutable records

Using a ledger-based model allows:
- Clear separation between commands and accounting
- Safe reprocessing of events
- Easier reasoning about balances and transactions

## 3. Monorepo Structure

A monorepo was chosen to:
- Simplify local development
- Share infrastructure configuration
- Keep architectural decisions visible in one place

Each service remains independently deployable.

## 4. PostgreSQL as Source of Truth

PostgreSQL is used as the transactional database due to:
- ACID guarantees
- Mature indexing and constraint support
- Suitability for financial data integrity

Other storage systems (e.g. Elasticsearch) are considered projections,
not sources of truth.
