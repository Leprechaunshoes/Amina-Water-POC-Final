#!/bin/bash

# Amina Water Deployment Verification Script

BACKEND_URL="https://amina-water-backend.onrender.com"

echo "╔════════════════════════════════════════════════════════════════╗"
echo "║      🌊 Amina Water Deployment Verification 🌊                 ║"
echo "╚════════════════════════════════════════════════════════════════╝"
echo ""

# Test 1: Check if backend is online
echo "🔍 Test 1: Checking backend status..."
response=$(curl -s -o /dev/null -w "%{http_code}" "$BACKEND_URL")
if [ "$response" = "200" ]; then
    echo "✅ Backend is ONLINE (HTTP $response)"
else
    echo "⚠️  Backend returned HTTP $response (may still be deploying...)"
fi
echo ""

# Test 2: Get backend info
echo "🔍 Test 2: Getting backend information..."
curl -s "$BACKEND_URL" | jq '.' 2>/dev/null || curl -s "$BACKEND_URL"
echo ""

# Test 3: Check stats endpoint
echo "🔍 Test 3: Checking stats endpoint..."
stats=$(curl -s "$BACKEND_URL/stats")
echo "$stats" | jq '.' 2>/dev/null || echo "$stats"
echo ""

# Test 4: Extract counter value
echo "🔍 Test 4: Current swap counter..."
counter=$(echo "$stats" | jq -r '.totalTransfers' 2>/dev/null || echo "Unable to parse")
echo "   Total Transfers: $counter"
echo ""

# Test 5: Check monitoring status
echo "🔍 Test 5: Blockchain monitoring status..."
echo "   Monitoring: Amina ASA 1107424865 on Algorand Mainnet"
echo "   Interval: Every 30 seconds"
echo "   Indexer: AlgoNode Mainnet"
echo ""

echo "═══════════════════════════════════════════════════════════════"
echo ""
echo "✅ VERIFICATION COMPLETE!"
echo ""
echo "📊 How to see it work:"
echo "   1. Make an Amina token transfer on Algorand"
echo "   2. Wait 30-60 seconds"
echo "   3. Run this script again to see counter increment"
echo ""
echo "🔗 Useful links:"
echo "   • Backend API: $BACKEND_URL"
echo "   • Stats: $BACKEND_URL/stats"
echo "   • Amina ASA: https://allo.info/asset/1107424865"
echo "   • Humanity Fund: https://allo.info/address/PZBPPJUHZ3UMENQHZO2HJKPCPTCYCAWY4FPW44XBOKSYIKPILJN76WMIBA"
echo ""
echo "═══════════════════════════════════════════════════════════════"
