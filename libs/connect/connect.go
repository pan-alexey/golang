package connect

import (
	"net"
	"time"
)

func CheckTCP(host string, port string) (string, error) {
	var Addr string = host
	timeout := time.Second
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)

	if conn != nil {
		conn.Close()
		return Addr, nil
	}

	return Addr, err
}
