# FIPS 140-3 Builds

The controllers can be built so that all cryptographic operations go through Go's
native [FIPS 140-3 module](https://go.dev/doc/security/fips140). The module is
frozen at build time via the `GOFIPS140` environment variable and is activated
in the binary by default — no runtime configuration is required.

This page documents only how to *build* a FIPS binary. The matching
FIPS-validated container image is built on top of
`ghcr.io/gardenlinux/gardenlinux-fips` and is described separately.

## How to use it

```bash
# Normal (non-FIPS) build — unchanged:
make release

# FIPS build (frozen v1.0.0 CMVP-certified module, the default):
make release-fips

# Pin a specific module version:
make release-fips GOFIPS140=v1.26.0
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

## Notes

- `release-fips` keeps `CGO_ENABLED=0`, so the resulting binaries remain fully
  static and work in `scratch` / `distroless` / `gardenlinux-fips` base images
  without any libc or OpenSSL dependency.
- Because `GOFIPS140` sets `DefaultGODEBUG=fips140=on` in the binary, FIPS mode
  is on out of the box. Setting `GODEBUG=fips140=on` at runtime is redundant but
  harmless; `GODEBUG=fips140=only` switches the binary into the stricter
  panic-on-non-approved-primitive mode (intended for self-tests, not
  production).
