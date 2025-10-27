# Amina Water Backend Deployment

## Deploy to Render (Free)

1. Go to [Render.com](https://render.com) and sign up/sign in
2. Click "New +" and select "Web Service"
3. Connect your GitHub repository: `Leprechaunshoes/Amina-Water-POC-Final`
4. Configure the service:
   - **Name**: `amina-water-backend`
   - **Runtime**: Go
   - **Build Command**: `go build -o amina-server main.go`
   - **Start Command**: `./amina-server`
   - **Instance Type**: Free
5. Click "Create Web Service"

Once deployed, Render will give you a URL like:
`https://amina-water-backend.onrender.com`

## Update Frontend

After deployment, update `scripts.js` line 3 with your actual Render URL:
```javascript
const BACKEND = window.location.hostname === 'localhost' 
  ? "http://localhost:8080" 
  : "https://YOUR-APP-NAME.onrender.com";
```

Then commit and push the change.

## Test Your Deployment

Visit: `https://YOUR-APP-NAME.onrender.com/stats`

You should see: `{"totalTransfers":0,"lastDonation":""}`
