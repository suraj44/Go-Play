package main

import (

	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

var addr = flag.String("addr", "","")

var port = flag.Int("port", 8000,"")

func main() {
	fmt.Println("Launching the server...")

	src := *addr + ":" + strconv.Itoa(*port)

	listener, _ := net.Listen("tcp", src)
	fmt.Printf("Listening on %s\n", src)

	defer listener.Close()


	for  {
		conn ,err := listener.Accept()

		if err != nil {
			fmt.Printf("Connection error: %s\n", err)
		}
		
		go handleConnection(conn)

	}
}

func handleConnection(conn net.Conn) {
	remoteAddr := conn.RemoteAddr().String()
	fmt.Println("Client connected from " + remoteAddr)
	scanner := bufio.NewScanner(conn)

	for {
		ok := scanner.Scan()

		handleMessage(scanner.Text(), conn)

		if !ok {
			break
		}

	}
	fmt.Println("Client at "+ remoteAddr + "disconnected.")

}

func handleMessage(message string, conn net.Conn) {
	fmt.Println("> "+message)

	if message[0] == '/' {

		switch {
		case message == "/time":

			resp := "It is " + time.Now().String() + "\n"
			fmt.Print("< "+resp)
			conn.Write([]byte(resp))

		case message == "/joke":

			resp := "A machine learning algorithm walks into a bar.\n The bartender asks, 'What will you have?'\nThe algorithm says, 'What's everyone else having?'\n"
			fmt.Print("< "+resp)
			conn.Write([]byte(resp))

		case message == "/quit":
			fmt.Println("GOODBYE!")
			conn.Write([]byte("I'm shutting down now, see you later!\n"))
			fmt.Println("< " + "%quit%")
			conn.Write([]byte("%quit%"))
			os.Exit(0)

		default:
			resp := "Unrecognized commmand\n"
			fmt.Print("< "+resp)
			conn.Write([]byte(resp))
		}
	}
}
	

