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

## Endpoints

### `GET /healthz`

Purpose: service readiness for compose and deployer.

Request: no body.

Response `200 text/plain`:

```text
ok
```

Failure:

| Status | Code | When |
|---|---|---|
| `503` | `service_unavailable` | Migrations not applied or database `SELECT 1` fails |

### `GET /api/greeting`

Purpose: return stored greeting for home page.

Request: no body, no query parameters.

Response `200 application/json`:

```json
{
  "text": "Hello Word"
}
```

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
- Frontend must not hardcode `Hello Word` outside tests or mocked UI branch.
