package handlers

import (
	"errors"
	"go-tcp-kvs/store"
)

func Put(command string) error {

	key, midIndex, err := validateProtocol(command, 0)
	if err != nil {
		return err
	}

	value, endIndex, err2 := validateProtocol(command, midIndex)
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
