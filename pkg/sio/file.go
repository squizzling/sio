package sio

import (
	"io"
	"os"
)

func FSeek(f *os.File, o int64, w int) int64 {
	if n, err := f.Seek(o, w); err != nil {
		panic(err)
	} else {
		return n
	}
}

func FRead(f *os.File, buf []byte) int {
	if n, err := f.Read(buf); err != nil && err != io.EOF {
		panic(err)
	} else {
		return n
	}
}

func FWrite(f *os.File, buf []byte) {
	remaining := len(buf)
	for remaining > 0 {
		if n, err := f.Write(buf); err != nil {
			panic(err)
		} else if n == 0 {
			panic("zero write")
		} else {
			remaining -= n
			buf = buf[n:]
		}
	}
}

func FCopy(fDst, fSrc *os.File, buf []byte) int {
	n := FRead(fSrc, buf)
	FWrite(fDst, buf[:n])
	return n
}

func FClose(f *os.File) {
	if err := f.Close(); err != nil {
		panic(err)
	}
}
