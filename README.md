# authpod

env requirements:

```env
USE_ROOTUSER=<desired root user name>
USE_MIGRATION=<false for downgrade, true for upgrade, not set - skip>
POSTGRESQL_CONNSTRING=postgres://postgres:postgres@127.0.0.1:5432/postgres
LISTEN_ADDR=0.0.0.0:1234
```