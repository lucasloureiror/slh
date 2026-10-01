/*
Copyright 2024 github.com/lucasloureiror/slh maintainers

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Service level calculations as pure functions; the
// page in index.html handles rendering. Tests live in test/calc.test.js.
const SLH = (() => {
  const PERIODS = [
    { key: "daily", label: "Daily", days: 1 },
    { key: "weekly", label: "Weekly", days: 7 },
    { key: "monthly", label: "Monthly", days: 30.44 },
    { key: "quarterly", label: "Quarterly", days: 90 },
    { key: "yearly", label: "Yearly", days: 365 },
  ];

  // Drops floating-point noise (8639.999999999998 -> 8640) before values get
  // truncated to whole seconds for display. 12 significant digits is far more
  // than any input needs.
  function denoise(n) {
    return Number(n.toPrecision(12));
  }

  function periods(hoursPerDay) {
    return PERIODS.map((p) => ({ ...p, seconds: p.days * hoursPerDay * 3600 }));
  }

  // Formats seconds as "1d 2h 3m 4s", or "0s" for zero.
  function secondsToTimeString(total) {
    total = Math.trunc(total);
    const d = Math.floor(total / 86400);
    const h = Math.floor((total % 86400) / 3600);
    const m = Math.floor((total % 3600) / 60);
    const s = total % 60;
    const parts = [];
    if (d > 0) parts.push(d + "d");
    if (h > 0) parts.push(h + "h");
    if (m > 0) parts.push(m + "m");
    if (s > 0) parts.push(s + "s");
    return parts.length ? parts.join(" ") : "0s";
  }

  // Parses "1d2h3m4s" into seconds. Whitespace is ignored so "1h 30m" works.
  function timeStringToSeconds(str) {
    const clean = String(str).replace(/\s+/g, "").toLowerCase();
    const match = /^(?:(\d+)d)?(?:(\d+)h)?(?:(\d+)m)?(?:(\d+)s)?$/.exec(clean);
    if (!clean || !match) {
      throw new Error(`"${str}" is not a valid duration. Use days, hours, minutes and seconds in that order, like 1d, 1h30m or 0h20m.`);
    }
    const [, d = 0, h = 0, m = 0, s = 0] = match;
    return Number(d) * 86400 + Number(h) * 3600 + Number(m) * 60 + Number(s);
  }

  // Accepts "," or "." decimals. A value from 1 up to (not including) 16 is
  // read as a count of nines (1 -> 90, 4 -> 99.99), unless it ends in "%",
  // which always means a plain percentage ("4%" -> 4).
  function parseServiceLevel(str) {
    const trimmed = String(str).trim();
    const explicitPercent = trimmed.endsWith("%");
    const clean = trimmed.replace(",", ".").replace(/%$/, "").trim();
    const value = clean === "" ? NaN : Number(clean);
    if (!Number.isFinite(value)) {
      throw new Error("Enter a service level as a percentage like 99.9, or a number of nines like 3.");
    }
    if (!explicitPercent && value >= 1 && value < 16) {
      return { percent: denoise(100 - Math.pow(10, 2 - value)), nines: value };
    }
    if (value <= 0 || value > 100) {
      throw new Error("A service level has to be above 0% and at most 100%.");
    }
    return { percent: value, nines: null };
  }

  function downtimeBudget(serviceLevel, hoursPerDay) {
    const sl = parseServiceLevel(serviceLevel);
    return {
      ...sl,
      rows: periods(hoursPerDay).map((p) => ({
        ...p,
        downtime: denoise((p.seconds * (100 - sl.percent)) / 100),
      })),
    };
  }

  // Probe frequency:
  // (downtime - MTTR * incidents) / (incidents * probes). A period whose budget
  // doesn't cover the repair time gets frequency null.
  function probeFrequency(rows, mttr, incidents, probes) {
    const repair = timeStringToSeconds(mttr) * incidents;
    return rows.map((row) => ({
      ...row,
      repair,
      frequency: Math.trunc(row.downtime) > repair ? (row.downtime - repair) / (incidents * probes) : null,
    }));
  }

  // Reverse: availability is 0 when downtime fills the period.
  function reverse(outage, hoursPerDay) {
    const downtime = timeStringToSeconds(outage);
    return periods(hoursPerDay).map((p) => ({
      ...p,
      downtime,
      availability: p.seconds > downtime ? 100 - (downtime / p.seconds) * 100 : 0,
    }));
  }

  return { periods, secondsToTimeString, timeStringToSeconds, parseServiceLevel, downtimeBudget, probeFrequency, reverse };
})();

if (typeof module !== "undefined") module.exports = SLH;
