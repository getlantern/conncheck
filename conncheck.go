// package conncheck provides utilities for checking the status of a net.Conn.
package conncheck

import (
	"net"
	"time"
)

var (
	zeroTime time.Time
)

// IsOpen checks whether the net.Conn is still open, based on suggestions from
// http://stackoverflow.com/questions/12741386/how-to-know-tcp-connection-is-closed-in-golang-net-package.
// Note - this function results in the read deadline being set back to the
// default value.
func IsOpen(conn net.Conn) bool {
	conn.SetReadDeadline(time.Now())
	defer conn.SetReadDeadline(zeroTime)
	_, err := conn.Read([]byte{})
	neterr, ok := err.(net.Error)
	if ok {
		if neterr.Timeout() {
			return true
		}
	}
	return false
}
