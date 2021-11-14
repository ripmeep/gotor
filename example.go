package main

import (
	"fmt"
	"log"
	"time"
	"strings"
	"github.com/ripmeep/gotor"
)

func main() {
	host := "ident.me"
	port := 80

	tor := tor.TorConnection { "127.0.0.1", 9050, 9051 } // Tor host, SOCKS5 port, Control Port
	tor_con, err := tor.Connect(host, port)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Connected to %s:%d through Tor service (%s:%d)\n", host, port, tor.Host, tor.SocksPort)

	req := []byte("GET / HTTP/1.1\r\nHost: www.ident.me\r\n\r\n") // IP website request
	res := make([]byte, 4096)

	tor_con.Write(req)
	tor_con.Read(res)

	res_split := strings.Split(string(res), "\r\n\r\n")
	ip := res_split[len(res_split) - 1]

	fmt.Println("Tor IP: " + ip)

	ok, err := tor.Refresh("torpassword") // REPLACE "torpassword" WITH YOUR CONTROL PASSWORD

	if !ok {
		fmt.Println(err)
		log.Fatal("Tor refresh failed")
	}

	fmt.Println("\nTor refreshed successfully! New IP assigned\n")
	time.Sleep(2 * time.Second)

	/* Make second request after refresh (IP should change) */

	tor_con, err = tor.Connect(host, port)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Connected to %s:%d through Tor service (%s:%d)\n", host, port, tor.Host, tor.SocksPort)

	req = []byte("GET / HTTP/1.1\r\nHost: www.ident.me\r\n\r\n") // IP website request
	res = make([]byte, 4096)

	tor_con.Write(req)
	tor_con.Read(res)

	res_split = strings.Split(string(res), "\r\n\r\n")
	ip = res_split[len(res_split) - 1]

	fmt.Println("Tor IP: " + ip)
}
