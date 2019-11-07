// +build !appengine,!gopherjs

package logrus

import (
	"io"
	"os"

	"github.com/golang/crypto/ssh/terminal"
)

func checkIfTerminal(w io.Writer) bool {
	switch v := w.(type) {
	case *os.File:
		return terminal.IsTerminal(int(v.Fd()))
	default:
		return false
	}
}
