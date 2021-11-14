// Package core registers the server and all plugins we support.
package core

import (
	// plug in the server
	_ "github.com/coredns/coredns/core/dnsserver"
    "fmt"
)

func init()  {
    fmt.Println("core/coredns.go run")
}
