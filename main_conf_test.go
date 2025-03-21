// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package adns

import (
	"testing"
)

func allResolvers(t *testing.T, f func(t *testing.T)) {
	t.Run("default resolver", f)
	t.Run("forced go resolver", func(t *testing.T) {
		if fixup := forceGoDNS(); fixup != nil {
			defer fixup()
			f(t)
		}
	})
	t.Run("forced cgo resolver", func(t *testing.T) {
		if fixup := forceCgoDNS(); fixup != nil {
			defer fixup()
			f(t)
		}
	})
}

// See main_conf_test.go for what these (don't) do.
func forceGoDNS() func() { return nil }

// See main_conf_test.go for what these (don't) do.
func forceCgoDNS() func() { return nil }
