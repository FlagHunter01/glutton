// [CUSTOM] Generic UDP handler

package udp

import (
	"context"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/mushorg/glutton/connection"
	"github.com/mushorg/glutton/protocols/interfaces"
)

func CustomUdpHandle(ctx context.Context, srcAddr, dstAddr *net.UDPAddr, data []byte, md connection.Metadata, logger interfaces.Logger, h interfaces.Honeypot) error {
	
	// Source IP
	source_ip := srcAddr.IP

	// Destination port
	destination_port := md.TargetPort

	// Log connection data
	log.SetOutput(os.Stdout)
	log.Printf("UDP connection from %-15s to port %s\n", source_ip, strconv.FormatInt(int64(destination_port), 10))

	return nil
}
