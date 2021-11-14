package plugin

import (
    "fmt"
    "github.com/coredns/caddy"
)

// Register registers your plugin with CoreDNS and allows it to be called when the server is running.
func Register(name string, action caddy.SetupFunc) {
    fmt.Println("plugin register run")
	caddy.RegisterPlugin(name, caddy.Plugin{
		ServerType: "dns",
		Action:     action,
	})
}
