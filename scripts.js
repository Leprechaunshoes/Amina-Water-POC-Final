// Auto-detect environment - use same server for everything
const BACKEND = window.location.origin;

async function load(){
  try {
    const s = await fetch(BACKEND + "/stats").then(r => r.json());
    const d = await fetch(BACKEND + "/last-donation").then(r => r.json());

    const total = s.totalTransfers || 0;
    const goal = 1000;

    // Check if we're on fund.html (has a span#count) or index.html (has div#count)
    const countEl = document.getElementById("count");
    if (countEl) {
      // If it's a span (fund.html), just set the number
      if (countEl.tagName === 'SPAN') {
        countEl.textContent = total.toString();
      } else {
        // If it's a div (index.html), set the full format
        countEl.textContent = `${total} / ${goal}`;
      }
    }
    
    const barEl = document.getElementById("bar");
    if (barEl) {
      barEl.style.width = Math.min(100, (total / goal) * 100) + "%";
    }

    const lastEl = document.getElementById("last");
    if (lastEl && d.txnId) {
      lastEl.innerHTML =
        `<a href="https://allo.info/mainnet/transaction/${d.txnId}" target="_blank">${d.txnId.slice(0,12)}…</a>`;
    }
  } catch(e) {
    console.log("Backend not available:", e);
    // Site will still work, just won't show live stats
  }
}

load();
setInterval(load, 15000);

