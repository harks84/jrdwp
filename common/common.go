package common

import (
	"log"
	"net"
	"path"
	"strconv"
	"strings"
	"time"
)

const (
	//HeaderToken security token
	HeaderToken = "jrdwptoken"
	//HeaderPort jdwp port
	HeaderPort = "jdwpport"
	//PublicKeyDir default client key path
	PublicKeyDir = "."
	//PublicKeyFile filename
	PublicKeyFile = ".jrdwp_key"
	//PrivateKeyDir default client key path
	PrivateKeyDir = "."
	//PrivateKeyFile filename
	PrivateKeyFile = ".jrdwp_privatekey"
)

var DeadlineDuration = time.Second * 60 * 30

//PublicKeyPath concat file path of client key
func PublicKeyPath() string {
	return path.Join(PublicKeyDir, PublicKeyFile)
}

//PrivateKeyPath concat file path of private key
func PrivateKeyPath() string {
	return path.Join(PrivateKeyDir, PrivateKeyFile)
}

//SplitToInt split comma delimited string to int array
func SplitToInt(commaDelimitedString string) []int {
	ports := []int{}

	for _, token := range strings.Split(commaDelimitedString, ",") {
		port, err := strconv.Atoi(string(strings.TrimSpace(token)))
		if err != nil {
			log.Fatalln("bad integer", token, err.Error)
		} else {
			ports = append(ports, port)
		}
	}

	return ports
}

//InitTCPConn initialize tcp connection to keep alive
func InitTCPConn(conn *net.TCPConn) {
	conn.SetKeepAlive(true)
	conn.SetNoDelay(true)
	conn.SetLinger(3)
	conn.SetDeadline(time.Now().Add(DeadlineDuration))
}
