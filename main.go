package main

import (
	"encoding/json"
	"fmt"
	"net"
)

func main() {

	s, err := net.ResolveUDPAddr("udp4", ":9010")
	if err != nil {
		fmt.Println(err)
		return
	}
	conn, err := net.ListenUDP("udp4", s)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	buffer := make([]byte, 1024)

	for {
		_, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			continue
		}

		stats, err := getStats()
		if err != nil {
			fmt.Println(err)
			continue
		}
		//fmt.Printf("stats %+v\n", stats)
		response, err := json.Marshal(&stats)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(string(response))
		_, err = conn.WriteToUDP(response, addr)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}
