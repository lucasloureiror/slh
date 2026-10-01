# Service Level Helper - slh

Service Level Helper (slh) is a web app for Site Reliability Engineers, DevOps professionals and similar roles. It turns a service level objective or agreement into the maximum allowable downtime for each period.

Use it at [slh.loureiro.tech](https://slh.loureiro.tech).

## Features

- Maximum allowable downtime for daily, weekly, monthly, quarterly and yearly periods, from a service level given as a percentage or a number of nines.
- Minimum probing frequency to keep your service level inside each period, from your mean time to repair (MTTR), expected incidents and the failed probes needed to alert.
- Reverse calculation: the availability left by a given amount of downtime.
- Shareable links: the page address holds your inputs.

## Development

The site is static, with no build step. `docs/calc.js` holds the calculations as pure functions, `docs/app.js` and `docs/index.html` are the calculator, and `usage.html`, `probe.html` and `faq.html` are the docs pages.

```bash
npm start   # serves docs/ at http://localhost:8000
npm test    # runs the calculation tests with node:test
```

## Probe frequency algorithm

```
probe frequency = buffer / (incidents × probes)
buffer          = downtime budget − (MTTR × incidents)
```

- **Incidents**: how many times you expect the service to become unavailable in the period.
- **Probes**: how many consecutive failed checks it takes to confirm an outage and fire an alert.
- **Buffer**: what is left of the downtime budget after repairing every expected incident.

If the repair time alone uses up the budget, no probing frequency can keep that service level for the period.

> This formula is experimental and there is no guarantee it fits your service.

## License

Service Level Helper is licensed under the [Apache License](LICENSE).
