---
version: 1
project: Grafana Test
sourceCommit: e68275f3283a48c57686ddcd40674d3a048d2293
generatedAt: 2026-11-12T00:47:48Z
---

# Design rules

Forged's implementation agents follow this file on every frontend change. Edit it here, in the repository. Forged reads the new version once it is on the default branch.

## Precedence
<!-- forged:section id=precedence status=recommendation -->

- The code is the source of truth for what the product currently is. This file is authoritative for what you write next. <!-- forged:rule id=precedence.current-vs-next origin=inferred confidence=assumed -->
- When this file and the code disagree, follow this file for new code and name the conflict in your change description. Do not rewrite existing code to match unless the task asks for it. <!-- forged:rule id=precedence.conflicts origin=inferred confidence=assumed -->
- The copy of this file on the default branch is the only copy that counts. <!-- forged:rule id=precedence.canonical origin=inferred confidence=assumed -->
- The repository's design documents and instruction files (listed under Where this came from) come before this file. Where one says something more specific, such as a size, a breakpoint or a state for one kind of element, follow the document. <!-- forged:rule id=precedence.design-docs origin=inferred confidence=assumed -->

## Where tokens live
<!-- forged:section id=tokens status=observed -->

- Tokens live in `packages/grafana-data/src/themes/palette_new.ts`. Never hardcode a value that already exists as a token. <!-- forged:rule id=tokens.file-packages-grafana-data-src-themes-palette-new-ts origin=extracted confidence=code -->
  from: packages/grafana-data/src/themes/palette_new.ts:3-197
- Tokens live in `packages/grafana-data/src/themes/createVisualizationColors.ts`. Never hardcode a value that already exists as a token. <!-- forged:rule id=tokens.file-packages-grafana-data-src-themes-createvisualizationcolors-ts origin=extracted confidence=code -->
  from: packages/grafana-data/src/themes/createVisualizationColors.ts:434-574

## Color
<!-- forged:section id=color status=observed -->

- Color tokens are declared in this repository. It covers about 60% of the colors written in the code. The rest are colors outside the declared tokens. <!-- forged:rule id=color.tokens-only origin=extracted confidence=code -->
  from: emails/templates/partials/layout/theme.css:2, public/sass/grafana.dark.css:124
