
#!/usr/bin/env bash
set -euo pipefail

REPO="OrbitalJin/michi"
BINARY="michi"

# Detect OS
OS="$(uname -s | tr '[:upper:]' '[:lower:]')"
ARCH="$(uname -m)"

case "$ARCH" in
    x86_64) ARCH="amd64" ;;
    aarch64 | arm64) ARCH="arm64" ;;
    *) echo "Unsupported architecture: $ARCH"; exit 1 ;;
esac

VERSION="${VERSION:-latest}"

if [ "$VERSION" = "latest" ]; then
    VERSION=$(curl -fsSL "https://api.github.com/repos/$REPO/releases/latest" | grep -oP '"tag_name": "\K(.*)(?=")')
fi

echo "Installing $BINARY $VERSION for $OS/$ARCH"

# Build download URL (michi_${OS}_${ARCH}.tar.gz)
TARFILE="${BINARY}_${OS}_${ARCH}.tar.gz"
URL="https://github.com/$REPO/releases/download/$VERSION/$TARFILE"
echo $URL

TMPDIR=$(mktemp -d)
curl -fsSL "$URL" -o "$TMPDIR/$TARFILE"

# Extract
tar -xzf "$TMPDIR/$TARFILE" -C "$TMPDIR"

# Install
DEST="$HOME/.local/bin/$BINARY"
mkdir -p "$(dirname "$DEST")"

mv "$TMPDIR/$BINARY" "$DEST"
chmod +x "$DEST"

echo "Installed $BINARY to $DEST"
