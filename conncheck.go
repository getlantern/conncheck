// package conncheck provides utilities for checking the status of a net.Conn.
package conncheck

import (
	"net"
	"time"
)

// IsOpen checks whether the net.Conn is still open, based on suggestions from
// http://stackoverflow.com/questions/12741386/how-to-know-tcp-connection-is-closed-in-golang-net-package.
// Note - this function results in the read deadline being set back to the
// default value.
func IsOpen(c net.Conn) bool {
	c.SetReadDeadline(time.Now())
	_, err := c.Read([]byte{})
	neterr, ok := err.(net.Error)
	if ok {
		if neterr.Timeout() {
			return true
		}
	}
	return false
}
