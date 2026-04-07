# Insight Stack + DB

- Backend: Golang (Gin + GORM + PostgreSQL)
- Frontend: Next.js + typescript
- Database: PostgreSQL

## Setup Database

1. PostgreSQL:

```sql
CREATE DATABASE psi_tech_test;
```

Jalankan melalui folder `backend/cmd`:

```bash
cd backend/cmd
go mod tidy
go run main.go
```

default runningnya di `http://localhost:8080`.

- saat running BE maka auto migrate tabel (`users`, `companies`, `vouchers`, `auth_users`).
- seed data awal saat tabel `users` masih kosong.

## Menjalankan Frontend

```bash
cd frontend
npm install
npm run dev
```

default FE `http://localhost:3000`.

## Endpoint

- `POST /auth/register`
- `POST /auth/login`
- `POST /auth/logout`
- `POST /api/checkout`
- `GET /api/users`
- `GET /api/users/external?results=10&page=1`
