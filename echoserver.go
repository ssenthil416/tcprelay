package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

func handleRead(conn net.Conn) {
	defer (func() {
		conn.Close()
		fmt.Println("Closed connection")
	})()

	for {

		buf := make([]byte, 256)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			return
		}
		conn.Write(buf)
	}
}

func main() {
	relayHost := flag.String("relayHost", "localhost", "provide relay server Host")
	relayPort := flag.String("relayPort", "8080", "provide relay server port")
	flag.Parse()
	if *relayHost == "" && *relayPort == "" {
		flag.Usage()
		os.Exit(0)
	}
	//host := "10.0.0.221:" + "8081"

	//connect to Replay
	conn, err := net.Dial("tcp", *relayHost+":"+*relayPort)
	if err != nil {
		panic(err)
	}
	//defer conn.Close()

	fmt.Println("Relay Host = ", *relayHost, "Relay Port =", *relayPort)

	buf := make([]byte, 256)
	_, err = conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		return
	}

	fmt.Println("Connect to", string(buf))
	//defer (func() {
	//	read.Close()
	//})()

	client, err := net.Listen("tcp", string(buf))
	if err != nil {
		fmt.Println(" Connection Listen Error =", err.Error())
		os.Exit(1)
	}

	// use this goroutine to wait for and process clients
	for {
		conn, err := client.Accept()
		if err != nil {
			fmt.Println("Couldn't accept :", err)
			break
		}
		go handleRead(conn)
	}

}
