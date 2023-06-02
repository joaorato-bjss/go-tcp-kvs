package handlers

import (
	"net"
	"testing"
)

// tests the Put handler via connection (requires server to be running)
func TestPut(t *testing.T) {
	c, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		t.Error("could not connect to server: ", err)
	}
	t.Run("TestSmallString", func(t *testing.T) {
		if _, err2 := c.Write([]byte("p\n")); err2 != nil {
			t.Error("error writing to server: ", err2)
		}
		b := make([]byte, 3)
		if _, err3 := c.Read(b); err3 != nil {
			t.Error("error reading from server: ", err3)
		}
		if string(b) != "err" {
			t.Error("expected 'err', got: ", string(b))
		}
	})
	t.Run("TestBadMethod", func(t *testing.T) {
		if _, err2 := c.Write([]byte("pu11k11v\n")); err2 != nil {
			t.Error("error writing to server: ", err2)
		}
		b := make([]byte, 3)
		if _, err3 := c.Read(b); err3 != nil {
			t.Error("error reading from server: ", err3)
		}
		if string(b) != "err" {
			t.Error("expected 'err', got: ", string(b))
		}
	})
	t.Run("TestValidSmallKeySmallValue", func(t *testing.T) {
		if _, err2 := c.Write([]byte("put11k11v\n")); err2 != nil {
			t.Error("error writing to server: ", err2)
		}
		b := make([]byte, 3)
		if _, err3 := c.Read(b); err3 != nil {
			t.Error("error reading from server: ", err3)
		}
		if string(b) != "ack" {
			t.Error("expected 'ack', got: ", string(b))
		}
	})
	t.Run("TestValidBigKeyBigValue", func(t *testing.T) {
		if _, err2 := c.Write([]byte("put210firstkey0115value\n")); err2 != nil {
			t.Error("error writing to server: ", err2)
		}
		b := make([]byte, 3)
		if _, err3 := c.Read(b); err3 != nil {
			t.Error("error reading from server: ", err3)
		}
		if string(b) != "ack" {
			t.Error("expected 'ack', got: ", string(b))
		}
	})
	t.Run("TestValidOnlyNumbers", func(t *testing.T) {
		if _, err2 := c.Write([]byte("put111111\n")); err2 != nil {
			t.Error("error writing to server: ", err2)
		}
		b := make([]byte, 3)
		if _, err3 := c.Read(b); err3 != nil {
			t.Error("error reading from server: ", err3)
		}
		if string(b) != "ack" {
			t.Error("expected 'ack', got: ", string(b))
		}
	})
	t.Run("TestInvalidZeroKey", func(t *testing.T) {
		if _, err2 := c.Write([]byte("put011v\n")); err2 != nil {
			t.Error("error writing to server: ", err2)
		}
		b := make([]byte, 3)
		if _, err3 := c.Read(b); err3 != nil {
			t.Error("error reading from server: ", err3)
		}
		if string(b) != "err" {
			t.Error("expected 'err', got: ", string(b))
		}
	})
	t.Run("TestInvalidZeroValue", func(t *testing.T) {
		if _, err2 := c.Write([]byte("put11k01v\n")); err2 != nil {
			t.Error("error writing to server: ", err2)
		}
		b := make([]byte, 3)
		if _, err3 := c.Read(b); err3 != nil {
			t.Error("error reading from server: ", err3)
		}
		if string(b) != "err" {
			t.Error("expected 'err', got: ", string(b))
		}
	})
	t.Run("TestInvalidSmallKeySmallValue1", func(t *testing.T) {
		if _, err2 := c.Write([]byte("put11k110v\n")); err2 != nil {
			t.Error("error writing to server: ", err2)
		}
		b := make([]byte, 3)
		if _, err3 := c.Read(b); err3 != nil {
			t.Error("error reading from server: ", err3)
		}
		if string(b) != "err" {
			t.Error("expected 'err', got: ", string(b))
		}
	})
	t.Run("TestInvalidSmallKeySmallValue2", func(t *testing.T) {
		if _, err2 := c.Write([]byte("put21k11v\n")); err2 != nil {
			t.Error("error writing to server: ", err2)
		}
		b := make([]byte, 3)
		if _, err3 := c.Read(b); err3 != nil {
			t.Error("error reading from server: ", err3)
		}
		if string(b) != "err" {
			t.Error("expected 'err', got: ", string(b))
		}
	})
	t.Run("TestInvalidSmallKeySmallValue3", func(t *testing.T) {
		if _, err2 := c.Write([]byte("put1\n")); err2 != nil {
			t.Error("error writing to server: ", err2)
		}
		b := make([]byte, 3)
		if _, err3 := c.Read(b); err3 != nil {
			t.Error("error reading from server: ", err3)
		}
		if string(b) != "err" {
			t.Error("expected 'err', got: ", string(b))
		}
	})
	t.Run("TestInvalidBigKeyBigValue1", func(t *testing.T) {
		if _, err2 := c.Write([]byte("put210firstkey0115vvalue\n")); err2 != nil {
			t.Error("error writing to server: ", err2)
		}
		b := make([]byte, 3)
		if _, err3 := c.Read(b); err3 != nil {
			t.Error("error reading from server: ", err3)
		}
		if string(b) != "err" {
			t.Error("expected 'err', got: ", string(b))
		}
	})
	t.Run("TestInvalidBigKeyBigValue2", func(t *testing.T) {
		if _, err2 := c.Write([]byte("put210firstkey\n")); err2 != nil {
			t.Error("error writing to server: ", err2)
		}
		b := make([]byte, 3)
		if _, err3 := c.Read(b); err3 != nil {
			t.Error("error reading from server: ", err3)
		}
		if string(b) != "err" {
			t.Error("expected 'err', got: ", string(b))
		}
	})
	t.Run("TestInvalidBigKeyBigValue3", func(t *testing.T) {
		if _, err2 := c.Write([]byte("put210firstkey0115val\n")); err2 != nil {
			t.Error("error writing to server: ", err2)
		}
		b := make([]byte, 3)
		if _, err3 := c.Read(b); err3 != nil {
			t.Error("error reading from server: ", err3)
		}
		if string(b) != "err" {
			t.Error("expected 'err', got: ", string(b))
		}
	})
	t.Run("TestInvalidBigKeyBigValue4", func(t *testing.T) {
		if _, err2 := c.Write([]byte("put210firstkey01511111\n")); err2 != nil {
			t.Error("error writing to server: ", err2)
		}
		b := make([]byte, 3)
		if _, err3 := c.Read(b); err3 != nil {
			t.Error("error reading from server: ", err3)
		}
		if string(b) != "err" {
			t.Error("expected 'err', got: ", string(b))
		}
	})
	t.Run("TestInvalidOnlyNumbers", func(t *testing.T) {
		if _, err2 := c.Write([]byte("put1111111\n")); err2 != nil {
			t.Error("error writing to server: ", err2)
		}
		b := make([]byte, 3)
		if _, err3 := c.Read(b); err3 != nil {
			t.Error("error reading from server: ", err3)
		}
		if string(b) != "err" {
			t.Error("expected 'err', got: ", string(b))
		}
	})
	t.Run("TestRewriteKey", func(t *testing.T) {
		if _, err2 := c.Write([]byte("put210firstkey0116value2\n")); err2 != nil {
			t.Error("error writing to server: ", err2)
		}
		b := make([]byte, 3)
		if _, err3 := c.Read(b); err3 != nil {
			t.Error("error reading from server: ", err3)
		}
		if string(b) != "ack" {
			t.Error("expected 'ack', got: ", string(b))
		}
	})
	err2 := c.Close()
	if err2 != nil {
		t.Error("error closing connection: ", err2)
	}
}

