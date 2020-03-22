package ioutils

import "io"

func CloseQuietly(c io.Closer) {
	if c != nil {
		_ = c.Close()
	}
}
