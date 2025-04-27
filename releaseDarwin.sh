#!/bin/bash

# Stop on first error
set -e

# CONFIGURATION
APP_NAME="SynkRip"                          # Name of your app (no .app)
SIGN_IDENTITY="SynkRip"                     # Signing identity name (from Keychain) or security find-identity -v -p codesigning
BUILD_DIR="build/bin"                       # Wails default output directory
OUTPUT_APP="$BUILD_DIR/$APP_NAME.app"
DMG_NAME="$APP_NAME.dmg"
DMG_OUTPUT_DIR="$BUILD_DIR"                 # Where to put final .dmg

# 1. Build the app with Wails
echo "🔨 Building app with Wails..."
wails build

# 2. Sign the .app bundle
echo "🔏 Signing the app..."
codesign --deep --force --verify --verbose --sign "$SIGN_IDENTITY" "$OUTPUT_APP"

# 3. Create a DMG
echo "📦 Packaging into DMG..."
mkdir -p "$DMG_OUTPUT_DIR"
hdiutil create -volname "$APP_NAME" -srcFolder "$OUTPUT_APP" -ov -format UDZO "$DMG_OUTPUT_DIR/$DMG_NAME"

echo "✅ Done! DMG is at: $DMG_OUTPUT_DIR/$DMG_NAME"
