package sio

import (
	"os"

	"github.com/edsrzf/mmap-go"
)

func MMapMap(f *os.File) mmap.MMap {
	if m, err := mmap.Map(f, mmap.RDONLY, 0); err != nil {
		panic(err)
	} else {
		return m
	}
}

func MMapMapWrite(f *os.File) mmap.MMap {
	if m, err := mmap.Map(f, mmap.RDWR, 0); err != nil {
		panic(err)
	} else {
		return m
	}
}

func MMapMapRegion(f *os.File, length, offset uint64) mmap.MMap {
	if m, err := mmap.MapRegion(f, int(length), mmap.RDONLY, 0, int64(offset)); err != nil {
		panic(err)
	} else {
		return m
	}
}

func MMapUnmap(m mmap.MMap) {
	if err := m.Unmap(); err != nil {
		panic(err)
	}
}

