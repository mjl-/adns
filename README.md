adns - copy of pure Go resolver from Go standard library, with modifications to facilitate use with DNSSEC.

Commits from Go "net" package not included:
- https://github.com/golang/go/commit/f77bba43aa223fc86fd223f3ea4ef60db8e0c583, "net: accept a valid IP address in LookupMX": let's try to not accept invalid MX records as long as possible...

Documentation: https://pkg.go.dev/github.com/mjl-/adns

License: Go's BSD license, see LICENSE.
