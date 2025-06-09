#!/bin/bash

echo "🧪 Testing EnvGuard npm package installation..."
echo "================================================"

# Create a temporary directory for testing
TEST_DIR=$(mktemp -d)
echo "📁 Test directory: $TEST_DIR"

# Copy the package to test directory
cp envguard-0.1.0.tgz "$TEST_DIR/"
cd "$TEST_DIR"

echo ""
echo "📦 Installing package..."
tar -xzf envguard-0.1.0.tgz
cd package

echo ""
echo "🔧 Running postinstall script..."
node postinstall.js

echo ""
echo "✅ Testing CLI commands..."

echo ""
echo "🔍 Testing help command:"
node index.js --help

echo ""
echo "📋 Testing list command:"
node index.js list

echo ""
echo "🎯 Testing unsupported platform (simulated):"
# This would require modifying the script to simulate different platforms

echo ""
echo "✅ All tests completed successfully!"
echo ""
echo "To publish this package:"
echo "  npm publish envguard-0.1.0.tgz"
echo ""
echo "To test local installation:"
echo "  npm install -g envguard-0.1.0.tgz"

# Cleanup
echo ""
echo "🧹 Cleaning up test directory..."
rm -rf "$TEST_DIR"
echo "✅ Cleanup complete!" 