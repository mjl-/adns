// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !js && !wasip1

package adns

import (
	"fmt"
	"io/fs"
	"net"
)

func parseLookupPortError(nestedErr error) error {
	if nestedErr == nil {
		return nil
	}

	switch nestedErr.(type) {
	case *net.AddrError, *DNSError:
		return nil
	case *fs.PathError: // for Plan 9
		return nil
	}
	return fmt.Errorf("unexpected type on 1st nested level: %T", nestedErr)
}
