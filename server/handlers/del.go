package handlers

import (
	"fmt"
	"go-tcp-kvs/server/logger"
	"go-tcp-kvs/store"
	"net"
)

func Delete(c net.Conn, command string) {
	// delete uses the same format as get
	key, endIndex, err := validateGetProtocol(command, 0)
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
		_, _ = fmt.Fprintf(c, "err")
		return
	}
	if endIndex != len(command) {
		logger.ErrorLogger.Println("Wrong format. Command does not satisfy protocol")
		_, _ = fmt.Fprintf(c, "err")
		return
	}

	// delete value with "key" from store
	resp := store.DoStoreDelete(key)
	if resp.Error != nil {
		logger.ErrorLogger.Println(resp.Error.Error())
		_, _ = fmt.Fprintf(c, "err")
		return
	}

	_, _ = fmt.Fprintf(c, "ack")
}