// tests the Get handler via connection (requires server to be running)
// we assume these tests are run immediately after the Put ones.
func TestGet(t *testing.T) {
	c, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		t.Error("could not connect to server: ", err)
	}
	t.Run("TestBadMethod", func(t *testing.T) {
		if _, err2 := c.Write([]byte("ge11k\n")); err2 != nil {
			t.Error("error writing to server: ", err2)
		}
		b := make([]byte, 3)
		if _, err3 := c.Read(b); err3 != nil {
			t.Error("error reading from server: ", err3)
		}
		if string(b) != "err" {
			t.Error("expected 'err', got: ", string(b))
		}
	})
	t.Run("TestBadGetFormat", func(t *testing.T) {
		if _, err2 := c.Write([]byte("get11k11v\n")); err2 != nil {
			t.Error("error writing to server: ", err2)
		}
		b := make([]byte, 3)
		if _, err3 := c.Read(b); err3 != nil {
			t.Error("error reading from server: ", err3)
		}
		if string(b) != "err" {
			t.Error("expected 'err', got: ", string(b))
		}
	})
	t.Run("TestValidSmallKeySmallValue", func(t *testing.T) {
		if _, err2 := c.Write([]byte("get11k\n")); err2 != nil {
			t.Error("error writing to server: ", err2)
		}
		b := make([]byte, 6)
		if _, err3 := c.Read(b); err3 != nil {
			t.Error("error reading from server: ", err3)
		}
		if string(b) != "val11v" {
			t.Error("expected 'val11v', got: ", string(b))
		}
	})
	t.Run("TestValidSmallKeySmallValueVariable1", func(t *testing.T) {
		if _, err2 := c.Write([]byte("get11k10\n")); err2 != nil {
			t.Error("error writing to server: ", err2)
		}
		b := make([]byte, 6)
		if _, err3 := c.Read(b); err3 != nil {
			t.Error("error reading from server: ", err3)
		}
		if string(b) != "val11v" {
			t.Error("expected 'val11v', got: ", string(b))
		}
	})
	t.Run("TestValidSmallKeySmallValueVariable2", func(t *testing.T) {
		if _, err2 := c.Write([]byte("get11k11\n")); err2 != nil {
			t.Error("error writing to server: ", err2)
		}
		b := make([]byte, 6)
		if _, err3 := c.Read(b); err3 != nil {
			t.Error("error reading from server: ", err3)
		}
		if string(b) != "val11v" {
			t.Error("expected 'val11v', got: ", string(b))
		}
	})
	t.Run("TestValidSmallKeySmallValueVariable3", func(t *testing.T) {
		if _, err2 := c.Write([]byte("get11k12\n")); err2 != nil {
			t.Error("error writing to server: ", err2)
		}
		b := make([]byte, 6)
		if _, err3 := c.Read(b); err3 != nil {
			t.Error("error reading from server: ", err3)
		}
		if string(b) != "val11v" {
			t.Error("expected 'val11v', got: ", string(b))
		}
	})
	t.Run("TestInvalidSmallKeySmallValueVariable", func(t *testing.T) {
		if _, err2 := c.Write([]byte("get11k111\n")); err2 != nil {
			t.Error("error writing to server: ", err2)
		}
		b := make([]byte, 3)
		if _, err3 := c.Read(b); err3 != nil {
			t.Error("error reading from server: ", err3)
		}
		if string(b) != "err" {
			t.Error("expected 'err', got: ", string(b))
		}
	})
	t.Run("TestInvalidSmallKeySmallValueVariable2", func(t *testing.T) {
		if _, err2 := c.Write([]byte("get11k11a\n")); err2 != nil {
			t.Error("error writing to server: ", err2)
		}
		b := make([]byte, 3)
		if _, err3 := c.Read(b); err3 != nil {
			t.Error("error reading from server: ", err3)
		}
		if string(b) != "err" {
			t.Error("expected 'err', got: ", string(b))
		}
	})
	t.Run("TestValidBigKeyBigValue", func(t *testing.T) {
		if _, err2 := c.Write([]byte("get210firstkey01\n")); err2 != nil {
			t.Error("error writing to server: ", err2)
		}
		b := make([]byte, 11)
		if _, err3 := c.Read(b); err3 != nil {
			t.Error("error reading from server: ", err3)
		}
		if string(b) != "val16value2" {
			t.Error("expected 'val16value2', got: ", string(b))
		}
	})
	t.Run("TestValidBigKeyBigValueVariable1", func(t *testing.T) {
		if _, err2 := c.Write([]byte("get210firstkey0113\n")); err2 != nil {
			t.Error("error writing to server: ", err2)
		}
		b := make([]byte, 8)
		if _, err3 := c.Read(b); err3 != nil {
			t.Error("error reading from server: ", err3)
		}
		if string(b) != "val13val" {
			t.Error("expected 'val13val', got: ", string(b))
		}
	})
	t.Run("TestValidOnlyNumbers", func(t *testing.T) {
		if _, err2 := c.Write([]byte("get111\n")); err2 != nil {
			t.Error("error writing to server: ", err2)
		}
		b := make([]byte, 6)
		if _, err3 := c.Read(b); err3 != nil {
			t.Error("error reading from server: ", err3)
		}
		if string(b) != "val111" {
			t.Error("expected 'val111', got: ", string(b))
		}
	})
	t.Run("TestInvalidZeroKey", func(t *testing.T) {
		if _, err2 := c.Write([]byte("get0\n")); err2 != nil {
			t.Error("error writing to server: ", err2)
		}
		b := make([]byte, 3)
		if _, err3 := c.Read(b); err3 != nil {
			t.Error("error reading from server: ", err3)
		}
		if string(b) != "err" {
			t.Error("expected 'err', got: ", string(b))
		}
	})
	t.Run("TestInvalidSmallKey1", func(t *testing.T) {
		if _, err2 := c.Write([]byte("get110k\n")); err2 != nil {
			t.Error("error writing to server: ", err2)
		}
		b := make([]byte, 3)
		if _, err3 := c.Read(b); err3 != nil {
			t.Error("error reading from server: ", err3)
		}
		if string(b) != "err" {
			t.Error("expected 'err', got: ", string(b))
		}
	})
	t.Run("TestInvalidSmallKey2", func(t *testing.T) {
		if _, err2 := c.Write([]byte("get21k\n")); err2 != nil {
			t.Error("error writing to server: ", err2)
		}
		b := make([]byte, 3)
		if _, err3 := c.Read(b); err3 != nil {
			t.Error("error reading from server: ", err3)
		}
		if string(b) != "err" {
			t.Error("expected 'err', got: ", string(b))
		}
	})
	t.Run("TestInvalidBigKeyBigValue", func(t *testing.T) {
		if _, err2 := c.Write([]byte("get210firstkey0\n")); err2 != nil {
			t.Error("error writing to server: ", err2)
		}
		b := make([]byte, 3)
		if _, err3 := c.Read(b); err3 != nil {
			t.Error("error reading from server: ", err3)
		}
		if string(b) != "err" {
			t.Error("expected 'err', got: ", string(b))
		}
	})
	t.Run("TestInvalidOnlyNumbers", func(t *testing.T) {
		if _, err2 := c.Write([]byte("get1111\n")); err2 != nil {
			t.Error("error writing to server: ", err2)
		}
		b := make([]byte, 3)
		if _, err3 := c.Read(b); err3 != nil {
			t.Error("error reading from server: ", err3)
		}
		if string(b) != "err" {
			t.Error("expected 'err', got: ", string(b))
		}
	})
	t.Run("TestKeyNotFound", func(t *testing.T) {
		if _, err2 := c.Write([]byte("get11a\n")); err2 != nil {
			t.Error("error writing to server: ", err2)
		}
		b := make([]byte, 3)
		if _, err3 := c.Read(b); err3 != nil {
			t.Error("error reading from server: ", err3)
		}
		if string(b) != "nil" {
			t.Error("expected 'nil', got: ", string(b))
		}
	})
	err2 := c.Close()
	if err2 != nil {
		t.Error("error closing connection: ", err2)
	}
}

