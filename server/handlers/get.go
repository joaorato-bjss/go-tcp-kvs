package handlers

import (
	"errors"
	"fmt"
	"go-tcp-kvs/server/logger"
	"go-tcp-kvs/store"
	"net"
	"strconv"
)

func Get(c net.Conn, command string) {

	key, endIndex, err := validateGetProtocol(command, 0)
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
		_, _ = fmt.Fprintf(c, "err")
		return
	}
	if endIndex != len(command) {
		logger.ErrorLogger.Println("Wrong format. Command does not satisfy the protocol")
		_, _ = fmt.Fprintf(c, "err")
		return
	}

	// get value with "key" from store
	value := store.DoStoreGet(key)
	if value.Error != nil {
		logger.ErrorLogger.Println(value.Error.Error())
		_, _ = fmt.Fprintf(c, "nil")
		return
	}

	// create response that satisfies the protocol
	resp := createResponseWithProtocol(value.Data)

	_, _ = fmt.Fprintf(c, "val"+resp)
}

func validateGetProtocol(command string, startIndex int) (string, int, error) {
	lengthByteStr := command[startIndex : startIndex+1]
	lengthByte, err := strconv.Atoi(lengthByteStr)
	if err != nil || lengthByte <= 0 || lengthByte > 9 {
		return "", 0, errors.New("wrong format. First part of argument has to be a single byte [1-9]")
	}

	lengthStr := command[startIndex+1 : startIndex+1+lengthByte]

	length, err := strconv.Atoi(lengthStr)
	if err != nil {
		return "", 0, errors.New("wrong format. Second part of argument have to be " + strconv.Itoa(lengthByte) + " numbers\n")
	}

	endIndex := startIndex + 1 + lengthByte + length

	return command[startIndex+1+lengthByte : endIndex], endIndex, nil
}

func createResponseWithProtocol(value string) string {
	length := len(value)

	lengthStr := strconv.Itoa(length)

	lengthByte := len(lengthStr)

	return strconv.Itoa(lengthByte) + lengthStr + value
}
