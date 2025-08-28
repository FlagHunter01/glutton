package custom

import (
	"context"
	"net"
	"strings"
	"time"
)

// Returns the hostname associated with the given IP, or the IP if it fails
func ResolveIp(ip string) string{
	// If not IP format, return it as-is
	if net.ParseIP(ip) == nil {
		return ip
	}

	// resolver uses the system DNS configuration by default.
	var resolver = net.Resolver{}

	// Context with 5 minute timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	// Resolve IP
	names, err := resolver.LookupAddr(ctx, ip)
	// If an error occurs, return IP
	if err != nil || len(names) == 0 {
		return ip
	}

	// Hostname
	hostname := strings.TrimSpace(names[0])
	// Remove ".in-addr.arpa." suffix if present
	hostname = strings.TrimSuffix(hostname, ".in-addr.arpa.")
	// Remove trailing dot if present
	hostname = strings.TrimSuffix(hostname, ".")
	// Add the IP address in parenthesis
	hostname += " (" + ip + ")"

	return hostname
}
