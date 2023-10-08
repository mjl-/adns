// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix

package adns

var (
	extraTestHookInstallers   []func()
	extraTestHookUninstallers []func()
)

func installTestHooks() {
	for _, fn := range extraTestHookInstallers {
		fn()
	}
}

func uninstallTestHooks() {
	for _, fn := range extraTestHookUninstallers {
		fn()
	}
}

// forceCloseSockets must be called only from TestMain.
func forceCloseSockets() {
}
