# Design System — hello-word-2

> Source of truth: the approved `index.html` (preview: unavailable in tool output).
> Every value below is extracted from it. Changing a value here without changing the approved design is a defect.

Last updated: 2025-08-14

## 1. Foundations

### 1.1 Color

Semantic tokens. Name by job, never by hue.

| Token | Value | Used for |
|---|---|---|
| `--color-bg` | `#ffffff` | Page background |
| `--color-surface` | `#ffffff` | Surface / panel background |
| `--color-border` | `#000000` | Default border and divider |
| `--color-text` | `#000000` | Body text and primary text |
| `--color-text-muted` | `#666666` | Secondary text, captions |
| `--color-focus` | `#000000` | Focus ring |

#### Contrast audit

| Foreground | Background | Ratio | Passes |
|---|---|---|---|
| `--color-text` | `--color-bg` | `21:1` | AA / AA Large |
| `--color-text-muted` | `--color-bg` | `5.74:1` | AA / AA Large |
| `--color-border` | `--color-surface` | `21:1` | UI border |
| `--color-focus` | `--color-bg` | `21:1` | AA / AA Large |

### 1.2 Spacing

Base unit: `4px`. Every margin, padding, and gap in product uses one of these.

| Token | Value |
|---|---|
| `--space-3` | `12px` |
| `--space-4` | `16px` |
| `--space-6` | `24px` |
| `--space-8` | `32px` |

### 1.3 Typography

Font families:

- Body: `Arial, Helvetica, sans-serif`
- Headings: `Arial, Helvetica, sans-serif`
- Mono: none used

| Token | Size | Line height | Weight | Used for |
|---|---|---|---|---|
| `--text-base` | `16px` | `1.4` | `400` | Body |
| `--text-sm` | `14px` | `1.4` | `400` | Secondary text |
| `--text-hero` | `clamp(2.75rem, 8vw, 5rem)` | `1` | `400` | Main greeting |

Heading levels are not used as a hierarchy; only `h1` appears.

### 1.4 Radius, border, shadow, motion

| Token | Value | Used for |
|---|---|---|
| `--border-width` | `1px` | Borders and dividers |
| `--radius-none` | `0` | No rounded corners in approved UI |
| `--shadow-none` | `none` | All surfaces |
| `--duration-none` | `0ms` | No motion in approved UI |

Motion respects `prefers-reduced-motion: reduce` by default because the approved UI has no animated transitions.

### 1.5 Layout and breakpoints

| Name | Min width | Container | Columns | Gutter |
|---|---|---|---|---|
| `sm` | `0px` | full viewport | 1 | `24px` |
| `md` | `768px` | full viewport | 1 | `24px` |
| `lg` | `1024px` | full viewport | 1 | `24px` |
| `xl` | `1280px` | full viewport | 1 | `24px` |

Z-index scale (only these values are allowed):

| Layer | Value |
|---|---|
| Base | `0` |
| Sticky header | not used |
| Dropdown | not used |
| Modal backdrop | not used |
| Modal | not used |
| Toast | not used |

## 2. Components

### 2.1 Greeting screen

**Purpose** — Centered one-line greeting on plain white page. Use for this single end-to-end proof only. Not for multi-section layouts.

**Anatomy** — `[main.frame] [section.state-*] [div.hero] [h1#greeting] [p] [div.controls] [div.status]`

**Variants**

| Variant | Tokens | When to use |
|---|---|---|
| Ready | `--color-bg`, `--color-text`, `--color-text-muted`, `--space-4`, `--space-6`, `--space-8` | Stored greeting loaded successfully |
| Loading | same tokens | Backend fetch in progress |
| Error | same tokens | Backend fetch failed |

**States**

| State | Visual change | Tokens |
|---|---|---|
| Default | Centered greeting, muted helper line, state buttons, status line | `--color-bg`, `--color-text`, `--color-text-muted`, `--space-3`, `--space-4`, `--space-6`, `--space-8`, `--text-base`, `--text-sm`, `--text-hero` |
| Hover | No hover styling in approved UI | none |
| Focus (keyboard) | Button gets 2px outline in black with 2px offset | `--color-focus` |
| Active / pressed | Pressed button shows no separate visual state in approved UI | none |
| Disabled | Not used | none |
| Loading | Greeting screen still centered; helper line changes to loading copy | same as default |
| Error | Greeting screen still centered; helper line changes to error copy | same as default |
| Empty | Not used; data source always expected to return one row | none |

**Accessibility** — `main`, `section`, and buttons use native semantics. Buttons remain at least 44×44px by content plus padding. Focus indicator visible. `aria-live="polite"` on greeting wrapper. `aria-pressed` on state preview buttons.

### 2.2 Button

**Purpose** — Preview-state toggle only. Not for product actions; no other button patterns exist in approved UI.

**Anatomy** — `[button]`

**Variants**

| Variant | Tokens | When to use |
|---|---|---|
| Default | `--color-bg`, `--color-text`, `--color-border` | Inactive preview toggle |
| Pressed | `--color-text`, `#ffffff` text | Active preview toggle |

**Sizes**

| Size | Height | Padding | Text token |
|---|---|---|---|
| Default | content-driven | `10px 14px` | `--text-base` |

**States**

| State | Visual change | Tokens |
|---|---|---|
| Default | White background, black text, black border | `--color-bg`, `--color-text`, `--color-border` |
| Hover | No separate hover style in approved UI | none |
| Focus (keyboard) | 2px black outline with 2px offset | `--color-focus` |
| Active / pressed | Inverted colors when `aria-pressed="true"` | `--color-text`, `#ffffff` |
| Disabled | Not used | none |
| Loading | Not used | none |
| Error | Not used | none |
| Empty | Not used | none |

**Accessibility** — Native `button` element. `aria-pressed` indicates toggle state. Hit target may exceed text size because of padding.

## 3. Content and formatting

- Voice and tone: plain, factual, no marketing language.
- Date format: not shown.
- Number format: plain digits only.
- Capitalization: sentence case for status copy; title case not used.
- Empty-state and error-message wording pattern: short, direct, and action-oriented; approved copy uses one sentence.

## 4. Known deviations

| Where | Deviation | Why it stands | Follow-up |
|---|---|---|---|
| `hero p` / `status` / buttons | Preview controls exist in approved mockup even though product itself is only greeting screen | Stakeholder approved the design as shown | Keep them only if later implementation needs local state preview; otherwise omit from product UI |
| `body` | No explicit radius, shadow, or animation tokens beyond none / zero | Simple proof-of-pipeline screen | Keep flat styling unless approved design changes |
| `h1` | Heading hierarchy skips other heading levels | Only one heading needed | No action |

## 5. Change log

| Date | Change | Design PR |
|---|---|---|
| 2025-08-14 | Initial design system extracted from approved `index.html` | pending |
