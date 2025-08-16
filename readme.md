# Database Migration

## Migration
Migration menggunakan Goose, untuk instalasi bisa di lihat pada [Goose Docs](https://github.com/pressly/goose)

### migration up
```code
goose up
```

# migration down
```code
goose down
```

## Environtment
Environtment sebagari berikut, bisa dilihat juga di .env.example
config menyesuaikan masih-masih local

```code
DB_HOST="127.0.0.1"
DB_POST="5432"
DB_USERNAME="pgsql"
DB_PASSWORD="pgsql"
DB_NAME="database"

GOOSE_DRIVER=postgres
GOOSE_DBSTRING=postgres://pgsql:pgsql@localhost:5432/database
GOOSE_MIGRATION_DIR=./migrations
GOOSE_TABLE=public.goose_migrations
```
