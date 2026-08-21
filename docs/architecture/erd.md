# ERD — hello-word-2

## Scope

One table stores the greeting shown on the only page. Exactly one active row supplies visible text.

## Story extension — Show stored greeting

Reviewed UI mock contract: `code/frontend/lib/mock/show-stored-greeting.ts` exports `GreetingResponse = { text: string }`. Schema supplies that field from `greetings.text`; no schema change beyond existing `greetings` table is needed.

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

Indexes:

- Primary key index on `id` serves `GET /api/greeting` lookup where `id = 1`; no extra index needed.

Seed data:

```sql
insert into greetings (id, text) values (1, 'Hello Word')
on conflict (id) do nothing;
```

## Relationships

No relationships. This project stores no user, session, or audit data.

## Migration plan — Show stored greeting

Forward:

1. Create `greetings` table if it does not exist with `id smallint not null primary key`, `text text not null`, `updated_at timestamptz not null default now()`, `check (id = 1)`, and `check (length(text) > 0)`.
2. Seed row `(1, 'Hello Word')` with `on conflict (id) do nothing`.

Backward:

1. Delete seed row where `id = 1`.
2. Drop `greetings` table if it exists.

Safety on populated tables:

- Forward migration is safe on empty or already-migrated database because table creation and seed use idempotent guards.
- Forward migration is not meant to preserve pre-existing nonconforming `greetings` data; no such data belongs to this project before this story.
- Backward migration is destructive for greeting data; deployment rollback restores database backup rather than automatic destructive rollback.

## Migration contract

- Migrations live under `code/backend/internal/migrations/sql/`.
- Backend applies pending `.up.sql` files on boot in filename order.
- Applied versions are tracked in `schema_migrations`.
- `.down.sql` files exist for symmetry, but deployment rollback restores from database backup rather than automatic destructive rollback.
