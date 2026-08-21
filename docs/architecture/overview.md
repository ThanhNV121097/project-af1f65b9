# Architecture Overview — hello-word-2

## 1. Project shape

| Area | Decision |
|---|---|
| Shape | fullstack |
| UI | Next.js 15 App Router, TypeScript, Tailwind v3 |
| API | Go 1.22 HTTP service |
| Database | PostgreSQL 16 |
| Runtime | `docker compose up` from repository root |

Purpose: prove one end-to-end path. One stored greeting row in PostgreSQL is read by Go API and rendered by Next.js page.

## 2. Repository layout

```text
code/
  backend/
    cmd/api/main.go
    internal/migrations/migrations.go
    internal/migrations/sql/*.sql
    go.mod
    .env.example
  frontend/
    app/layout.tsx
    app/page.tsx
    app/globals.css
    package.json
    .env.example
docs/
  architecture/overview.md
  architecture/erd.md
  architecture/services.md
  home/SRS.md
```

Committed container files stay as project contract:

- Root `docker-compose.yml` builds `code/backend` and `code/frontend`.
- `code/backend/Dockerfile` builds one Go binary.
- `code/frontend/Dockerfile` runs Next standalone output.
- `.github/workflows/ci.yml` runs build, vet, tests, lint, token checks.

## 3. Runtime data flow

1. Browser opens frontend.
2. Frontend story component later calls `NEXT_PUBLIC_API_URL`.
3. Backend handles `GET /api/greeting`.
4. Backend reads current greeting from PostgreSQL.
5. Backend returns JSON response.
6. Frontend renders returned text or minimal error state.

No greeting text may be hardcoded in frontend code. Mock data may exist only in story UI branch and must be removed when API lands.

## 4. Backend architecture

| Concern | Contract |
|---|---|
| Entry point | `code/backend/cmd/api/main.go` |
| Module | `github.com/ThanhNV121097/project-af1f65b9/backend` |
| HTTP mux | Standard library `net/http` |
| DB driver | `github.com/jackc/pgx/v5/stdlib` |
| Port | Read `PORT`, fallback `APP_PORT`, fallback `8080` |
| Database | Read `DATABASE_URL`; no separate DB host/user/password vars |
| Migrations | Embedded SQL files under `internal/migrations/sql`, applied on boot |
| Health | `/healthz` returns 200 only after migrations and `SELECT 1` succeed |
| Errors | Public API uses documented JSON envelope; logs keep internal detail |

Backend starts only if `DATABASE_URL` exists, migrations apply, and database ping succeeds. Migration tracking uses `schema_migrations(version text primary key, applied_at timestamptz not null default now())`.

## 5. Frontend architecture

| Concern | Contract |
|---|---|
| App router root | `code/frontend/app/page.tsx` |
| Shared CSS | `code/frontend/app/globals.css` |
| Story components | `code/frontend/components/{Component}.tsx` |
| Story CSS modules | `code/frontend/components/{Component}.module.css` |
| Story mocks | `code/frontend/lib/mock/{story-slug}.ts`, deleted by API story |
| Config | One file: `next.config.js` |
| Styling | Tailwind v3 plus CSS custom properties in `globals.css` |
| Client components | First line must be literal `"use client"` when browser APIs or hooks are used |

`app/page.tsx` is composition root only. Later UI story adds one import and one element. It must not become feature implementation.

## 6. Design tokens

`globals.css` owns shared tokens for all six required categories:

| Category | Tokens |
|---|---|
| Color | `--color-bg`, `--color-surface`, `--color-border`, `--color-text`, `--color-text-muted`, `--color-focus` |
| Spacing | `--space-3`, `--space-4`, `--space-6`, `--space-8` |
| Typography | `--font-body`, `--font-heading`, `--text-base`, `--text-sm`, `--text-hero` |
| Radius | `--radius-none` |
| Shadow | `--shadow-none` |
| Motion | `--duration-none` |

