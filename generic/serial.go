package generic

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/tarm/serial"
)

const (
	// serial@name@baud
	// eg:
	// 	serial@/dev/ttys002@9600
	// 	serial@COM4@9600
	SERIAL_PROTO = "serial@"
)

func ParseSerialParams(addr string) (*serial.Config, error) {
	// serial://NAME
	strs := strings.Split(addr, "@")
	if len(strs) < 3 {
		return nil, errors.Errorf("serial format error:%v", addr)
	}

	port := strs[1]
	// parse baud
	baud, err := strconv.Atoi(strs[2])
	if err != nil {
		return nil, err
	}

	config := new(serial.Config)
	config.Name = port
	config.Baud = baud

	return config, nil
}
