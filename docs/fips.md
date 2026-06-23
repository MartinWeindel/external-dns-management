# FIPS 140-3 Builds

The controllers can be built so that all cryptographic operations go through Go's
native [FIPS 140-3 module](https://go.dev/doc/security/fips140). The module is
frozen at build time via the `GOFIPS140` environment variable and is activated
in the binary by default — no runtime configuration is required.

Two separate artifacts are produced:

- **Binary** — built with `GOFIPS140=v1.0.0` so the CMVP-certified Go crypto
  module is frozen in at compile time and FIPS mode is on by default.
- **Container image** — built on `gcr.io/distroless/static-debian13:nonroot`,
  the same minimal base as the standard image. Tagged with a `-fips` suffix to
  distinguish it from the standard image.

## How to use it

```bash
# Normal (non-FIPS) build — unchanged:
make release
make docker-images

# FIPS binary only (frozen v1.0.0 CMVP-certified module, the default):
make release-fips

# Pin a specific module version:
make release-fips GOFIPS140=v1.26.0

# FIPS container images (calls release-fips internally, tags with -fips suffix):
make docker-images-fips
```

Supported `GOFIPS140` values include any frozen module version (e.g. `v1.0.0`,
`v1.26.0`), as well as the lifecycle aliases `latest`, `inprocess`, and
`certified`. See the [Go documentation](https://go.dev/doc/security/fips140)
for the meaning of each.

## Verifying a binary is FIPS-enabled

`go version -m` exposes the build settings that were baked into the binary:

```bash
$ go version -m ./dns-controller-manager | grep -E 'GOFIPS140|fips140|DefaultGODEBUG'
        build   -tags=fips140v1.0
        build   DefaultGODEBUG=fips140=on
        build   GOFIPS140=v1.0.0-c2097c7c
```

At runtime, both controllers log their FIPS status on startup:

```
FIPS 140-3 mode: ON (Go crypto module v1.0.0)
```

In Go code, the same information is available via `pkg/fips`:

```go
import "github.com/gardener/external-dns-management/pkg/fips"

if fips.Enabled() {
    // approved-algorithm enforcement is active
}
```

## Base image choice

The FIPS image uses the same `gcr.io/distroless/static-debian13:nonroot` base
as the standard image, not a FIPS-validated OS image such as
`ghcr.io/gardenlinux/gardenlinux-fips`.

The reasoning: the Go binary is **fully static** (`CGO_ENABLED=0`) and never
calls into the system's OpenSSL or any native crypto library. All cryptographic
operations go through the pure-Go FIPS 140-3 module that is baked into the
binary at compile time via `GOFIPS140`. The surrounding container OS has no
influence on which algorithms the binary uses.

Switching to `gardenlinux-fips` would add roughly 120–150 MB of OS user-space
(glibc, the OpenSSL FIPS provider, apt, coreutils, …) that the binary never
touches — larger image, larger attack surface, no functional gain.

## Notes

- `release-fips` keeps `CGO_ENABLED=0`, so the resulting binaries remain fully
  static and do not require libc or OpenSSL in the container.
- Because `GOFIPS140` sets `DefaultGODEBUG=fips140=on` in the binary, FIPS mode
  is on out of the box. Setting `GODEBUG=fips140=on` at runtime is redundant but
  harmless; `GODEBUG=fips140=only` switches the binary into the stricter
  panic-on-non-approved-primitive mode (intended for self-tests, not
  production).
