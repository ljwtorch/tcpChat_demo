package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.199.120:8888")
	if err != nil {
		fmt.Println("client dial err =", err)
		return
	}
	//fmt.Println("conn success =", conn)

	//1、客户端可以发送单行数据，然后就退出
	reader := bufio.NewReader(os.Stdin)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("readString err =", err)
		}
		//如果用户输入exit就退出
		line = strings.Trim(line, " \r\n")
		if line == "exit" {
			fmt.Println("客户端已退出！")
			break
		}
		//再将line发送给服务器
		_, err = conn.Write([]byte(line))
		if err != nil {
			fmt.Println("conn.Write err =", err)
		}
		//fmt.Printf("客户端发送了 %d 字节的数据，并退出", n)
	}
}
