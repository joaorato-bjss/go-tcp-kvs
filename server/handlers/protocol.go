package handlers

import (
	"errors"
	"strconv"
)

func validateProtocol(command string, startIndex int) (string, int, error) {
	// Part 1
	lengthByteStr := command[startIndex : startIndex+1]
	lengthByte, err := strconv.Atoi(lengthByteStr)
	if err != nil || lengthByte <= 0 || lengthByte > 9 {
		return "", 0, errors.New("wrong format. First part of argument has to be a single byte [1-9]")
	}

	// Part 2
	var lengthStr string

	if startIndex+1+lengthByte <= len(command) {
		lengthStr = command[startIndex+1 : startIndex+1+lengthByte]
	} else {
		return "", 0, errors.New("wrong format. Second part of argument have to be " + strconv.Itoa(lengthByte) + " numbers\n")
	}
	length, err := strconv.Atoi(lengthStr)
	if err != nil {
		return "", 0, errors.New("wrong format. Second part of argument have to be " + strconv.Itoa(lengthByte) + " numbers\n")
	}

	endIndex := startIndex + 1 + lengthByte + length
	if endIndex > len(command) {
		return "", 0, errors.New("wrong format. Third part of argument has to be " + strconv.Itoa(length) + " characters\n")
	}

	return command[startIndex+1+lengthByte : endIndex], endIndex, nil
}
