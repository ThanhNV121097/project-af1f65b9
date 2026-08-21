# Services — hello-word-2

## API conventions

Base path: backend service root.

Success responses use JSON. Error responses use one envelope everywhere:

```json
{
  "error": {
    "code": "string",
    "message": "string"
  }
}
```

Rules:

- `code` is stable and machine-readable.
- `message` is short and safe for users.
- Internal database errors are logged by backend, not exposed.
- Public endpoints require no authentication.

## Story extension — Show stored greeting

Reviewed UI mock contract: `code/frontend/lib/mock/show-stored-greeting.ts` exports `GreetingResponse = { text: string }` and states `ready | loading | error | empty`. Backend supplies only ready success response and error envelope; frontend maps request pending to loading and `greeting_empty` to empty.

## Endpoints

### `GET /healthz`

Purpose: service readiness for compose and deployer.

Auth: none.

Request: no body.

Response `200 text/plain`:

```text
ok
```

Failure:

| Status | Code | Message | When |
|---|---|---|---|
| `503` | `service_unavailable` | `Service is unavailable.` | Migrations not applied or database `SELECT 1` fails |

### `GET /api/greeting`

Purpose: return stored greeting for home page.

Auth: none.

Request: no body, no query parameters.

Response `200 application/json`:

```json
{
  "text": "Hello Word"
}
```

Response shape matches reviewed mock `GreetingResponse` exactly. No envelope, list, ID, timestamp, or metadata is added because frontend only renders `text`.

Validation and failure:

| Status | Code | Message | When |
|---|---|---|---|
| `404` | `greeting_not_found` | `Greeting is not available.` | Row `id = 1` missing |
| `422` | `greeting_empty` | `Greeting is empty.` | Stored text is empty after validation |
| `500` | `internal_error` | `Greeting could not be loaded.` | Database read fails |

Frontend behavior tied to this contract:

- Loading state while request is pending.
- Default state renders `text` exactly from API.
- Error state shows minimal failure copy and no fallback greeting.
- Empty state is used for `greeting_empty` and renders no greeting text.
- Frontend must not hardcode `Hello Word` outside tests or mocked UI branch.

## Migration plan — Show stored greeting

Forward:

1. Add `greetings` table and seed row through SQL migration before enabling `GET /api/greeting`.
2. Backend reads `select text from greetings where id = 1` and returns `{ "text": value }`.

Backward:

1. Remove `GET /api/greeting` implementation with backend rollback.
2. Run down migration only in disposable/local environments; production rollback restores database backup.

Safety on populated tables:

- Forward route addition is safe; endpoint is new and public.
- Seed is idempotent with `on conflict (id) do nothing`.
- Backward route removal breaks frontend builds expecting API, so rollback frontend and backend together.
