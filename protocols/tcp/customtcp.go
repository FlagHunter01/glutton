// [CUSTOM] Generic TCP handler

package tcp

import (
	"context"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/mushorg/glutton/connection"
	"github.com/mushorg/glutton/protocols/interfaces"
)


func CustomTcpHandle(ctx context.Context, conn net.Conn, md connection.Metadata, logger interfaces.Logger, h interfaces.Honeypot) error {
	
	// Source IP
	source_ip, _,_ := net.SplitHostPort(conn.RemoteAddr().String())

	// Destination port
	destination_port := md.TargetPort

	// Log connection data
	log.SetOutput(os.Stdout)
	log.Printf("TCP connection from %-15s to port %s\n", source_ip, strconv.FormatInt(int64(destination_port), 10))

	return conn.Close()
}
