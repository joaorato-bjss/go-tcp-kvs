package handlers

import (
	"go-tcp-kvs/server/logger"
	"net"
)

func Bye(c net.Conn, done chan bool) {
	err := c.Close()
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
	}

	done <- true
}