// tests the Delete handler without connection (requires store to be initialised)
// we assume these tests are run after the Put ones.
func TestDelete(t *testing.T) {
	c, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		t.Error("could not connect to server: ", err)
	}
	t.Run("TestBadMethod", func(t *testing.T) {
		if _, err2 := c.Write([]byte("de11k\n")); err2 != nil {
			t.Error("error writing to server: ", err2)
		}
		b := make([]byte, 3)
		if _, err3 := c.Read(b); err3 != nil {
			t.Error("error reading from server: ", err3)
		}
		if string(b) != "err" {
			t.Error("expected 'err', got: ", string(b))
		}
	})
	t.Run("TestBadDeleteFormat", func(t *testing.T) {
		if _, err2 := c.Write([]byte("del11k11v\n")); err2 != nil {
			t.Error("error writing to server: ", err2)
		}
		b := make([]byte, 3)
		if _, err3 := c.Read(b); err3 != nil {
			t.Error("error reading from server: ", err3)
		}
		if string(b) != "err" {
			t.Error("expected 'err', got: ", string(b))
		}
	})
	t.Run("TestValidSmallKeySmallValue", func(t *testing.T) {
		if _, err2 := c.Write([]byte("del11k\n")); err2 != nil {
			t.Error("error writing to server: ", err2)
		}
		b := make([]byte, 3)
		if _, err3 := c.Read(b); err3 != nil {
			t.Error("error reading from server: ", err3)
		}
		if string(b) != "ack" {
			t.Error("expected 'ack', got: ", string(b))
		}
	})
	t.Run("TestValidBigKeyBigValue", func(t *testing.T) {
		if _, err2 := c.Write([]byte("del210firstkey01\n")); err2 != nil {
			t.Error("error writing to server: ", err2)
		}
		b := make([]byte, 3)
		if _, err3 := c.Read(b); err3 != nil {
			t.Error("error reading from server: ", err3)
		}
		if string(b) != "ack" {
			t.Error("expected 'ack', got: ", string(b))
		}
	})
	t.Run("TestValidOnlyNumbers", func(t *testing.T) {
		if _, err2 := c.Write([]byte("del111\n")); err2 != nil {
			t.Error("error writing to server: ", err2)
		}
		b := make([]byte, 3)
		if _, err3 := c.Read(b); err3 != nil {
			t.Error("error reading from server: ", err3)
		}
		if string(b) != "ack" {
			t.Error("expected 'ack', got: ", string(b))
		}
	})
	t.Run("TestInvalidZeroKey", func(t *testing.T) {
		if _, err2 := c.Write([]byte("del0\n")); err2 != nil {
			t.Error("error writing to server: ", err2)
		}
		b := make([]byte, 3)
		if _, err3 := c.Read(b); err3 != nil {
			t.Error("error reading from server: ", err3)
		}
		if string(b) != "err" {
			t.Error("expected 'err', got: ", string(b))
		}
	})
	t.Run("TestInvalidSmallKey1", func(t *testing.T) {
		if _, err2 := c.Write([]byte("del110k\n")); err2 != nil {
			t.Error("error writing to server: ", err2)
		}
		b := make([]byte, 3)
		if _, err3 := c.Read(b); err3 != nil {
			t.Error("error reading from server: ", err3)
		}
		if string(b) != "err" {
			t.Error("expected 'err', got: ", string(b))
		}
	})
	t.Run("TestInvalidSmallKey2", func(t *testing.T) {
		if _, err2 := c.Write([]byte("get21k\n")); err2 != nil {
			t.Error("error writing to server: ", err2)
		}
		b := make([]byte, 3)
		if _, err3 := c.Read(b); err3 != nil {
			t.Error("error reading from server: ", err3)
		}
		if string(b) != "err" {
			t.Error("expected 'err', got: ", string(b))
		}
	})
	t.Run("TestInvalidBigKeyBigValue", func(t *testing.T) {
		if _, err2 := c.Write([]byte("del210firstkey0\n")); err2 != nil {
			t.Error("error writing to server: ", err2)
		}
		b := make([]byte, 3)
		if _, err3 := c.Read(b); err3 != nil {
			t.Error("error reading from server: ", err3)
		}
		if string(b) != "err" {
			t.Error("expected 'err', got: ", string(b))
		}
	})
	t.Run("TestInvalidOnlyNumbers", func(t *testing.T) {
		if _, err2 := c.Write([]byte("del1111\n")); err2 != nil {
			t.Error("error writing to server: ", err2)
		}
		b := make([]byte, 3)
		if _, err3 := c.Read(b); err3 != nil {
			t.Error("error reading from server: ", err3)
		}
		if string(b) != "err" {
			t.Error("expected 'err', got: ", string(b))
		}
	})
	t.Run("TestKeyNotFound", func(t *testing.T) {
		if _, err2 := c.Write([]byte("del11a\n")); err2 != nil {
			t.Error("error writing to server: ", err2)
		}
		b := make([]byte, 3)
		if _, err3 := c.Read(b); err3 != nil {
			t.Error("error reading from server: ", err3)
		}
		if string(b) != "err" {
			t.Error("expected 'err', got: ", string(b))
		}
	})
	err2 := c.Close()
	if err2 != nil {
		t.Error("error closing connection: ", err2)
	}
}
