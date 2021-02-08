/*
@Time : 2021/2/8 11:20
@Author : wangtao
@File : server
*/
package server

import (
	"fmt"
	"net"
	"time"
)

const netWork = "udp4"

func Server(ip string, port string) {
	addr := ip + ":" + port
	udpAddr, _ := net.ResolveUDPAddr(netWork, addr)
	udpConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println(err)
	}
	defer udpConn.Close()
	fmt.Println("udp listening on ", addr)

	//udp不需要Accept
	for {
		handleConnection(udpConn)
	}
}

func handleConnection(conn *net.UDPConn) {
	data := make([]byte, 1500)
	n, remoteAddr, err := conn.ReadFromUDP(data)
	if err != nil {
		fmt.Println("failed to read UDP msg because of ", err.Error())
		return
	}
	daytime := time.Now()
	fmt.Println("server read len:", n, remoteAddr)
	b := make([]byte, 1500)
	b = []byte(daytime.Format(time.RFC3339))
	conn.WriteToUDP(b, remoteAddr)
}
