# Database Configuration

This project supports two methods for database configuration:

## Method 1: DB_LINK (Recommended)

Use a single environment variable `DB_LINK` with the full database connection string:

```bash
DB_LINK=mysql://username:password@host:port/database_name
```

### Examples:

**PostgreSQL:**
```bash
DB_LINK=postgres://user:password@localhost:5432/myapp_dev
```

**MySQL:**
```bash
DB_LINK=mysql://root:password123@localhost:3306/myapp_dev
```

**SQLite:**
```bash
DB_LINK=sqlite:./database.db
```

## Method 2: Individual Environment Variables (Fallback)

If `DB_LINK` is not provided, the system will fall back to individual environment variables:

```bash
DB_TYPE=postgres
DB_HOST=localhost
DB_PORT=5432
DB_USER=root
DB_PASSWORD=password123
DB_NAME=myapp_dev
```

## Configuration Priority

1. If `DB_LINK` is set, it will be used exclusively
2. If `DB_LINK` is not set, individual `DB_*` variables will be used
3. If neither is set, default values will be applied

## GoFrame Configuration

The database configuration is automatically loaded through:
- `config.yaml` - Uses `${DB_LINK}` environment variable
- `internal/utility/env.go` - Provides helper functions for both approaches

## CLI Tools

For GoFrame CLI tools (like `gf gen dao`), make sure to set the `DB_LINK` environment variable:

```bash
export DB_LINK=postgres://root:password@localhost:5432/myapp_dev
gf gen dao
```