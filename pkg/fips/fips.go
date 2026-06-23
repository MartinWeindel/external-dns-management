// SPDX-FileCopyrightText: SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

// Package fips reports whether the Go FIPS 140-3 crypto module is active at
// runtime. The module is baked into the binary at build time via GOFIPS140
// and activated at process start via GODEBUG=fips140=on (or "only").
//
// LogStatus is intended to be called once during startup so the operational
// state is visible in pod logs — useful both for proving FIPS mode is on in
// release builds and for catching accidental regressions in non-FIPS builds.
package fips

import (
	"crypto/fips140"
	"fmt"
)

// Enabled reports whether the FIPS 140-3 module is currently enforcing
// approved-algorithm selection. It returns true only when the binary was
// built with GOFIPS140 set AND the process was started with
// GODEBUG=fips140=on or GODEBUG=fips140=only.
func Enabled() bool {
	return fips140.Enabled()
}

// Status returns a single human-readable line describing the FIPS state of
// the running binary. Suitable for an INFO log at startup.
func Status() string {
	if fips140.Enabled() {
		return fmt.Sprintf("FIPS 140-3 mode: ON (Go crypto module %s)", fips140.Version())
	}
	return "FIPS 140-3 mode: off"
}
