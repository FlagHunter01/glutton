// [CUSTOM] Generic TCP handler

package tcp

import (
	"bufio"
	"context"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/mushorg/glutton/connection"
	"github.com/mushorg/glutton/custom"
	"github.com/mushorg/glutton/protocols/interfaces"
)

// Read file and return map of values
func ReadWhitelist(path string) (map[string]bool, error) {
	// Map containing IPs and ports
	whitelist := make(map[string]bool)
	// Open the whitelist file for reading
	file, error := os.Open(path)
	// Return file error
	if error != nil {
		return nil, error
	}
	// Close the file at the end of the scope
	defer file.Close()

	// Read file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Trim spaces
		line := strings.TrimSpace(scanner.Text())
		// Skip empty lines or comment lines starting with "#"
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		// Add line to map
		whitelist[line] = true
	}

	return whitelist, scanner.Err()
}

// Check that the IP, port or combination of both is whitelisted
func Whitelisted(ip string, port string, whitelist map[string]bool) bool {
	// Combination of IP:port
	str := ip + ":" + port
	if whitelist[str] {
		return true
	}

	// IP only
	if whitelist[ip] {
		return true
	}

	// Port only
	if whitelist[port] {
		return true
	}

	return false
}

func CustomTcpHandle(ctx context.Context, conn net.Conn, md connection.Metadata, logger interfaces.Logger, h interfaces.Honeypot) error {
	// Whitelist and related error
	whitelist, whitelist_error := ReadWhitelist("/root/llamapot/whitelist")
	// Source IP
	source_ip, _, _ := net.SplitHostPort(conn.RemoteAddr().String())
	// Destination port
	destination_port := md.TargetPort
	// Desctination port converted to string
	destination_port_string := strconv.FormatInt(int64(md.TargetPort), 10)
	// Log connection only if not whitelisted
	if !Whitelisted(source_ip, destination_port_string, whitelist) {
		// Set output to stdout
		log.SetOutput(os.Stdout)
		// Write log
		log.Printf("TCP connection from %-40s to port %-5s %s\n", custom.ResolveIp(source_ip), destination_port_string, custom.PortDescription(destination_port))
	}
	// Log error
	if whitelist_error != nil {
		// Set output to stderr
		log.SetOutput(os.Stderr)
		// Write log
		log.Printf("Error while reading whitelist: %s\n", whitelist_error)
	}

	return conn.Close()
}
