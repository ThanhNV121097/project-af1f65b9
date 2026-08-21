# ERD — hello-word-2

## Scope

One table stores the greeting shown on the only page. Exactly one active row supplies visible text.

## Tables

### `greetings`

| Column | Type | Null | Default | Notes |
|---|---|---|---|---|
| `id` | `smallint` | no | none | Primary key. Scaffold uses `1` for single row. |
| `text` | `text` | no | none | Stored greeting returned by API. Must not be empty. |
| `updated_at` | `timestamptz` | no | `now()` | Latest committed value wins if row is changed later. |

Constraints:

- `primary key (id)`.
- `check (id = 1)` enforces one logical greeting row.
- `check (length(text) > 0)` prevents blank greeting.

Seed data:

```sql
insert into greetings (id, text) values (1, 'Hello Word')
on conflict (id) do nothing;
```

## Relationships

No relationships. This project stores no user, session, or audit data.

## Migration contract

- Migrations live under `code/backend/internal/migrations/sql/`.
- Backend applies pending `.up.sql` files on boot in filename order.
- Applied versions are tracked in `schema_migrations`.
- `.down.sql` files exist for symmetry, but deployment rollback restores from database backup rather than automatic destructive rollback.
