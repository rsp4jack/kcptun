package main

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/tarm/serial"
	kcp "github.com/xtaci/kcp-go/v5"
	"github.com/xtaci/kcptun/generic"
	"github.com/xtaci/serialpacket"
	"github.com/xtaci/tcpraw"
)

func dial(config *Config, block kcp.BlockCrypt) (*kcp.UDPSession, error) {
	// serial://NAME
	if strings.HasPrefix(config.RemoteAddr, generic.SERIAL_PROTO) {
		c, err := generic.ParseSerialParams(config.RemoteAddr)
		if err != nil {
			return nil, err
		}

		// open serial
		s, err := serial.OpenPort(c)
		if err != nil {
			return nil, err
		}

		// wrapp to net.PacketConn
		serialConn, err := serialpacket.NewConn(s)
		if err != nil {
			return nil, err
		}

		// kcp conn over serial
		return kcp.NewConn("", block, config.DataShard, config.ParityShard, serialConn)
	} else if config.TCP {
		conn, err := tcpraw.Dial("tcp", config.RemoteAddr)
		if err != nil {
			return nil, errors.Wrap(err, "tcpraw.Dial()")
		}
		return kcp.NewConn(config.RemoteAddr, block, config.DataShard, config.ParityShard, conn)
	}
	return kcp.DialWithOptions(config.RemoteAddr, block, config.DataShard, config.ParityShard)
}
