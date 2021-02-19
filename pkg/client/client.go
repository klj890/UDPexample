/*
@Time : 2021/2/8 11:34
@Author : wangtao
@File : client
*/
package client

import (
	"UDPexample/pkg/utils"
	"fmt"
	"net"
	"time"
)

func Client(ip string, port string, dataPath string, mtu int) {
	addr := ip + ":" + port
	data, err := utils.LoadFile(dataPath)
	if err != nil {
		panic(err)
	}
	udpAddr, _ := net.ResolveUDPAddr("udp4", addr)
	udpConn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		panic(err)
	}
	fmt.Println("udp dial ok, send to ", addr)
	for {
		handleConnection(udpConn, data, mtu)
		time.Sleep(1 * time.Second)
	}
}

func handleConnection(conn *net.UDPConn, data string, mtu int) {
	var tmpByte = []byte(data)
	if len(tmpByte) > mtu {
		tmpByte = tmpByte[:mtu]
	}
	len, err := conn.Write(tmpByte)
	if err != nil {
		panic(err)
	}
	fmt.Println("client write len:", len)
	//读取数据
	buf := make([]byte, 1500)
	len, _ = conn.Read(buf)
	fmt.Println("client read len:", len)
	fmt.Println("client read data:", string(buf[:]))
}
