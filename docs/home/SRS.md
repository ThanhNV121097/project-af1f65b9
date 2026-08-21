# SRS — home

Module: `home`
Last updated: 2025-08-13
Design: [View the approved design](http://localhost:8080/design/af1f65b9-87cd-4577-9cb9-09b5865dd8e7)
Design system: `design/design-system.md`

## 1. Purpose

`home` delivers the only user-facing screen in `hello-word-2`. It shows one stored greeting, centered on a plain white page, so the project can prove data flows from PostgreSQL through backend API to frontend render. If this module does not work, the whole end-to-end proof fails.

## 2. Actors

| Actor | Who they are | What they may do in this module |
|---|---|---|
| Guest | Any visitor opening project page | View greeting screen and read stored text |
| API | Backend service exposing greeting value | Return stored greeting to frontend |
| Database | PostgreSQL row store | Persist greeting text used by page |

## 3. Scope

**In scope** — the functions specified below, by their plan titles:

- Show stored greeting

**Out of scope**

- Any second screen, settings page, or navigation — not part of this proof.
- Editing greeting text — belongs to no planned module and is deliberately not built.
- Styling beyond plain white background, black text, and centered layout — excluded by project brief.

## 4. Functional requirements

### 4.1 Show stored greeting

**Requirement HOME-001 — Load greeting from stored data**

*As a* Guest, *I want to* see greeting text from stored data, *so that* page proves frontend reads backend data instead of hardcoding it.

Behaviour:

1. Guest opens home page.
2. Frontend requests current greeting from backend API.
3. Backend returns greeting value read from PostgreSQL.
4. Frontend renders returned value as visible page text.

**Acceptance criteria** — each maps one-to-one onto a test case in `docs/home/test-cases/show-stored-greeting.md`. Given/When/Then, no compound conditions: one behaviour per criterion.

| # | Given | When | Then |
|---|---|---|---|
| AC-1 | stored greeting row exists with value `Hello Word` | Guest opens home page | page shows `Hello Word` |
| AC-2 | stored greeting row exists with value `Hello Word` | Guest inspects page source or network response | frontend does not contain hardcoded greeting text and displayed text comes from API response |
| AC-3 | stored greeting row exists with value `Hello Word` | Guest requests greeting through API | API returns stored value exactly |
| AC-4 | stored greeting text is unavailable from backend or database | Guest opens home page | page shows error state, not stale or empty greeting |

**Failure, boundary and permission behaviour** — the part most often skipped and most often the source of bugs. Every row needs a defined outcome; "should not happen" is not an outcome.

| Case | Condition | Expected behaviour |
|---|---|---|
| Invalid input | Greeting row contains empty text | Page shows empty-state error rather than silent blank content |
| Boundary | Greeting text is 1 character or a long single line within storage limit | Page renders full text without truncation or layout break |
| Not found | Greeting row missing | Page shows error state with no greeting text |
| Not permitted | N/A for public home page | No permission gate; any guest may view page |
| Conflict | Two updates to greeting exist in storage | Latest committed greeting is the value API returns |
| Upstream failure | Backend or PostgreSQL unavailable | Page shows error state and does not invent fallback text |

**Data touched** — the fields this function reads and writes, in product terms.

| Field | Type | Required | Rule |
|---|---|---|---|
| greeting text | text | yes | Exactly one stored row supplies visible home page text; no frontend copy source may override it |

### 4.2 Show stored greeting screen state

**Requirement HOME-002 — Center greeting on plain screen**

*As a* Guest, *I want to* see greeting centered on plain screen, *so that* layout matches minimal project brief.

Behaviour:

1. Guest opens home page.
2. Page background stays white.
3. Greeting text stays black and centered horizontally and vertically.
4. No extra decoration, motion, palette, or navigation appears.

**Acceptance criteria** — each maps one-to-one onto a test case in `docs/home/test-cases/show-stored-greeting.md`. Given/When/Then, no compound conditions: one behaviour per criterion.

| # | Given | When | Then |
|---|---|---|---|
| AC-1 | page is loaded | Guest views page | background is white |
| AC-2 | page is loaded with greeting text | Guest views page | greeting text is black |
| AC-3 | page is loaded with greeting text | Guest views page | greeting is centered horizontally and vertically |
| AC-4 | page is loaded | Guest views page | no animation, palette, or extra screen content appears |

**Failure, boundary and permission behaviour**

| Case | Condition | Expected behaviour |
|---|---|---|
| Invalid input | Browser viewport is very small | Content still stays centered without horizontal scroll |
| Boundary | Browser viewport is large | Content remains centered and does not drift to a corner |
| Not found | N/A | No alternate screen; same single-page layout remains the only screen |
| Not permitted | N/A | Any guest may view page; no auth requirement |
| Conflict | N/A | Layout rules do not change based on data race |
| Upstream failure | Greeting fetch fails | Error state uses same minimal layout and does not add fallback styling |

**Data touched** — the fields this function reads and writes, in product terms.

| Field | Type | Required | Rule |
|---|---|---|---|
| layout state | presentation | yes | One screen only; greeting stays centered on white background with black text |

## 5. Screens

| Screen | Section in the design | Functions it serves | States that must exist |
|---|---|---|---|
| Single greeting screen | Greeting screen | HOME-001, HOME-002 | loading, default, error |

## 6. Non-functional requirements

| Area | Requirement |
|---|---|
| Performance | Home page initial render completes within 2 seconds on a typical connection after backend response is available |
| Accessibility | Greeting text remains readable with keyboard focus order preserved for any future interactive element; contrast stays at least 4.5:1 between black text and white background |
| Responsive | Page works from 320px width upward with no horizontal scroll |
| Localisation | Copy remains in English exactly as supplied: `Hello Word` |
| Privacy | No personal data is stored or displayed; only shared greeting text is read from database |

## 7. Dependencies and assumptions

- **Depends on:** PostgreSQL, for persistent greeting row.
- **Depends on:** Backend API, for reading stored greeting and serving it to frontend.
- **Assumption:** Exactly one greeting row exists for this page; if it is missing, page shows error state until data is restored.

| Open question | Proposed default | Who decides |
|---|---|---|
| Should error state show copy or just blank failure screen? | Show minimal error state with no greeting text | Stakeholder / TL |

## 8. Traceability

| Plan item | Requirement ids | Test cases |
|---|---|---|
| Show stored greeting | HOME-001, HOME-002 | `test-cases/show-stored-greeting.md` |
