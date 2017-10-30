package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
)

func relayHandler(conn net.Conn, listenChan chan int) {
	//defer conn.Close()

	port := <-listenChan
	echoSerAdd := "localhost:" + strconv.Itoa(port)

	//fmt.Println("echoSerAdd =", echoSerAdd)

	// Sending
	_, err := conn.Write([]byte(echoSerAdd))
	if err != nil {
		fmt.Println("Cannot Write to connection")
		return
	}

	//fmt.Println("Bytes Written =", wc)
}

func main() {
	port := flag.String("port", "8080", "provide 8080 to relay start in that port")
	flag.Parse()
	if *port != "8080" {
		flag.Usage()
		os.Exit(0)
	}

	listen, err := net.Listen("tcp", ":"+*port)
	if err != nil {
		fmt.Println(" Connection Listen Error =", err.Error())
		os.Exit(1)
	}
	//defer listen.Close()

	fmt.Println("Relay is running on : ", *port)

	numService := 1
	listenChan := make(chan int, numService)
	listenChan <- 8081

	//listen for ever
	for {
		//Wait for conn
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(" Connection Accept Error =", err.Error())
			break
		}

		fmt.Println("Conneciton Accepted")
		go relayHandler(conn, listenChan)
	}

	fmt.Println("Replay is Stopped!!!!!")
}
