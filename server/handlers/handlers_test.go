package handlers

import (
	"errors"
	"go-tcp-kvs/store"
	"testing"
)

// tests the Put handler without connection (requires store to be initialised)
func TestPut2(t *testing.T) {
	store.InitStore()

	t.Run("TestValidSmallKeySmallValue", func(t *testing.T) {
		err := Put("11k11v")
		if err != nil {
			t.Error("expected no error, got: ", err.Error())
		}
	})
	t.Run("TestValidBigKeyBigValue", func(t *testing.T) {
		err := Put("210firstkey0115value")
		if err != nil {
			t.Error("expected no error, got: ", err.Error())
		}
	})
	t.Run("TestValidOnlyNumbers", func(t *testing.T) {
		err := Put("111111")
		if err != nil {
			t.Error("expected no error, got: ", err.Error())
		}
	})
	t.Run("TestInvalidZeroKey", func(t *testing.T) {
		err := Put("011v")
		if err == nil {
			t.Error("expected error, got none")
		}
	})
	t.Run("TestInvalidZeroValue", func(t *testing.T) {
		err := Put("11k01v")
		if err == nil {
			t.Error("expected error, got none")
		}
	})
	t.Run("TestInvalidSmallKeySmallValue1", func(t *testing.T) {
		err := Put("11k110v")
		if err == nil {
			t.Error("expected error, got none")
		}
	})
	t.Run("TestInvalidSmallKeySmallValue2", func(t *testing.T) {
		err := Put("21k11v")
		if err == nil {
			t.Error("expected error, got none")
		}
	})
	t.Run("TestInvalidBigKeyBigValue1", func(t *testing.T) {
		err := Put("210firstkey0115vvalue")
		if err == nil {
			t.Error("expected error, got none")
		}
	})
	t.Run("TestInvalidBigKeyBigValue2", func(t *testing.T) {
		err := Put("210firstkey0115val")
		if err == nil {
			t.Error("expected error, got none")
		}
	})
	t.Run("TestInvalidOnlyNumbers", func(t *testing.T) {
		err := Put("1111111")
		if err == nil {
			t.Error("expected error, got none")
		}
	})
	t.Run("TestRewriteKey", func(t *testing.T) {
		err := Put("210firstkey0116value2")
		if err != nil {
			t.Error("expected no error, got: ", err.Error())
		}
	})
}

// tests the Get handler without connection (requires store to be initialised)
// we assume these tests are run immediately after the Put ones.
func TestGet2(t *testing.T) {
	t.Run("TestBadGetFormat", func(t *testing.T) {
		_, err := Get("11k11v")
		if err == nil {
			t.Error("expected error, got none")
		}
	})
	t.Run("TestValidSmallKeySmallValue", func(t *testing.T) {
		resp, err := Get("11k")
		if err != nil || resp != "11v" {
			t.Error("expected '11v' and no error, got: ", resp, err.Error())
		}
	})
	t.Run("TestValidBigKeyBigValue", func(t *testing.T) {
		resp, err := Get("210firstkey01")
		if err != nil || resp != "16value2" {
			t.Error("expected '16value2' and no error, got: ", resp, err.Error())
		}
	})
	t.Run("TestValidOnlyNumbers", func(t *testing.T) {
		resp, err := Get("111")
		if err != nil || resp != "111" {
			t.Error("expected '111' and no error, got: ", resp, err.Error())
		}
	})
	t.Run("TestInvalidZeroKey", func(t *testing.T) {
		_, err := Get("0")
		if err == nil {
			t.Error("expected error, got none")
		}
	})
	t.Run("TestInvalidSmallKey1", func(t *testing.T) {
		_, err := Get("110k")
		if err == nil {
			t.Error("expected error, got none")
		}
	})
	t.Run("TestInvalidSmallKey2", func(t *testing.T) {
		_, err := Get("21k")
		if err == nil {
			t.Error("expected error, got none")
		}
	})
	t.Run("TestInvalidBigKeyBigValue", func(t *testing.T) {
		_, err := Get("210firstkey0")
		if err == nil {
			t.Error("expected error, got none")
		}
	})
	t.Run("TestInvalidOnlyNumbers", func(t *testing.T) {
		_, err := Get("1111")
		if err == nil {
			t.Error("expected error, got none")
		}
	})
	t.Run("TestKeyNotFound", func(t *testing.T) {
		resp, err := Get("11a")
		if !errors.Is(err, store.ErrNotFound) || resp != "" {
			t.Error("expected 'not found', got: ", err.Error())
		}
	})
}

// tests the Delete handler without connection (requires store to be initialised)
// we assume these tests are run after the Put ones.
func TestDelete2(t *testing.T) {
	t.Run("TestBadDeleteFormat", func(t *testing.T) {
		err := Delete("11k11v")
		if err == nil {
			t.Error("expected error, got none")
		}
	})
	t.Run("TestValidSmallKeySmallValue", func(t *testing.T) {
		err := Delete("11k")
		if err != nil {
			t.Error("expected no error, got: ", err.Error())
		}
	})
	t.Run("TestValidBigKeyBigValue", func(t *testing.T) {
		err := Delete("210firstkey01")
		if err != nil {
			t.Error("expected no error, got: ", err.Error())
		}
	})
	t.Run("TestValidOnlyNumbers", func(t *testing.T) {
		err := Delete("111")
		if err != nil {
			t.Error("expected no error, got: ", err.Error())
		}
	})
	t.Run("TestInvalidZeroKey", func(t *testing.T) {
		err := Delete("0")
		if err == nil {
			t.Error("expected error, got none")
		}
	})
	t.Run("TestInvalidSmallKey1", func(t *testing.T) {
		err := Delete("110k")
		if err == nil {
			t.Error("expected error, got none")
		}
	})
	t.Run("TestInvalidSmallKey2", func(t *testing.T) {
		err := Delete("21k")
		if err == nil {
			t.Error("expected error, got none")
		}
	})
	t.Run("TestInvalidBigKeyBigValue", func(t *testing.T) {
		err := Delete("210firstkey0")
		if err == nil {
			t.Error("expected error, got none")
		}
	})
	t.Run("TestInvalidOnlyNumbers", func(t *testing.T) {
		err := Delete("1111")
		if err == nil {
			t.Error("expected error, got none")
		}
	})
	t.Run("TestKeyNotFound", func(t *testing.T) {
		err := Delete("11a")
		if !errors.Is(err, store.ErrNotFound) {
			t.Error("expected 'not found', got: ", err.Error())
		}
	})
}
