#!/usr/bin/env bash
# SPDX-FileCopyrightText: SAP SE or an SAP affiliate company and Gardener contributors
#
# SPDX-License-Identifier: Apache-2.0

set -euo pipefail
BIN="${1:?usage: fips-audit.sh <binary>}"

echo "== go version -m =="
go version -m "$BIN" | grep -E 'GOFIPS140|GOEXPERIMENT|CGO_ENABLED' || true

echo "== dynamic deps =="
if command -v ldd >/dev/null 2>&1; then
  ldd "$BIN" 2>&1 || true
elif command -v otool >/dev/null 2>&1; then
  # macOS: ldd doesn't exist; otool -L is the equivalent
  otool -L "$BIN" 2>&1 || true
else
  echo "  skipped (neither ldd nor otool available)"
fi

echo "== native crypto symbols (should be empty) =="
# Match real native-crypto indicators only. The Go stdlib has a stub package
# `crypto/internal/boring` whose symbols (e.g. crypto/internal/boring.* ,
# runtime.boringCaches) appear in every Go binary as no-op placeholders — those
# are NOT BoringCrypto being linked in. Real BoringCrypto would show
# `_goboringcrypto_*` or `_cgo_*boring*` symbols.
go tool nm "$BIN" 2>/dev/null \
  | grep -Ei '_goboringcrypto_|_cgo_.*boring|openssl|libcrypto|EVP_[A-Z]|sodium|wolfssl|mbedtls|gnutls' \
  | grep -v 'crypto/internal/boring\.\|runtime\.boringCaches' \
  || echo "  none"

echo "== embedded crypto-library strings =="
strings "$BIN" | grep -Ei 'OpenSSL [0-9]|BoringSSL|libsodium|wolfSSL|mbed TLS|GnuTLS' | sort -u || echo "  none"

echo "== disallowed-primitive source references =="
grep -rEn --include='*.go' 'crypto/md5|crypto/sha1\b|crypto/des|crypto/rc4|"math/rand"' . || echo "  none"