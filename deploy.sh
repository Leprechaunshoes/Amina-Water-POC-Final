#!/bin/bash

# Deploy Amina Water Project
# This script helps you deploy both frontend and backend

echo "╔════════════════════════════════════════════════════════════════╗"
echo "║           🌊 Amina Water Deployment Script 🌊                  ║"
echo "╚════════════════════════════════════════════════════════════════╝"
echo ""

# Check if backend builds
echo "🔧 Step 1: Testing backend build..."
if go build -o amina-server main.go; then
    echo "✅ Backend builds successfully"
else
    echo "❌ Backend build failed!"
    exit 1
fi

# Test backend locally
echo ""
echo "🧪 Step 2: Testing backend locally..."
./amina-server &
SERVER_PID=$!
sleep 2

if curl -s http://localhost:8080/stats > /dev/null; then
    echo "✅ Backend works locally"
    kill $SERVER_PID 2>/dev/null
else
    echo "❌ Backend failed to start"
    kill $SERVER_PID 2>/dev/null
    exit 1
fi

# Check if HTML files exist
echo ""
echo "📄 Step 3: Checking frontend files..."
if [ -f "index.html" ] && [ -f "fund.html" ] && [ -f "scripts.js" ]; then
    echo "✅ All frontend files present"
else
    echo "❌ Missing frontend files"
    exit 1
fi

# Show deployment options
echo ""
echo "═══════════════════════════════════════════════════════════════"
echo ""
echo "✅ Everything is ready to deploy!"
echo ""
echo "📊 DEPLOYMENT OPTIONS:"
echo ""
echo "OPTION 1: Deploy Backend to Render (Recommended)"
echo "  1. Go to https://dashboard.render.com"
echo "  2. Click 'New +' → 'Web Service'"
echo "  3. Connect: Leprechaunshoes/Amina-Water-POC-Final"
echo "  4. Render will auto-detect render.yaml"
echo "  5. Click 'Create Web Service'"
echo "  6. Wait 3-5 minutes"
echo "  7. Backend URL: https://amina-water-backend.onrender.com"
echo ""
echo "OPTION 2: Deploy Frontend to GitHub Pages"
echo "  1. Push to main branch (already done)"
echo "  2. Go to GitHub repo → Settings → Pages"
echo "  3. Source: Deploy from branch 'main'"
echo "  4. Folder: / (root)"
echo "  5. Save and wait 1-2 minutes"
echo "  6. Frontend URL: https://leprechaunshoes.github.io/Amina-Water-POC-Final/"
echo ""
echo "OPTION 3: Deploy Both to Render (More Complex)"
echo "  See RENDER-SETUP.md for detailed instructions"
echo ""
echo "═══════════════════════════════════════════════════════════════"
echo ""
echo "🎯 RECOMMENDED APPROACH:"
echo "  • Backend API → Render (handles blockchain monitoring)"
echo "  • Frontend → GitHub Pages (free, fast, simple)"
echo ""
echo "The frontend will automatically connect to the backend!"
echo ""
echo "═══════════════════════════════════════════════════════════════"
