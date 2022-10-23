package sender

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func StartSending(network, address string) {
	conn, err := net.Dial(network, address)
	if err != nil {
		log.Fatal("Unable to open connection:", err.Error())
	}
	defer conn.Close()

	for {
		fmt.Print(">> ")

		clientRequest, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			log.Fatal("Unable to read from CLI:", err.Error())
		}
		clientRequest = strings.Replace(clientRequest, "\r", "", -1)
		conn.Write([]byte(clientRequest))

		serverResponse, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatal("Unable to get response:", err.Error())
		}
		log.Print("Server: ", serverResponse)
	}
}
