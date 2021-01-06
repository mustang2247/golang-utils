package main

import (
	"log"
	"net"
)

func main() {
	log.Println(GetLocalIP())
}

// 获取本地网址，在生产服务器可以使用，主要用于微服务注册
func GetLocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Println(err)
		return "", err
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	panic("unable to determine locla ip")
}
