// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !js && !wasip1

package adns

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"sort"
	"strings"
	"sync"
	"testing"
)

var (
	// uninstallTestHooks runs just before a run of benchmarks.
	testHookUninstaller sync.Once
)

var (
	testDNSFlood = flag.Bool("dnsflood", false, "whether to test DNS query flooding")

	// If external IPv4 connectivity exists, we can try dialing
	// non-node/interface local scope IPv4 addresses.
	// On Windows, Lookup APIs may not return IPv4-related
	// resource records when a node has no external IPv4
	// connectivity.
	testIPv4 = flag.Bool("ipv4", true, "assume external IPv4 connectivity exists")

	// If external IPv6 connectivity exists, we can try dialing
	// non-node/interface local scope IPv6 addresses.
	// On Windows, Lookup APIs may not return IPv6-related
	// resource records when a node has no external IPv6
	// connectivity.
	testIPv6 = flag.Bool("ipv6", false, "assume external IPv6 connectivity exists")
)

func TestMain(m *testing.M) {
	installTestHooks()

	st := m.Run()

	testHookUninstaller.Do(uninstallTestHooks)
	if testing.Verbose() {
		printRunningGoroutines()
	}
	forceCloseSockets()
	os.Exit(st)
}

func printRunningGoroutines() {
	gss := runningGoroutines()
	if len(gss) == 0 {
		return
	}
	fmt.Fprintf(os.Stderr, "Running goroutines:\n")
	for _, gs := range gss {
		fmt.Fprintf(os.Stderr, "%v\n", gs)
	}
	fmt.Fprintf(os.Stderr, "\n")
}

// runningGoroutines returns a list of remaining goroutines.
func runningGoroutines() []string {
	var gss []string
	b := make([]byte, 2<<20)
	b = b[:runtime.Stack(b, true)]
	for _, s := range strings.Split(string(b), "\n\n") {
		_, stack, _ := strings.Cut(s, "\n")
		stack = strings.TrimSpace(stack)
		if !strings.Contains(stack, "created by net") {
			continue
		}
		gss = append(gss, stack)
	}
	sort.Strings(gss)
	return gss
}
