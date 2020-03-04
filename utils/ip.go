package utils

import (
"net"
"strings"
)

// GetOutboundIP 获取本地对外IP
func GetOutboundIP(addr string) (ip string, err error) {
	conn, err := net.Dial("udp", addr)
	if err != nil {
		return
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	//fmt.Println(localAddr.String())
	ip = strings.Split(localAddr.IP.String(), ":")[0]
	return
}