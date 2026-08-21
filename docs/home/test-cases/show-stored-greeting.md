# Test Cases — Show stored greeting

Risk level: low. Single read-only home flow, but contract and layout both must be exact because this story proves end-to-end wiring.

## Automated coverage

### Scenario: Page shows stored greeting
**Given** stored greeting row exists with value `Hello Word`
**When** Guest opens home page
**Then** page shows `Hello Word`

Traceability: HOME-001 AC-1

### Scenario: Page text comes from API, not hardcoded frontend copy
**Given** stored greeting row exists with value `Hello Word`
**When** Guest inspects page source or network response
**Then** frontend does not contain hardcoded greeting text and displayed text comes from API response

Traceability: HOME-001 AC-2

### Scenario: API returns stored greeting exactly
**Given** stored greeting row exists with value `Hello Word`
**When** Guest requests greeting through API
**Then** API returns stored value exactly as JSON `{"text":"Hello Word"}` with status `200`

Traceability: HOME-001 AC-3

### Scenario: Home page shows error when greeting unavailable
**Given** stored greeting text is unavailable from backend or database
**When** Guest opens home page
**Then** page shows error state, not stale or empty greeting

Traceability: HOME-001 AC-4, failure path

### Scenario: Empty stored greeting becomes error state
**Given** greeting row exists with empty text
**When** Guest opens home page
**Then** page shows empty-state error rather than silent blank content

Traceability: HOME-001 invalid input

### Scenario: Missing greeting row shows error state
**Given** greeting row is missing
**When** Guest opens home page
**Then** page shows error state with no greeting text

Traceability: HOME-001 not found

### Scenario: Backend or database failure shows error state
**Given** backend or PostgreSQL is unavailable
**When** Guest opens home page
**Then** page shows error state and does not invent fallback text

Traceability: HOME-001 upstream failure

### Scenario: Page background is white
**Given** page is loaded
**When** Guest views page
**Then** background is white `#FFFFFF`

Traceability: HOME-002 AC-1

### Scenario: Greeting text is black
**Given** page is loaded with greeting text
**When** Guest views page
**Then** greeting text is black `#000000`

Traceability: HOME-002 AC-2

### Scenario: Greeting is centered both ways
**Given** page is loaded with greeting text
**When** Guest views page
**Then** greeting is centered horizontally and vertically

Traceability: HOME-002 AC-3

### Scenario: No extra content appears
**Given** page is loaded
**When** Guest views page
**Then** no animation, palette, or extra screen content appears

Traceability: HOME-002 AC-4

### Scenario: Small viewport keeps centered layout
**Given** browser viewport is very small
**When** Guest views page
**Then** content still stays centered without horizontal scroll

Traceability: HOME-002 small-viewport boundary

### Scenario: Large viewport keeps centered layout
**Given** browser viewport is large
**When** Guest views page
**Then** content remains centered and does not drift to a corner

Traceability: HOME-002 large-viewport boundary

### Scenario: Public guest may view page
**Given** any guest opens home page
**When** Guest views page
**Then** page is visible with no auth gate

Traceability: HOME-001 not permitted, HOME-002 not permitted

## Manual coverage

### Scenario: Centering and visual plainness on real browser
**Given** page is loaded in real browser
**When** Guest views page at normal size and scrolls if needed
**Then** greeting stays visually centered on plain white screen with black text and no extra UI

Reason: visual centering and plainness need human check in browser.

Traceability: HOME-002 AC-1 to AC-4
