package server

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
)

func Run() {
	listener, err := net.ListenPacket("udp", ":4040")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	for {
		buff := make([]byte, 1024)
		n, addr, err := listener.ReadFrom(buff)
		if err != nil {
			fmt.Println("connection error")
			continue
		}
		go handleConnection(listener, addr, buff[:n])
	}
}

func handleConnection(listener net.PacketConn, addr net.Addr, buff []byte) error {
	reader := bufio.NewReader(bytes.NewReader(buff))
	filename, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println(err)
		return err
	}
	file, err := os.Create("ftp\\public\\" + string(filename))
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()
	if _, err := io.Copy(file, reader); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
