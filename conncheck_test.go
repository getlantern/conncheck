package conncheck

import (
	"net"
	"testing"

	"github.com/getlantern/testify/assert"
)

func TestIsOpen(t *testing.T) {
	conn, err := net.Dial("tcp", "www.google.com:80")
	if err != nil {
		t.Fatalf("Unable to dial Google")
	}
	assert.True(t, IsOpen(conn), "Conn should still be open")
	conn.Close()
	assert.False(t, IsOpen(conn), "Conn should now be closed")
}
