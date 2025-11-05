# Amina Water - Turning Digital Value into Human Impact

A transparent blockchain-based system that automatically donates to clean water projects based on real Amina token swaps.

## 🌊 How It Works

Every **1,000 verified Amina transfers** on the Algorand blockchain triggers a **$100 ALGO donation** from the Humanity Fund to a live clean-water project. All donations are recorded on-chain for complete transparency.

## 🔍 Real Blockchain Monitoring

The backend monitors the Algorand blockchain in real-time:
- **Amina ASA ID**: 1107424865
- **Network**: Algorand Mainnet
- **Monitoring**: Real transfers every 30 seconds via Algorand Indexer
- **Transparency**: All swaps are verifiable on-chain

## 🚀 Live Deployment

- **Backend API**: https://amina-water-backend.onrender.com
- **Monitoring**: Automatic detection of Amina token transfers
- **Updates**: Counter updates every 30 seconds with real blockchain data

## 📊 API Endpoints

### GET /stats
Returns current swap counter and donation status
```json
{
  "totalTransfers": 245,
  "lastDonation": "2025-11-05 12:34:56"
}
```

### GET /last-donation
Returns details of the most recent donation
```json
{
  "txnId": "ABC123...",
  "date": "2025-11-05 12:34:56"
}
```

### POST /simulate
For testing purposes only - manually increment counter

## 🧪 How to Verify It's Working

### Method 1: Check Live Stats
```bash
curl https://amina-water-backend.onrender.com/stats
```

### Method 2: Monitor Real Swaps
1. Make an Amina token transfer on Algorand
2. Wait ~30-60 seconds for the monitor to detect it
3. Check `/stats` endpoint to see counter increment

### Method 3: View on Blockchain
- Check Amina ASA transactions: https://allo.info/asset/1107424865
- Verify Humanity Fund: https://allo.info/address/PZBPPJUHZ3UMENQHZO2HJKPCPTCYCAWY4FPW44XBOKSYIKPILJN76WMIBA

### Method 4: Use Test Page
Open `test.html` in your browser for a comprehensive testing interface

## 🛠️ Local Development

### Prerequisites
- Go 1.16 or higher
- Git

### Run Locally
```bash
# Build the server
go build -o amina-server main.go

# Run the server
./amina-server

# Server runs on http://localhost:8080
```

### Test Locally
```bash
# Check stats
curl http://localhost:8080/stats

# Simulate a swap (testing only)
curl -X POST http://localhost:8080/simulate
```

## 📱 Mobile-Friendly Design

The website is fully responsive with:
- Fluid typography using CSS clamp()
- Touch-friendly buttons (140px+ minimum)
- Responsive layouts that stack on mobile
- Optimized for screens 320px to 1920px wide

## 🔧 Technical Stack

- **Backend**: Go (Golang)
- **Blockchain**: Algorand
- **API**: Algorand Indexer (AlgoNode)
- **Deployment**: Render.com
- **Frontend**: HTML, CSS, Vanilla JavaScript

## 📈 How the Counter Works

1. **Blockchain Monitor**: Queries Algorand Indexer API every 30 seconds
2. **Detection**: Identifies Amina ASA (1107424865) transfer transactions
3. **Counting**: Increments counter for each detected transfer
4. **Milestone**: At 1,000 transfers, triggers donation and resets counter
5. **Transparency**: All transactions are verifiable on Algorand blockchain

## 🎯 Humanity Fund

- **Address**: `PZBPPJUHZ3UMENQHZO2HJKPCPTCYCAWY4FPW44XBOKSYIKPILJN76WMIBA`
- **Purpose**: Holds funds for clean water project donations
- **Transparency**: All inbound/outbound transactions are public
- **Verification**: Anyone can audit the fund on Algorand explorers

## 📄 Files

- `index.html` - Main landing page
- `fund.html` - Humanity Fund transparency page with live counter
- `test.html` - Testing interface for developers
- `style.css` - Responsive styling
- `scripts.js` - Frontend JavaScript
- `main.go` - Backend server with blockchain monitoring
- `render.yaml` - Deployment configuration

## 🚀 Deployment

The project auto-deploys to Render when changes are pushed to the `main` branch:

```bash
git add .
git commit -m "Your changes"
git push origin main
```

Render will automatically:
1. Detect the changes
2. Build the Go application
3. Deploy the backend
4. Start monitoring the blockchain

## 📞 Support

For issues or questions about:
- **Blockchain monitoring**: Check Algorand Indexer status
- **API errors**: Verify backend logs on Render dashboard
- **Counter not updating**: Ensure real Amina transfers are happening

## 📜 License

MIT License - Feel free to use and modify for similar charitable projects!

---

© 2025 Amina Water • Transparency by Design • Powered by Algorand
