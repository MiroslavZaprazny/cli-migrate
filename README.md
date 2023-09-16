# Cli database migration tool

# Usage

## Creating a migration
```bash
$ migrate -path database/migrations/create_user_table.sql create
```

## Migrating up
```bash
$ migrate -database mysql://user:password@localhost:3306/database -path database/migrations up
```

# Tests
Run tests with 
```bash
$ docker compose build && docker compose run test go test ./...
```
