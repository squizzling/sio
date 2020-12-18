package sio

import (
	"io"
	"os"
)

func OsOpen(fn string) *os.File {
	if f, err := os.Open(fn); err != nil {
		panic(err)
	} else {
		return f
	}
}

func OsOpenAppend(fn string) *os.File {
	if f, err := os.OpenFile(fn, os.O_RDWR | os.O_APPEND, 0o644); err != nil {
		panic(err)
	} else {
		_ = FSeek(f, 0, io.SeekEnd)
		return f
	}
}

func OsCreate(fn string) *os.File {
	if f, err := os.Create(fn); err != nil {
		panic(err)
	} else {
		return f
	}
}

func OsRemove(fn string) {
	if err := os.Remove(fn); err != nil {
		panic(err)
	}
}

func OsRemoveTry(fn string) {
	_ = os.Remove(fn)
}

func OsRename(from, to string) {
	if err := os.Rename(from, to); err != nil {
		panic(err)
	}
}

func OsLstat(filename string) os.FileInfo {
	if fi, err := os.Lstat(filename); err != nil {
		panic(err)
	} else {
		return fi
	}
}
