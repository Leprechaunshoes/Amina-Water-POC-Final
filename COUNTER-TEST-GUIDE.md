# How to Verify the Swap Counter is Working

## 🎯 Quick Test Methods

### Method 1: Use the Test Page (Easiest)
1. Open `test.html` in your browser
2. The page will automatically check if the backend is connected
3. Click "Simulate +1 Swap" to increment the counter
4. Click "Get Stats" to see the current count
5. Use "Start Auto-Refresh" to watch the counter update in real-time

### Method 2: Check the Live Website
1. Go to your deployed site: `https://YOUR-SITE.onrender.com/fund.html?demo=1`
2. The `?demo=1` parameter shows the demo button
3. Click the "+1 (demo)" button to increment the counter
4. Watch the progress bar and counter update

### Method 3: Use cURL (Terminal)
```bash
# Get current stats
curl https://amina-water-backend.onrender.com/stats

# Simulate a swap (increment counter)
curl -X POST https://amina-water-backend.onrender.com/simulate

# Check stats again to see the increase
curl https://amina-water-backend.onrender.com/stats
```

### Method 4: Browser Console
Open your browser's developer console (F12) and run:
```javascript
// Get current counter
fetch('https://amina-water-backend.onrender.com/stats')
  .then(r => r.json())
  .then(d => console.log('Counter:', d.totalTransfers));

// Simulate a swap
fetch('https://amina-water-backend.onrender.com/simulate', {method: 'POST'})
  .then(() => console.log('Swap simulated!'));
```

## 📊 What the Counter Shows

- **totalTransfers**: Number of swaps/transfers counted
- **lastDonation**: Timestamp of last $100 donation (empty until first milestone)
- **Progress Bar**: Visual representation of progress toward next 1,000 swap milestone

## ✅ Signs It's Working

1. Counter increments when you click the demo button
2. Progress bar moves forward
3. `/stats` endpoint returns updated count
4. Counter persists across page refreshes
5. Auto-refresh on the fund page shows live updates

## 🔧 Troubleshooting

### Backend Not Responding
- Check if backend is deployed on Render
- Wait 1-2 minutes for Render to wake up (free tier sleeps when inactive)
- Check Render dashboard for deployment status

### Counter Not Incrementing
- Check browser console for errors (F12)
- Verify BACKEND constant in scripts.js points to correct URL
- Test the `/simulate` endpoint directly with cURL

### Local Testing
If testing locally:
```bash
# Start the backend
go build -o amina-server main.go
./amina-server &

# Test it
curl http://localhost:8080/stats
curl -X POST http://localhost:8080/simulate
```

## 🚀 Mobile Friendly Updates

The website is now optimized for mobile devices with:
- Responsive text sizing (clamp() for fluid typography)
- Flexible layouts that stack on small screens
- Touch-friendly button sizes
- Optimized spacing for mobile viewports
- Horizontal scrolling prevented on code blocks
- Tested on screens from 320px to 1920px width

Test on mobile by:
1. Opening Chrome DevTools (F12)
2. Click the device toolbar icon (Ctrl+Shift+M)
3. Select different device sizes
4. Or use your actual phone/tablet

## 📱 Deployment Status

Your changes are now live at:
- **Backend API**: https://amina-water-backend.onrender.com
- **Frontend**: Your static site host (GitHub Pages, Netlify, etc.)
- **Test Page**: Upload `test.html` to your static host to use it online

The backend will automatically redeploy when you push changes to the `main` branch.
