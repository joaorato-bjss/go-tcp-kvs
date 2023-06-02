package server

import (
	"bufio"
	"errors"
	"fmt"
	"go-tcp-kvs/server/handlers"
	"go-tcp-kvs/server/logger"
	"go-tcp-kvs/store"
	"net"
)

func HandleConnection(c net.Conn, done chan bool) {
	logger.InfoLogger.Println("Handling client request")

	s := bufio.NewScanner(c)
	for s.Scan() {
		txt := s.Text()
		logger.InfoLogger.Printf("Received '%s'\n", txt)

		// extract method. first 3 characters must be "put", "get" or "del" otherwise error
		if len(txt) < 3 {
			logger.ErrorLogger.Println("command does not have valid method")
			_, _ = fmt.Fprintf(c, "err")
		} else {
			method := txt[:3]

			switch method {
			case "put":
				err := handlers.Put(txt[3:])
				if err != nil {
					logger.ErrorLogger.Println(err.Error())
					_, _ = fmt.Fprintf(c, "err")
				} else {
					_, _ = fmt.Fprintf(c, "ack")
				}
			case "get":
				resp, err := handlers.Get(txt[3:])
				if errors.Is(err, store.ErrNotFound) {
					logger.ErrorLogger.Println(err.Error())
					_, _ = fmt.Fprintf(c, "nil")
				} else if err != nil {
					logger.ErrorLogger.Println(err.Error())
					_, _ = fmt.Fprintf(c, "err")
				} else {
					_, _ = fmt.Fprintf(c, "val"+resp)
				}
			case "del":
				err := handlers.Delete(txt[3:])
				if err != nil {
					logger.ErrorLogger.Println(err.Error())
					_, _ = fmt.Fprintf(c, "err")
				} else {
					_, _ = fmt.Fprintf(c, "ack")
				}
			case "bye":
				done <- true
				err := c.Close()
				if err != nil {
					logger.ErrorLogger.Println(err.Error())
				}
			default:
				logger.ErrorLogger.Println("No valid method. Use 'put', 'get' or 'del'")
				_, _ = fmt.Fprintf(c, "err")
			}
		}
	}
}
