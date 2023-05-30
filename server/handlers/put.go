package handlers

import (
	"fmt"
	"go-tcp-kvs/server/logger"
	"net"
	"strconv"
)

func Put(c net.Conn, command string) {

	key, midIndex := validatePutProtocol(c, command, 0)

	value, endIndex := validatePutProtocol(c, command, midIndex)
	if endIndex != len(command)-1 {
		logger.ErrorLogger.Println("Wrong format. Command does not satisfy the protocol")
		_, _ = fmt.Fprintf(c, "err")
	}

	// put "value" in store with key "key"

	_, _ = fmt.Fprintf(c, "ack")
}

func validatePutProtocol(c net.Conn, command string, startIndex int) (string, int) {
	lengthByte := command[startIndex]
	if lengthByte <= 0 || lengthByte > 9 {
		logger.ErrorLogger.Println("Wrong format. First part of argument has to be a single byte [1-9]")
		_, _ = fmt.Fprintf(c, "err")
	}

	lengthStr := command[startIndex+1 : startIndex+1+int(lengthByte)]

	length, err := strconv.Atoi(lengthStr)
	if err != nil {
		logger.ErrorLogger.Printf("Wrong format. Second part of argument have to be %v numbers\n", lengthByte)
		_, _ = fmt.Fprintf(c, "err")
	}

	endIndex := startIndex + 1 + int(lengthByte) + length

	return command[startIndex+1+int(lengthByte) : endIndex], endIndex
}
