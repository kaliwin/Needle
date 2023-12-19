package main

import (
	"github.com/kaliwin/Needle/network/dns"
	"log"
)

func main() {
	err := dns.ServeDNS(":53", "com", "192.168.3.104")
	if err != nil {
		log.Println(err)
	}
}
