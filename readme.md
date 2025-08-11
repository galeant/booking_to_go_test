# Database Migration

# .env
```code
DB_DSN=root:@tcp(127.0.0.1:3307)/go_api?parseTime=true
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