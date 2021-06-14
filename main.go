package main

import (
	"fmt"
	"net"
)

/**
*@Author lyer
*@Date 6/14/21 16:01
*@Describe
**/
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	addrs, err := net.InterfaceAddrs()
	checkError(err)

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP)
			}
		}
	}
}
