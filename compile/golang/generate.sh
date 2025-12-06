#!/usr/bin/env bash
set -e

# --- SETTINGS ---
SCHEMA_URL="https://raw.githubusercontent.com/ExplosionHm/schemas-opticode/main/golang.fbs"
SCHEMA_NAME="golang.fbs"
OUTPUT_DIR="./"
FLATC="flatc"     # assumes flatc is in PATH; change to absolute path if needed

# --- CREATE OUTPUT DIR ---
mkdir -p "$OUTPUT_DIR"

echo "Downloading schema..."
# Works in Windows Git Bash, Linux, macOS
if command -v curl >/dev/null 2>&1; then
    curl -L "$SCHEMA_URL" -o "$SCHEMA_NAME"
elif command -v wget >/dev/null 2>&1; then
    wget -O "$SCHEMA_NAME" "$SCHEMA_URL"
else
    echo "Error: curl or wget required!"
    exit 1
fi

echo "Running flatc..."

$FLATC --go -o "$OUTPUT_DIR" "$SCHEMA_NAME"

echo "Cleaning up..."
rm "$SCHEMA_NAME"

echo "Done!"