CSS modules must use tokens. No hardcoded color, `rgb()`, or spacing values. No `var(--token, fallback)`.

## 7. Environment variables

### Root compose

| Key | Used by | Purpose |
|---|---|---|
| `POSTGRES_USER` | compose/db/backend | Local database role |
| `POSTGRES_PASSWORD` | compose/db/backend | Local database password |
| `POSTGRES_DB` | compose/db/backend | Local database name |
| `BACKEND_PORT` | compose | Host port for backend |
| `FRONTEND_PORT` | compose | Host port for frontend |
| `NEXT_PUBLIC_API_URL` | frontend build | Browser-visible backend URL |

### Backend

| Key | Required | Purpose |
|---|---|---|
| `DATABASE_URL` | yes | PostgreSQL connection URL |
| `PORT` | no | HTTP listen port |
| `APP_PORT` | no | Secondary port fallback |

### Frontend

| Key | Required | Purpose |
|---|---|---|
| `NEXT_PUBLIC_API_URL` | yes | Public backend API base URL |

`.env.example` files list keys and comments only. No secrets committed.

## 8. Naming conventions

| Thing | Convention |
|---|---|
| Go packages | lowercase, short, no underscores |
| Go exported names | only when used across packages |
| HTTP paths | `/api/{resource}` for API, `/healthz` for health |
| JSON fields | lower camel case |
| SQL tables | snake_case plural nouns |
| SQL columns | snake_case |
| React components | PascalCase default function export |
| CSS tokens | semantic custom properties, `--category-purpose` |

## 9. Run and verify

Local development:

```bash
cp .env.example .env
docker compose --profile local up --build
```

Backend checks:

```bash
cd code/backend
go build ./...
go vet ./...
go test ./...
```

Frontend checks:

```bash
cd code/frontend
npm ci
npm run lint
npm run build
npm test --if-present
```

CI gate: `.github/workflows/ci.yml` runs same checks on pull requests. Container workflow also builds and boots stack.

## 10. Decisions and tradeoffs

| Decision | Chosen | Rejected | Tradeoff |
|---|---|---|---|
| Shape | fullstack | static page | Requirement says text stored in PostgreSQL and served by API. More moving parts accepted to prove pipeline. |
| Backend router | Go stdlib `net/http` | Gin/Echo | Stdlib enough for one endpoint; less dependency surface. |
| DB access | `database/sql` with pgx stdlib | ORM | One table and one query do not need ORM complexity. |
| Migrations | Self-applied SQL on boot | External migration job | Runtime database starts empty; boot-time migration avoids missing setup step. |
| Frontend | Next.js App Router | Static HTML | Default stack and future story mounting support. |
| Styling | Tailwind plus CSS tokens | Component-local hardcoded CSS | CI can enforce tokens and prevent visual drift. |
| Health | DB-backed `/healthz` | Process-only ping | Detects broken DB/migrations before service marked healthy. |
| Error envelope | `{ error: { code, message } }` | Plain text errors | One contract for all API failures; slightly more JSON boilerplate. |

## 11. Risks and constraints

| Risk | Mitigation |
|---|---|
| Migration SQL embedded from wrong path | Embed lives in `internal/migrations`, beside `sql/`. |
| Frontend hardcodes greeting | Story review checks API usage and no `Hello Word` in frontend source. |
| Health check false positive | `/healthz` runs `SELECT 1` after migrations. |
| Local compose profile missed | Use `docker compose --profile local up --build` for local DB. Deployment injects `DATABASE_URL`. |
| Token CI rejects CSS | Keep shared values in `globals.css`; CSS modules use `var(--token)`. |

## 12. Not in scaffold

- No product endpoint beyond health in scaffold; `GET /api/greeting` contract is in `services.md` for story implementation.
- No frontend greeting component yet; story owns component, API fetch, states, and mock removal.
- No seed content route; greeting row is inserted by migration because project needs one stored row.
