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

## Testing

Register :

```bash
curl -X POST http://localhost:8080/auth/register -H "Content-Type: application/json" -d "{\"username\":\"testing\",\"password\":\"12345678\"}"
```

Login :

```bash
curl -X POST http://localhost:8080/auth/login -H "Content-Type: application/json" -d "{\"username\":\"testing\",\"password\":\"12345678\"}"
```

Logout :

```bash
curl -X POST http://localhost:8080/auth/logout -H "Authorization: TOKEN_DISINI"
```

Cehckout :

```bash
curl -X POST http://localhost:8080/api/checkout -H "Content-Type: application/json" -H "Authorization: ISI_TOKEN_LOGIN" -d "{\"harga_barang\":5000000,\"voucher_code\":\"DISKON50\"}"
```

Get users internal :

```bash
curl http://localhost:8080/api/users
```

Refresh token :

```bash
curl -X POST http://localhost:8080/auth/refresh -H "Content-Type: application/json" --cookie "refresh_token=ISI_REFRESH_TOKEN_DARI_RESPON_LOGIN"
```

External API :

```bash
curl http://localhost:8080/api/users/external
```

```bash
curl "http://localhost:8080/api/users/external?results=5&page=1"
```
