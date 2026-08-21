# Test Cases — Show stored greeting

Risk level: low. Single public read-only screen, but case must cover data flow, failure state, and layout because whole proof depends on them.

## Automated cases

### Scenario: Home page shows stored greeting

**Given** stored greeting row exists with value `Hello Word`

**When** guest opens home page

**Then** page shows visible text `Hello Word` centered on screen

Trace to: HOME-001 AC-1, HOME-002 AC-3

### Scenario: Home page text comes from API, not hardcoded frontend copy

**Given** stored greeting row exists with value `Hello Word`

**When** guest inspects page source and network response after page load

**Then** frontend source contains no hardcoded greeting text and visible greeting equals API response body text

Trace to: HOME-001 AC-2

### Scenario: API returns stored greeting value exactly

**Given** stored greeting row exists with value `Hello Word`

**When** guest requests `GET /api/greeting`

**Then** response status is `200` and JSON body is exactly `{"text":"Hello Word"}`

Trace to: HOME-001 AC-3

### Scenario: Missing stored greeting shows error state

**Given** greeting row `id = 1` is missing

**When** guest opens home page

**Then** page shows error state and no greeting text

Trace to: HOME-001 AC-4, HOME-001 not found failure, HOME-002 upstream failure

### Scenario: Empty stored greeting shows empty-state error

**Given** greeting row exists but stored text is empty after validation

**When** guest opens home page

**Then** page shows empty-state error and no blank greeting content

Trace to: HOME-001 invalid input

### Scenario: Backend or database unavailable shows error state

**Given** backend cannot load greeting because database read fails

**When** guest opens home page

**Then** page shows error state and does not invent fallback greeting text

Trace to: HOME-001 upstream failure, HOME-001 AC-4

### Scenario: White background on loaded page

**Given** page is loaded

**When** guest views page

**Then** background color is white (`#FFFFFF`)

Trace to: HOME-002 AC-1

### Scenario: Black greeting text on loaded page

**Given** page is loaded with greeting text

**When** guest views page

**Then** greeting text color is black (`#000000`)

Trace to: HOME-002 AC-2

### Scenario: Greeting stays centered on loaded page

**Given** page is loaded with greeting text

**When** guest views page

**Then** greeting is centered horizontally and vertically

Trace to: HOME-002 AC-3

### Scenario: No extra UI on loaded page

**Given** page is loaded

**When** guest views page

**Then** no animation, palette, navigation, or extra screen content appears

Trace to: HOME-002 AC-4

## Manual checks

### Scenario: Page remains centered in very small viewport

**Given** browser viewport is 320px wide

**When** guest opens home page

**Then** content remains centered with no horizontal scroll

Reason manual: viewport rendering and scroll behavior are browser-observable, not asserted by API contract alone.

Trace to: HOME-002 invalid input

### Scenario: Page remains centered in large viewport

**Given** browser viewport is large

**When** guest opens home page

**Then** content remains centered and does not drift to a corner

Reason manual: visual alignment across viewport sizes needs browser inspection.

Trace to: HOME-002 boundary

### Scenario: Home page loads within performance target

**Given** backend response is available

**When** guest opens home page

**Then** initial render completes within 2 seconds on typical connection

Reason manual: timing depends on environment and network.

Trace to: HOME non-functional performance

### Scenario: Page remains English exactly as supplied

**Given** page is loaded

**When** guest views greeting text

**Then** copy matches `Hello Word` exactly

Reason manual: localization copy is product-level visual check for this proof.

Trace to: HOME non-functional localisation
