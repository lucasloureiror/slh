# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project

`slh` (Service Level Helper) is a static website, served at slh.loureiro.tech. Given an SLO/SLA percentage it shows the maximum allowable downtime per period. It can also show the minimum probing frequency (from MTTR, incidents and probes), and turn a downtime duration into an availability percentage. It used to be a Go CLI. The CLI is gone and the project is web only.

## Commands

```bash
npm test      # node --test test/ (what CI runs)
npm start     # python3 -m http.server -d docs 8000, then open http://localhost:8000
```

There is no build step and no dependencies. Don't add a JS framework or bundler unless the site outgrows a single page.

## Architecture

- `docs/calc.js`: all the logic, as pure functions. It exports through `module.exports` for the tests and defines the global `SLH` in the browser.
  - Periods are Daily, Weekly, Monthly (30.44 days), Quarterly (90 days) and Yearly (365 days). Period lengths scale with hours per day (1–24).
  - `parseServiceLevel` accepts `,` or `.` as the decimal separator. A value `1 ≤ x < 16` is read as a **count of nines** (`1` → 90, `3` → 99.9), so `5` means 99.999%, not 5%. A trailing `%` always means a plain percentage (`5%` → 5). Values ≤ 0 or > 100 are rejected.
  - Probe formula: `(downtime − MTTR×incidents) / (incidents × probes)`. Periods where the budget doesn't cover the repair time get `frequency: null` and the page shows "Not achievable".
- `docs/index.html`: the calculator page only. Inputs sync to query params (`?sl=99&mttr=0h20m&incidents=3&probes=2`, `?mode=reverse&outage=1d2h3m4s`, `&hours=8`). Reverse takes four number fields (days, hours, minutes, seconds) but keeps the `outage` duration string in the URL, so old links work.
- `docs/app.js`: the calculator UI (rendering, URL sync). It uses the global `SLH` from `calc.js`.
- `docs/style.css`: styles shared by every page.
- `docs/usage.html`, `docs/probe.html`, `docs/faq.html`: the documentation, one static page each. They share the header, nav and footer markup with `index.html`, so a nav change has to be made in all four pages. Pages link to each other as `foo.html`, which works locally and on Cloudflare Pages (it serves the clean URL).
- `test/calc.test.js`: tests for `calc.js` using `node:test`.

## CI and hosting

`.github/workflows/ci.yml` runs `npm test` on pushes and pull requests to main. There is no deploy workflow in the repo. Check how `docs/` is published before changing the hosting setup.
