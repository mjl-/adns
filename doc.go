/*
adns is a copy of the Go standard library, modified to provide details about
the DNSSEC status of responses.

The MX, NS, SRV types from the "net" package are used to make to prevent churn
when switching from net to adns.

Modifications

  - Each Lookup* also returns a Result with the "Authentic" field representing if
    the response had the "authentic data" bit (and is trusted), i.e. was
    DNSSEC-signed according to the recursive resolver.
*/
package adns
