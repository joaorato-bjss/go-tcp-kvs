package handlers

import (
	"errors"
	"go-tcp-kvs/store"
)

func Delete(command string) error {
	// delete uses the same format as get
	key, endIndex, err := validateGetProtocol(command, 0)
	if err != nil {
		return err
	}
	if endIndex != len(command) {
		return errors.New("wrong format. Command does not satisfy protocol")
	}

	// delete value with "key" from store
	resp := store.DoStoreDelete(key)
	if resp.Error != nil {
		return resp.Error
	}

	return nil
}
