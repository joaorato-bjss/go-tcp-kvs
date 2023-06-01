package handlers

import (
	"errors"
	"go-tcp-kvs/store"
	"strconv"
)

func Put(command string) error {

	key, midIndex, err := validatePutProtocol(command, 0)
	if err != nil {
		return err
	}

	value, endIndex, err2 := validatePutProtocol(command, midIndex)
	if err2 != nil {
		return err2
	}

	if endIndex != len(command) {
		return errors.New("wrong format. Command does not satisfy the protocol")
	}

	// put "value" in store with key "key"
	store.DoStorePut(key, value)

	return nil
}

func validatePutProtocol(command string, startIndex int) (string, int, error) {
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
