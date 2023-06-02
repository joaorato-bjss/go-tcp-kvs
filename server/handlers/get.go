package handlers

import (
	"errors"
	"go-tcp-kvs/store"
	"strconv"
)

func Get(command string) (string, error) {

	key, endIndex, err := validateProtocol(command, 0)
	if err != nil {
		return "", err
	}

	var n int
	if endIndex < len(command) {
		n, endIndex, err = validateLengthProtocol(command, endIndex)
		if err != nil {
			return "", err
		}
	}

	// get value with "key" from store
	value := store.DoStoreGet(key)
	if value.Error != nil {
		return "", value.Error
	}

	// create response that satisfies the protocol
	resp := createGetResponseWithProtocol(value.Data, n)

	return resp, nil
}

func validateLengthProtocol(command string, startIndex int) (int, int, error) {
	// Part 1
	lengthByteStr := command[startIndex : startIndex+1]
	lengthByte, err := strconv.Atoi(lengthByteStr)
	if err != nil || lengthByte <= 0 || lengthByte > 9 {
		return 0, 0, errors.New("wrong format. First part of argument has to be a single byte [1-9]")
	}

	// Part 2
	var lengthStr string
	endIndex := startIndex + lengthByte
	if endIndex+1 == len(command) {
		lengthStr = command[startIndex+1 : endIndex+1]
	} else {
		return 0, 0, errors.New("wrong format. Second part of argument have to be " + strconv.Itoa(lengthByte) + " numbers\n")
	}
	length, err := strconv.Atoi(lengthStr)
	if err != nil {
		return 0, 0, errors.New("wrong format. Second part of argument have to be " + strconv.Itoa(lengthByte) + " numbers\n")
	}

	return length, endIndex + 1, nil
}

func createGetResponseWithProtocol(value string, n int) string {
	length := len(value)

	if n > 0 && n < length {
		length = n
	}

	lengthStr := strconv.Itoa(length)

	lengthByte := len(lengthStr)

	return strconv.Itoa(lengthByte) + lengthStr + value[:length]
}
