# Render Deployment Setup

Your code is now configured for **TWO separate services** on Render:

## 🔧 Service 1: Backend API (Go Server)

This runs your Go server that monitors the blockchain.

**Service Name**: `amina-water-backend`
**URL**: https://amina-water-backend.onrender.com

### To Deploy:
1. Go to https://dashboard.render.com
2. Click "New +" → "Web Service"
3. Connect your GitHub repo: `Leprechaunshoes/Amina-Water-POC-Final`
4. Render will auto-detect the `render.yaml` and create the backend service
5. Wait for deployment (2-3 minutes)

**Test it**:
```bash
curl https://amina-water-backend.onrender.com/stats
```

## 🌐 Service 2: Frontend Website (Static Site)

This serves your HTML pages.

**Service Name**: `amina-water-site`  
**URL**: https://amina-water-site.onrender.com (or custom domain)

### To Deploy:
1. In Render dashboard, the frontend service should auto-create from `render.yaml`
2. If not, click "New +" → "Static Site"
3. Connect same GitHub repo
4. Build Command: `mkdir -p public && cp index.html fund.html style.css scripts.js coin.PNG public/`
5. Publish Directory: `public`
6. Auto-deploy: Yes

## ✅ Verification

Once both are deployed:

1. **Backend API**: https://amina-water-backend.onrender.com/stats
   - Should return: `{"totalTransfers":0,"lastDonation":""}`

2. **Frontend Site**: https://amina-water-site.onrender.com
   - Should show your Amina Water website
   - Counter should update from backend

## 🔄 How It Works Together

```
User visits Frontend
     ↓
Frontend (scripts.js) calls Backend API
     ↓
Backend monitors Algorand blockchain
     ↓
Backend returns swap counter
     ↓
Frontend displays the count
```

## 🚨 If You Only Want ONE Service

If you prefer everything on one URL, use **GitHub Pages or Netlify** for the frontend:

1. **Frontend**: Deploy to GitHub Pages (free, easy)
2. **Backend**: Keep on Render (handles the API)

### Quick GitHub Pages Setup:
```bash
# Create gh-pages branch
git checkout -b gh-pages
git push -u origin gh-pages
```

Then in GitHub repo settings → Pages → Source: gh-pages branch

Your site will be at: https://leprechaunshoes.github.io/Amina-Water-POC-Final/

## 📝 Current Status

After pushing the latest changes, Render will:
1. Auto-deploy the backend API
2. Auto-deploy the frontend static site

**Wait 3-5 minutes** for both services to deploy.

## 🧪 Quick Test

Run this after 5 minutes:
```bash
./verify-deployment.sh
```

If backend shows "Not Found", check Render dashboard for deployment logs.
