package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
)

var (
	host  string
	op    string
	key   string
	value string
)

func main() {
	flag.StringVar(&host, "h", "localhost:8888", "ip:port格式")
	flag.StringVar(&op, "c", "get", "操作命令:set|get|del|stat")
	flag.StringVar(&key, "k", "", "key")
	flag.StringVar(&value, "v", "", "value")
	flag.Parse()

	// 命令校验
	if op != "set" && op != "get" && op != "del" && op != "stat" {
		log.Println("命令不正确")
		return
	}

	conn, err := net.Dial("tcp", host)
	if err != nil {
		panic(err)
	}
	// 关闭连接
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	cmd := fmt.Sprintf("%s %s %s ", op, key, value)
	_, err = conn.Write([]byte(cmd))
	if err != nil {
		panic(err)
	}
	// TODO 后续做结果读取操作
	result := make([]byte, 1024)
	for {
		n, err := conn.Read(result)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		fmt.Print(string(result[:n]))
	}

}
