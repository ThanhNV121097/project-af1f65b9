# Story — Show stored greeting

## User story

As a Guest, I want to see greeting text loaded from stored data, so that the page proves frontend reads backend data instead of hardcoding it.

## In scope

- Single home screen that requests current greeting from backend API and renders returned value.
- Centered plain page layout with white background and black text.
- Loading and error states for the same single screen.
- Stored greeting value is `Hello Word` in PostgreSQL and is the source of visible page text.

## Out of scope

- Any second screen, navigation, settings, or editing flow.
- Any copy, palette, motion, or decoration beyond plain white background, black text, and centered layout.
- Hardcoded greeting text in frontend source.

## UI scope

- One screen only: the single greeting screen in approved design.
- States: loading, default, error.
- Screen stays centered horizontally and vertically on white background with black text.
- No extra interactive UI beyond what the approved design already shows for this screen.

## Acceptance criteria

1. Given stored greeting row exists with value `Hello Word`, when Guest opens home page, then page shows `Hello Word`.
2. Given stored greeting row exists with value `Hello Word`, when Guest inspects page source or network response, then frontend does not contain hardcoded greeting text and displayed text comes from API response.
3. Given stored greeting row exists with value `Hello Word`, when Guest requests greeting through API, then API returns stored value exactly.
4. Given stored greeting text is unavailable from backend or database, when Guest opens home page, then page shows error state, not stale or empty greeting.
5. Given browser viewport is very small or large, when Guest views page, then greeting remains centered without horizontal scroll.

## Dependencies

- PostgreSQL with one stored greeting row.
- Backend API that serves current greeting value.
- Frontend page that reads greeting from backend instead of hardcoding it.
- Approved design and design system for the single greeting screen.
