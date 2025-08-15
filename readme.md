# Database Migration

# .env
```code
JWT_SECRET=supersecretjwtkey
```

## Cara Up Migration
Menjalankan semua migration yang belum dijalankan:

```code
export $(grep -v '^#' .env | xargs) && goose -dir ./migrations mysql "$DB_DSN" up
```

# cara down
```code
export $(grep -v '^#' .env | xargs) && goose -dir ./migrations mysql "$DB_DSN" down
```