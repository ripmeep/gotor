/*
	Author:    ripmeep
	Instagram: @rip.meep
	GitHub:    https://github.com/ripmeep/
*/

package tor

import (
	"net"
	"errors"
	"strconv"
	"encoding/binary"
)

type TorConnection struct {
	Host string
	SocksPort int
	ControlPort int
}

func (t TorConnection) Connect(host string, port int) (net.Conn, error) {
	con, err := net.Dial("tcp", t.Host + ":" + strconv.Itoa(t.SocksPort))

	if err != nil {
		return nil, err
	}

	sport := uint16(port) // Convert integer to uint16 (short) for network usage
	init := "\x05\x01\x00" // SOCKS5, one auth
	reply := make([]byte, 64)
	var tor_ok byte = 0

	con.Write([]byte(init))
	con.Read(reply)

	if reply[1] != tor_ok {
		return nil, errors.New("Tor authentication message was not accepted")
	}

	reply = make([]byte, 64)
	ptns := make([]byte, 2)

	binary.BigEndian.PutUint16(ptns, sport)

	host_req := "\x05\x01\x00\x03" // Connect
	host_req += string(len(host))
	host_req += host // Host
	host_req += string(ptns) // Port

	con.Write([]byte(host_req))
	con.Read(reply)

	if reply[1] != tor_ok {
		return nil, errors.New("Tor connection message was not accepted")
	}

	return con, nil
}

func (t TorConnection) Refresh(password string) (bool, error) {
	con, err := net.Dial("tcp", t.Host + ":" + strconv.Itoa(t.ControlPort))

	if err != nil {
		return false, err
	}

	reply := make([]byte, 1024)
	var tor_ok string = "250 OK"
	auth := []byte("authenticate \"" + password + "\"\r\n")

	con.Write(auth)
	con.Read(reply)

	if string(reply[:6]) != tor_ok {
		return false, errors.New("Tor control authentication failed: " + string(reply))
	}

	reply = make([]byte, 1024)
	refresh := []byte("signal newnym\r\n")

	con.Write(refresh)
	con.Read(reply)

	if string(reply[:6]) != tor_ok {
		return false, errors.New("Tor control authentication failed: " + string(reply))
	}

	return true, nil
}
