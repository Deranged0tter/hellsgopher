package hellsgopher

import (
	"net"
)

// get the internal ip of the computer
func GetInternalIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return ""
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
}

// return a slice of all interfaces on the computer
func GetAllInterfaces() []string {
	ips := []string{}

	addresses, err := net.InterfaceAddrs()
	if err != nil {
		return nil
	}

	for _, address := range addresses {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ips = append(ips, ipnet.IP.String())
			}
		}
	}

	return ips
}
