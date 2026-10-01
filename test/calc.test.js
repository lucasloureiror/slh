const test = require("node:test");
const assert = require("node:assert/strict");
const SLH = require("../docs/calc.js");

const budget = (sl, hours = 24) => SLH.downtimeBudget(sl, hours).rows.map((r) => SLH.secondsToTimeString(r.downtime));

test("timeStringToSeconds parses every unit combination", () => {
  const cases = {
    "1d": 86400, "1h": 3600, "1m": 60, "1s": 1, "1d1h1m1s": 90061, "1d1h1m": 90060, "1d1h1s": 90001,
    "1d1m1s": 86461, "1h1m1s": 3661, "1d1h": 90000, "1d1m": 86460, "1d1s": 86401, "1h1m": 3660,
    "1h1s": 3601, "1m1s": 61, "1h 30m": 5400, "0h20m": 1200,
  };
  for (const [str, seconds] of Object.entries(cases)) assert.equal(SLH.timeStringToSeconds(str), seconds, str);
});

test("timeStringToSeconds rejects malformed durations", () => {
  for (const str of ["", "1d1h1m1", "1d1h1m1s1", "1s1s", "xd", "xm", "xh", "xs", "xdxmxmxs", "1m1h"]) {
    assert.throws(() => SLH.timeStringToSeconds(str), str);
  }
});

test("secondsToTimeString", () => {
  assert.equal(SLH.secondsToTimeString(0), "0s");
  assert.equal(SLH.secondsToTimeString(8.64), "8s");
  assert.equal(SLH.secondsToTimeString(90061), "1d 1h 1m 1s");
  assert.equal(SLH.secondsToTimeString(3600), "1h");
});

test("parseServiceLevel", () => {
  assert.equal(SLH.parseServiceLevel("99.9").percent, 99.9);
  assert.equal(SLH.parseServiceLevel("99,9%").percent, 99.9);
  assert.equal(SLH.parseServiceLevel("100").percent, 100);
  assert.ok(Math.abs(SLH.parseServiceLevel("3").percent - 99.9) < 1e-9);
  assert.equal(SLH.parseServiceLevel("3").nines, 3);
  assert.ok(Math.abs(SLH.parseServiceLevel("5").percent - 99.999) < 1e-9);
  assert.ok(Math.abs(SLH.parseServiceLevel("1").percent - 90) < 1e-9);
  assert.equal(SLH.parseServiceLevel("1").nines, 1);
  assert.equal(SLH.parseServiceLevel("1%").percent, 1);
  assert.equal(SLH.parseServiceLevel("1%").nines, null);
  assert.equal(SLH.parseServiceLevel("5 %").percent, 5);
  assert.equal(SLH.parseServiceLevel("0.5").percent, 0.5);
  assert.equal(SLH.parseServiceLevel("16").percent, 16);
  for (const bad of ["", "abc", "0", "-1", "101", "%", "99%9"]) assert.throws(() => SLH.parseServiceLevel(bad), bad);
});

test("downtime budget for 99.9%", () => {
  assert.deepEqual(budget("99.9"), ["1m 26s", "10m 4s", "43m 50s", "2h 9m 36s", "8h 45m 36s"]);
});

test("downtime budget for 4 nines", () => {
  assert.equal(budget("4")[0], "8s");
  assert.deepEqual(budget("1"), ["2h 24m", "16h 48m", "3d 1h 3m 21s", "9d", "36d 12h"]);
});

test("hours per day shortens every period", () => {
  assert.deepEqual(budget("99", 12), ["7m 12s", "50m 24s", "3h 39m 10s", "10h 48m", "1d 19h 48m"]);
});

test("probe frequency: 99%, 20m MTTR, 3 incidents, 2 probes", () => {
  const rows = SLH.probeFrequency(SLH.downtimeBudget("99", 24).rows, "0h20m", 3, 2);
  assert.equal(rows[0].frequency, null);
  assert.deepEqual(
    rows.slice(1).map((r) => SLH.secondsToTimeString(r.frequency)),
    ["6m 48s", "1h 3m 3s", "3h 26m", "14h 26m"],
  );
});

test("probe frequency is null when repairs use the whole budget", () => {
  const rows = SLH.probeFrequency(SLH.downtimeBudget("99", 24).rows, "1d", 1, 1);
  assert.ok(rows.every((r) => r.frequency === null || r.frequency > 0));
  assert.equal(rows[0].frequency, null);
});

test("probe frequency formula", () => {
  const rows = SLH.probeFrequency([{ downtime: 6048 }], "20m", 3, 2);
  assert.equal(rows[0].frequency, 408);
});

test("reverse", () => {
  const rows = SLH.reverse("8s", 24);
  assert.equal(rows[0].availability.toFixed(2), "99.99");
  const full = SLH.reverse("5h45m30s", 24).map((r) => r.availability.toFixed(4));
  assert.deepEqual(full, ["76.0069", "96.5724", "99.2118", "99.7334", "99.9343"]);
  assert.equal(SLH.reverse("2d", 24)[0].availability, 0);
});
