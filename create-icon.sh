#!/bin/bash
# Create a simple app icon using ImageMagick
# This creates a 256x256 icon with a monitor/screen symbol

if ! command -v convert &> /dev/null; then
    echo "ImageMagick is not installed. Installing..."
    sudo apt install -y imagemagick
fi

# Create a simple icon with a monitor/tablet symbol
convert -size 256x256 xc:none \
    -fill '#2563eb' \
    -draw "roundrectangle 40,50 216,190 15,15" \
    -fill '#1e40af' \
    -draw "rectangle 40,170 216,190" \
    -fill '#60a5fa' \
    -draw "roundrectangle 50,60 206,170 8,8" \
    -fill '#3b82f6' \
    -draw "rectangle 100,190 156,210" \
    -fill '#1e40af' \
    -draw "rectangle 80,210 176,215" \
    build/appicon.png

echo "Icon created at build/appicon.png"
