package sio

import (
	"path/filepath"
)

func FilepathAbs(filename string) string {
	if absoluteFilename, err := filepath.Abs(filename); err != nil {
		panic(err)
	} else {
		return absoluteFilename
	}
}

func FilepathGlob(pattern string) []string {
	if filenames, err := filepath.Glob(pattern); err != nil {
		panic(err)
	} else {
		return filenames
	}
}
