package client

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
)

func Run() {
	Conn, _ := net.Dial("udp", "localhost:4040")
	defer Conn.Close()
	file, err := os.Open("D:\\install-list.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	buff := bytes.NewBufferString("install-list.txt\n")
	if _, err := io.Copy(buff, file); err != nil {
		fmt.Println(err)
		return
	}
	if _, err := io.Copy(Conn, buff); err != nil {
		fmt.Println(err)
		return
	}
}
