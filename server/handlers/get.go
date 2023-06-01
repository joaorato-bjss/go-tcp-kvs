package handlers

import (
	"errors"
	"go-tcp-kvs/store"
	"strconv"
)

func Get(command string) (string, error) {

	key, endIndex, err := validateGetProtocol(command, 0)
	if err != nil {
		return "", err
	}
	if endIndex != len(command) {
		return "", errors.New("wrong format. Command does not satisfy the protocol")
	}

	// get value with "key" from store
	value := store.DoStoreGet(key)
	if value.Error != nil {
		return "", value.Error
	}

	// create response that satisfies the protocol
	resp := createResponseWithProtocol(value.Data)

	return resp, nil
}

func validateGetProtocol(command string, startIndex int) (string, int, error) {
	// Part 1
	lengthByteStr := command[startIndex : startIndex+1]
	lengthByte, err := strconv.Atoi(lengthByteStr)
	if err != nil || lengthByte <= 0 || lengthByte > 9 {
		return "", 0, errors.New("wrong format. First part of argument has to be a single byte [1-9]")
	}

	// Part 2
	lengthStr := command[startIndex+1 : startIndex+1+lengthByte]

	length, err := strconv.Atoi(lengthStr)
	if err != nil {
		return "", 0, errors.New("wrong format. Second part of argument have to be " + strconv.Itoa(lengthByte) + " numbers\n")
	}

	endIndex := startIndex + 1 + lengthByte + length
	if endIndex > len(command) {
		return "", 0, errors.New("wrong format. Third part of argument has to be " + strconv.Itoa(length) + " character\n")
	}

	return command[startIndex+1+lengthByte : endIndex], endIndex, nil
}

func createResponseWithProtocol(value string) string {
	length := len(value)

	lengthStr := strconv.Itoa(length)

	lengthByte := len(lengthStr)

	return strconv.Itoa(lengthByte) + lengthStr + value
}
