#!/usr/bin/env bash
set -euo pipefail

BINARY="michi"
DEST="$HOME/.local/bin/$BINARY"

# Stop the service if running
if command -v "$DEST" &>/dev/null; then
    echo "Stopping $BINARY service..."
    "$DEST" stop || echo "$BINARY stop command failed or service not running"
else
    echo "$BINARY binary not found, skipping stop"
fi

# Remove binary
if [ -f "$DEST" ]; then
    echo "Removing $BINARY from $DEST"
    rm -f "$DEST"
    echo "$BINARY removed"
else
    echo "$BINARY not found at $DEST, nothing to do"
    echo "Perhaps it was already removed?"
    echo "If you wish to install it, run the install script."
    exit 0
fi

echo "Note: All configuration under ~/.michi still exist. Manually remove them if needed."
echo "Thank you for using michi!"