- Text colors: use `palette-blue-light-text` (#1F62E0) by default. Where a different text color is needed, use `palette-blue-dark-text` (#6E9FFF), `palette-ink-100` (#DADEE3), `palette-ink-150` (#C9CED4), `palette-ink-200` (#B9BEC6) or `palette-ink-250` (#A5ABB4). <!-- forged:rule id=color.role-text origin=recommendation confidence=code -->
  from: packages/grafana-data/src/themes/palette_new.ts:3-197
- Status colors, for success, warnings and errors: `palette-red-dark-main` (#D10E5C), `palette-red-light-main` (#E0226E), `palette-orange-dark-main` (#FF9900), `basic-colors-yellow` (#ECBB13), `palette-green-dark-main` (#1A7F4B), `palette-green-light-main` (#1B855E). <!-- forged:rule id=color.role-status origin=extracted confidence=code -->
  from: packages/grafana-data/src/themes/palette_new.ts:3-197
- Base palette colors behind the roles: `basic-colors-white` (#FFFFFF), `palette-gray-15` (#22252B), `basic-colors-black` (#000000), `basic-colors-gray-4` (#D8D9DA), `basic-colors-gray-7` (#FBFBFB), `palette-gray-95` (#F4F5F5), `basic-colors-gray-6` (#F4F5F8), `palette-gray-05` (#111217). Prefer a role color above for new work. <!-- forged:rule id=color.base-palette origin=recommendation confidence=code -->
  from: packages/grafana-data/src/themes/palette_new.ts:3-197

## Type
<!-- forged:section id=type status=not-established -->

Left empty. Eighteen text styles are in use and they do not resolve to a declared type scale. Not enough evidence to state a rule.

## Spacing
<!-- forged:section id=spacing status=not-established -->

Left empty. Thirty-two spacing values are in use and they do not resolve to a declared scale. Not enough evidence to state a rule.

## Components
<!-- forged:section id=components status=observed -->

- Reuse the existing components before writing new markup: `Page`, `Icon`, `Stack`, `Button`, `Field`, `Input`, `OptionsPaneItemDescriptor`, `OptionsPaneCategoryDescriptor`, `IconButton`, `LogListModel`. They are defined in `public/app/core/components/Page`, `packages/grafana-ui/src/components/Icon`, `packages/grafana-ui/src/components/Layout/Stack`, `packages/grafana-ui/src/components/Button`, `packages/grafana-ui/src/components/Forms`, `packages/grafana-ui/src/components/Input`, `public/app/features/dashboard/components/PanelEditor`, `packages/grafana-ui/src/components/IconButton`, `public/app/features/logs/components/panel`. <!-- forged:rule id=components.reuse origin=extracted confidence=code -->
  from: public/app/core/components/Page/Page.tsx:1
- Buttons use `Button` from `packages/grafana-ui/src/components/Button/Button.tsx`. Do not build a button out of a `div` or a plain `button` element. A control that only goes to another page stays a link, styled as a button where it needs to look like one. <!-- forged:rule id=components.button origin=extracted confidence=code -->
  from: packages/grafana-ui/src/components/Button/Button.tsx:1
- Text fields use `Input` from `packages/grafana-ui/src/components/Input/Input.tsx`. Do not use a plain `input` element. <!-- forged:rule id=components.input origin=extracted confidence=code -->
  from: packages/grafana-ui/src/components/Input/Input.tsx:1
- Selects use `Select` from `packages/grafana-ui/src/components/Select/Select.tsx`. Do not use a plain `select` element. <!-- forged:rule id=components.select origin=extracted confidence=code -->
  from: packages/grafana-ui/src/components/Select/Select.tsx:1
- Tables use `Table` from `packages/grafana-ui/src/components/Table/Table.tsx`. Do not use a plain `table` element. <!-- forged:rule id=components.table origin=extracted confidence=code -->
  from: packages/grafana-ui/src/components/Table/Table.tsx:1
- `Button` has 5 variants. Use one of them rather than restyling it. <!-- forged:rule id=components.button-variants origin=extracted confidence=code -->
  from: packages/grafana-ui/src/components/Button/Button.tsx:1
- `IconButton` has 3 variants. Use one of them rather than restyling it. <!-- forged:rule id=components.iconbutton-variants origin=extracted confidence=code -->
  from: packages/grafana-ui/src/components/IconButton/IconButton.tsx:1
- Do not rebuild a component that already exists, even as a close copy. <!-- forged:rule id=components.no-rebuild origin=extracted confidence=code -->

## Page patterns
<!-- forged:section id=patterns status=observed -->

- **`frontend-testing-strategy`** — general frontend test conventions: assert real behavior (not existence), avoiding AI-slop tests, and the anti-flake rules. (From packages/grafana-ui/AGENTS.md, @grafana/ui — agent guide; applies only in packages/grafana-ui/) <!-- forged:rule id=patterns.doc-agents-frontend-testing-strategy-general-fronte origin=imported confidence=docs -->
  from: packages/grafana-ui/AGENTS.md:8
- **`panel-testing-strategy`** — builds on the above with what's specific to visualization code: the canvas draw-call snapshot harness and panel a11y/interaction-snapshot E2E. (From packages/grafana-ui/AGENTS.md, @grafana/ui — agent guide; applies only in packages/grafana-ui/) <!-- forged:rule id=patterns.doc-agents-panel-testing-strategy-builds-on-the origin=imported confidence=docs -->
  from: packages/grafana-ui/AGENTS.md:10
- **`add-e2e-selectors`** — add versioned `@grafana/e2e-selectors` and wire `data-testid` into JSX. (From packages/grafana-ui/AGENTS.md, @grafana/ui — agent guide; applies only in packages/grafana-ui/) <!-- forged:rule id=patterns.doc-agents-add-e2e-selectors-add-versioned-grafana origin=imported confidence=docs -->
  from: packages/grafana-ui/AGENTS.md:12
- Hover: change the color by one step with a token, and keep the element's size and position the same. <!-- forged:rule id=patterns.state-hover origin=inferred confidence=assumed -->
- Focus: follow the focus style under Interaction and motion. <!-- forged:rule id=patterns.state-focus origin=extracted confidence=code -->
- Pressed: keep the element's size and layout while it is pressed, and follow the repository's own pressed style where it has one, with no new pressed color. <!-- forged:rule id=patterns.state-pressed origin=inferred confidence=assumed -->
- Disabled: lower the contrast, keep the element's size and label, and prefer `aria-disabled` over `disabled`, and block the action in its handler: `aria-disabled` alone does not stop a click or typing. `aria-disabled` keeps the control focusable, so the reason it is unavailable can be reached. <!-- forged:rule id=patterns.state-disabled origin=inferred confidence=assumed -->
- Loading: show `PageLoader` in place of content that is still loading. <!-- forged:rule id=patterns.state-loading origin=extracted confidence=code -->
- Empty: when a list or page has nothing to show, say so and give the next step. <!-- forged:rule id=patterns.state-empty origin=inferred confidence=assumed -->
- Errors: say what went wrong and how to recover, next to where it happened, and keep what the person entered. <!-- forged:rule id=patterns.state-error origin=inferred confidence=assumed -->

## Interaction and motion
<!-- forged:section id=interaction status=observed -->

- Use 200ms for transitions. 100ms, 150ms and 400ms appear in 14 places; use 200ms for new work. <!-- forged:rule id=interaction.durations origin=recommendation confidence=code -->
- Interactive elements show keyboard focus with `:focus` styles drawn as a box-shadow. <!-- forged:rule id=interaction.focus-visible origin=extracted confidence=code -->

## Responsive behavior
<!-- forged:section id=responsive status=observed -->

- Breakpoints in use: min-width 0px, min-width 544px, min-width 769px, min-width 992px, min-width 1200px. <!-- forged:rule id=responsive.breakpoints origin=extracted confidence=code -->

## Devices and browsers
<!-- forged:section id=devices status=recommendation -->

- Check every page at 320, 375, 390, 768 and 1280px wide. <!-- forged:rule id=devices.widths origin=inferred confidence=assumed -->
- Never let a page scroll sideways at those widths. Wrap, stack or truncate content instead. <!-- forged:rule id=devices.no-horizontal-scroll origin=inferred confidence=assumed -->
- Keep page headers usable at every width: the title, back link and main action stay visible and tappable. <!-- forged:rule id=devices.headers origin=inferred confidence=assumed -->
- Make every tap target at least 44 by 44px on touch screens. <!-- forged:rule id=devices.touch-targets origin=inferred confidence=assumed -->
- Do not size full-height layouts with `100vh`, which iOS Safari measures to include the space behind its toolbars. Use `100dvh`. <!-- forged:rule id=devices.viewport-height origin=inferred confidence=assumed -->
- Keep the focused field and its submit button in view when the on-screen keyboard opens. <!-- forged:rule id=devices.keyboard origin=inferred confidence=assumed -->
- Give text fields a font size of at least 16px, or iOS Safari zooms the page when they get focus. <!-- forged:rule id=devices.input-zoom origin=inferred confidence=assumed -->
- Test date and number fields on iOS Safari and Android Chrome, whose pickers and keyboards differ. Use `inputmode` to choose the keyboard, and do not use `type="number"` for codes or IDs. <!-- forged:rule id=devices.date-number-inputs origin=inferred confidence=assumed -->
- Never hide an action behind hover alone. Touch screens have no hover, so every control must work by tap and by keyboard. Put hover styles inside `@media (hover: hover)`. <!-- forged:rule id=devices.hover origin=inferred confidence=assumed -->
- Keep keyboard focus visible in every browser. Never remove the outline without drawing a replacement. <!-- forged:rule id=devices.focus origin=inferred confidence=assumed -->

## Accessibility floor
<!-- forged:section id=accessibility status=recommendation -->

- Text has a contrast ratio of at least 4.5:1 against its background. Text 24px and larger, or 18.66px and bold, needs at least 3:1. <!-- forged:rule id=accessibility.contrast origin=inferred confidence=assumed -->
- Every interactive element has a visible focus style. <!-- forged:rule id=accessibility.focus origin=inferred confidence=assumed -->
- Interactive targets are at least 24 by 24 CSS pixels. <!-- forged:rule id=accessibility.targets origin=inferred confidence=assumed -->
- Every form field has a visible label or an accessible name. <!-- forged:rule id=accessibility.labels origin=inferred confidence=assumed -->
- Images that carry meaning have alt text. Decorative images have an empty alt. <!-- forged:rule id=accessibility.alt origin=inferred confidence=assumed -->

## Writing style
<!-- forged:section id=writing status=not-established -->

Left empty. Voice and terminology cannot be read from code, and no design document in this repository covers them.

## Things not to do
<!-- forged:section id=prohibited status=not-established -->

Left empty. Nothing beyond the rules above, and no design document in this repository lists things not to do.

## Assumptions
<!-- forged:section id=assumptions status=recommendation -->

- Where the code uses several values for the same job, Forged picked one as the default for new work and named the values it replaces. If a default is wrong, change it in this file. <!-- forged:rule id=assumptions.recommendations origin=recommendation confidence=code -->

## Where this came from
<!-- forged:section id=provenance status=observed -->

- Generated by Forged from commit `e68275f` on November 12, 2026. <!-- forged:rule id=provenance.commit origin=extracted confidence=code -->
- Read from the code: color, components. Not established: type, spacing, corners, shadows. <!-- forged:rule id=provenance.categories origin=extracted confidence=code -->
- Design documents read: `contribute/backend/style-guide.md`, `packages/grafana-ui/src/DesignPrinciples.mdx`, `packages/grafana-ui/src/VoiceAndTone.mdx`, `pkg/router/specs/2026-08-17-router-discovery-openapi-design.md`, `pkg/router/specs/2026-08-19-router-circuit-breaker-design.md`, `pkg/router/specs/2026-09-11-router-aggregate-discovery-design.md`, `pkg/router/specs/2026-09-25-router-design-notes.md`, `pkg/router/specs/2026-10-28-router-group-visibility-design.md`. Rules that end with "(From ...)" are quoted from them. <!-- forged:rule id=provenance.design-docs origin=extracted confidence=code -->
- Instruction files with design rules: `packages/grafana-ui/AGENTS.md`. Rules that end with "(From ...)" are quoted from them. <!-- forged:rule id=provenance.instruction-docs origin=extracted confidence=code -->
- Detected and not read: Storybook. <!-- forged:rule id=provenance.not-read origin=extracted confidence=code -->
