(() => {
  const $ = (id) => document.getElementById(id);
  const fmt = SLH.secondsToTimeString;
  const els = {
    form: $("form"), sl: $("sl"), out: { d: $("out-d"), h: $("out-h"), m: $("out-m"), s: $("out-s") }, hours: $("hours"),
    probeOn: $("probe-on"), probeFields: $("probe-fields"), mttr: $("mttr"), incidents: $("incidents"), probes: $("probes"),
    headline: $("headline"), summary: $("summary"), error: $("error"), table: $("table"), legend: $("legend"),
  };
  let mode = "budget";

  const esc = (s) => String(s).replace(/[&<>"]/g, (c) => ({ "&": "&amp;", "<": "&lt;", ">": "&gt;", '"': "&quot;" }[c]));
  const lengthLabel = (p, hours) => p.key === "daily" ? `${hours} hours` : p.key === "monthly" ? "30.44 days" : `${p.days} days`;
  const formatPercent = (n) => parseFloat(n.toFixed(6)).toString() + "%";

  function positiveInt(input, name, max) {
    const n = Number(input.value);
    const ok = Number.isInteger(n) && n >= 1 && (!max || n <= max);
    input.setAttribute("aria-invalid", String(!ok));
    if (!ok) throw new Error(max ? `${name} has to be a whole number from 1 to ${max}.` : `${name} has to be a whole number of at least 1.`);
    return n;
  }

  function withInvalid(input, fn) {
    try { const v = fn(); input.setAttribute("aria-invalid", "false"); return v; }
    catch (e) { input.setAttribute("aria-invalid", "true"); throw e; }
  }

  // The four fields as a duration string like "5h45m30s". Blank counts as 0.
  function outageString() {
    let text = "", bad = null;
    for (const [unit, input] of Object.entries(els.out)) {
      const n = input.value.trim() === "" ? 0 : Number(input.value);
      const ok = Number.isInteger(n) && n >= 0;
      input.setAttribute("aria-invalid", String(!ok));
      if (!ok) bad = bad || input;
      if (n > 0) text += n + unit;
    }
    if (bad) throw new Error("Days, hours, minutes and seconds have to be whole numbers, zero or more.");
    return text || "0s";
  }

  function setOutage(text) {
    const seconds = SLH.timeStringToSeconds(text);
    const parts = { d: Math.floor(seconds / 86400), h: Math.floor((seconds % 86400) / 3600), m: Math.floor((seconds % 3600) / 60), s: seconds % 60 };
    for (const unit in parts) els.out[unit].value = parts[unit];
  }

  function renderBudget(hours) {
    const budget = withInvalid(els.sl, () => SLH.downtimeBudget(els.sl.value, hours));
    let rows = budget.rows;
    const probing = els.probeOn.checked;
    if (probing) {
      const incidents = positiveInt(els.incidents, "Incidents");
      const probes = positiveInt(els.probes, "Probes to alert");
      rows = withInvalid(els.mttr, () => SLH.probeFrequency(rows, els.mttr.value, incidents, probes));
    }

    const pct = formatPercent(budget.percent);
    els.summary.innerHTML = `<span class="big">${esc(pct)}</span><span class="sub">${budget.nines !== null ? esc(budget.nines) + (budget.nines === 1 ? " nine · " : " nines · ") : ""}maximum downtime per period</span>`;

    let html = `<thead><tr><th>Period</th><th>Downtime budget</th>${probing ? "<th>Probe at least every</th>" : ""}</tr></thead><tbody>`;
    for (const r of rows) {
      html += `<tr><td><span class="period">${r.label}</span><span class="len">${lengthLabel(r, hours)}</span></td><td class="num">${fmt(r.downtime)}</td>`;
      if (probing) {
        const share = r.downtime > 0 ? Math.min(r.repair / r.downtime, 1) : 1;
        const bar = `<div class="bar" aria-hidden="true"><i class="repair" style="width:${(share * 100).toFixed(2)}%"></i><i class="buffer" style="width:${((1 - share) * 100).toFixed(2)}%"></i></div>`;
        if (r.frequency === null) {
          html += `<td class="num na">Not achievable: repairs need ${fmt(r.repair)}${bar}</td>`;
        } else {
          html += `<td class="num">${r.frequency < 1 ? "under 1s" : fmt(r.frequency)}${bar}</td>`;
        }
      }
      html += "</tr>";
    }
    els.table.innerHTML = html + "</tbody>";
    els.legend.hidden = !probing;

    if (probing && rows.every((r) => r.frequency === null)) {
      showError("With this MTTR and incident count, repairs alone use up the budget in every period. No probing frequency can keep this service level.");
    }
  }

  function renderReverse(hours) {
    const rows = SLH.reverse(outageString(), hours);
    els.summary.innerHTML = `<span class="big">${esc(fmt(rows[0].downtime))}</span><span class="sub">of downtime leaves this availability</span>`;
    let html = "<thead><tr><th>Period</th><th>Availability</th></tr></thead><tbody>";
    for (const r of rows) {
      const cell = r.availability > 0
        ? `<td class="num">${r.availability.toFixed(4)}%</td>`
        : `<td class="num na">0%: the downtime is longer than the period</td>`;
      html += `<tr><td><span class="period">${r.label}</span><span class="len">${lengthLabel(r, hours)}</span></td>${cell}</tr>`;
    }
    els.table.innerHTML = html + "</tbody>";
    els.legend.hidden = true;
  }

  function showError(message) {
    els.error.textContent = message;
    els.error.hidden = false;
  }

  function render() {
    els.error.hidden = true;
    els.probeFields.hidden = !els.probeOn.checked;
    document.querySelectorAll("[data-for]").forEach((el) => { el.hidden = el.dataset.for !== mode; });
    document.querySelectorAll(".modes button").forEach((b) => b.setAttribute("aria-selected", String(b.dataset.mode === mode)));
    let current = null;
    try { current = SLH.parseServiceLevel(els.sl.value).percent; } catch (_) { /* nothing is pressed while the input is invalid */ }
    document.querySelectorAll("#presets button").forEach((b) => b.setAttribute("aria-pressed", String(current !== null && Math.abs(current - Number(b.dataset.sl)) < 1e-9)));

    renderHeadline(current);

    try {
      const hours = positiveInt(els.hours, "Hours per day", 24);
      if (mode === "reverse") renderReverse(hours); else renderBudget(hours);
    } catch (e) {
      els.summary.innerHTML = "";
      els.table.innerHTML = "";
      els.legend.hidden = true;
      showError(e.message);
    }
    saveToUrl();
  }

  // The page heading and tab title follow the input. While the input is
  // invalid they keep the last valid value, so the heading doesn't flicker.
  function renderHeadline(percent) {
    let value, html, title;
    if (mode === "reverse") {
      try { value = fmt(SLH.timeStringToSeconds(outageString())); } catch (_) { return; }
      html = `What does <em>${esc(value)}</em> of downtime leave?`;
      title = `${value} of downtime`;
    } else {
      if (percent === null) return;
      value = formatPercent(percent);
      html = `How much downtime does <em>${esc(value)}</em> allow?`;
      title = `${value} downtime budget`;
    }
    els.headline.innerHTML = html;
    document.title = `${title} · Service Level Helper`;
  }

  function saveToUrl() {
    const p = new URLSearchParams();
    if (mode === "reverse") { p.set("mode", "reverse"); try { p.set("outage", outageString()); } catch (_) { /* leave it out while the fields are invalid */ } }
    else {
      p.set("sl", els.sl.value.trim());
      if (els.probeOn.checked) { p.set("mttr", els.mttr.value.trim()); p.set("incidents", els.incidents.value); p.set("probes", els.probes.value); }
    }
    if (els.hours.value !== "24") p.set("hours", els.hours.value);
    try { history.replaceState(null, "", "?" + p.toString() + location.hash); } catch (_) { /* sandboxed viewers */ }
  }

  function loadFromUrl() {
    let p;
    try { p = new URLSearchParams(location.search); } catch (_) { return; }
    if (p.get("mode") === "reverse") mode = "reverse";
    for (const key of ["sl", "hours", "mttr", "incidents", "probes"]) {
      if (p.has(key)) els[key].value = p.get(key);
    }
    if (p.has("outage")) { try { setOutage(p.get("outage")); } catch (_) { /* keep the defaults */ } }
    if (p.has("mttr")) els.probeOn.checked = true;
  }

  document.querySelectorAll(".modes button").forEach((b) => b.addEventListener("click", () => { mode = b.dataset.mode; render(); }));
  document.querySelectorAll("#presets button").forEach((b) => b.addEventListener("click", () => { els.sl.value = b.dataset.sl; render(); }));
  els.form.addEventListener("input", render);
  els.form.addEventListener("submit", (e) => e.preventDefault());

  loadFromUrl();
  render();
})();
