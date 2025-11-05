// Auto-detect environment - use same server for everything
const BACKEND = window.location.origin;

async function load(){
  try {
    const s = await fetch(BACKEND + "/stats").then(r => r.json());
    const d = await fetch(BACKEND + "/last-donation").then(r => r.json());

    const total = s.totalTransfers || 0;
    const goal = 1000;

    document.getElementById("count").textContent = `${total} / ${goal}`;
    document.getElementById("bar").style.width = Math.min(100, (total / goal) * 100) + "%";

    if (d.txnId) {
      document.getElementById("last").innerHTML =
        `<a href="https://allo.info/mainnet/transaction/${d.txnId}" target="_blank">${d.txnId.slice(0,12)}…</a>`;
    }
  } catch(e) {
    console.log("Backend not available:", e);
    // Site will still work, just won't show live stats
  }
}

load();
setInterval(load, 15000);

