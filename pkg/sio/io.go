package sio

import (
	"io"
)

func IoCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		panic(err)
	}
}
