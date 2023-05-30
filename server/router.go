package server

import (
	"bufio"
	"fmt"
	"go-tcp-kvs/server/handlers"
	"go-tcp-kvs/server/logger"
	"net"
)

func HandleConnection(c net.Conn) {
	logger.InfoLogger.Println("Handling client request")

	s := bufio.NewScanner(c)
	for s.Scan() {
		txt := s.Text()
		logger.InfoLogger.Printf("Received '%s'\n", txt)

		// extract method. first 3 characters must be "put", "get" or "del" otherwise error
		method := txt[:3]

		switch method {
		case "put":
			handlers.Put(c, txt[3:])
		case "get":
			handlers.Get(c, txt[3:])
		case "del":
			handlers.Delete(c, txt[3:])
		default:
			logger.ErrorLogger.Println("No valid method. Use 'put', 'get' or 'del'")
			_, _ = fmt.Fprintf(c, "err")
		}
	}
}
