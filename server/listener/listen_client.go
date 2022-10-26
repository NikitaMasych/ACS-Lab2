package listener

import (
	"bufio"
	"log"
	"net"
	"server/client"
	"strings"
)

func StartListen(network, address string) {
	server, err := net.Listen(network, address)
	if err != nil {
		log.Fatal(err)
	}
	defer server.Close()

	conn, err := server.Accept()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for {
		clientRequest, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		clientRequest = strings.TrimSuffix(clientRequest, "\n")
		result := client.ProcessRequest(clientRequest)

		conn.Write([]byte(result + "\r"))
	}
}
