package sio

import (
	"context"
	"path/filepath"
)

// RecursiveGlob does a top level scan of each pattern to find files and directories that match the
// glob, sending every file found to the channel.  Every matching directory is scanned for files,
// and those files are sent to the channel.
//
// Only regular files are sent, links are neither processed not recursively entered.
//
// The order in which files are sent to the channel is not deterministic.
//
// Buffer controls the size of the channel, so it can be buffered or not as the caller chooses.
//
// To prevent a goroutine leak, the caller must either cancel the context (which will always clean up),
// or receive from the channel until it is closed.
//
// Any error will trigger a panic
func RecursiveGlob(ctx context.Context, buffer int, patterns ...string) <-chan string {
	ch:= make(chan string, buffer)
	go func() {
		defer close(ch)
		ctxDone := ctx.Done()
		dirStack := make([]string, 0, len(patterns))
		for _, pattern := range patterns {
			for _, name := range FilepathGlob(pattern) {
				absName := FilepathAbs(name)
				if fi := OsLstat(absName); fi.IsDir() {
					dirStack = append(dirStack, absName)
				} else if fi.Mode().IsRegular() {
					if !sendFilename(ctxDone, ch, absName) {
						return
					}
				}

			}
		}

		for len(dirStack) > 0 {
			dirName := dirStack[len(dirStack)-1]
			dirStack = dirStack[:len(dirStack)-1]

			for _, childInfo := range FReadDirAll(dirName) {
				absName := filepath.Join(dirName, childInfo.Name())
				if childInfo.IsDir() {
					dirStack = append(dirStack, absName)
				} else if childInfo.Mode().IsRegular() {
					if !sendFilename(ctxDone, ch, absName) {
						return
					}
				} else {
					// not a directory
					// not a regular file
					// ignore it
				}
			}
		}
	}()
	return ch
}

// RecursiveGlobDirs is like RecursiveGlob except it only sends directories.
//
// Note that a pattern of /foo/* will not send /foo.  To include the base directory,
// use a pattern of /foo.
func RecursiveGlobDirs(ctx context.Context, buffer int, patterns ...string) <-chan string {
	ch := make(chan string, buffer)
	go func() {
		defer close(ch)
		ctxDone := ctx.Done()
		dirStack := make([]string, 0, len(patterns))
		for _, pattern := range patterns {
			for _, name := range FilepathGlob(pattern) {
				absName := FilepathAbs(name)
				if fi := OsLstat(absName); fi.IsDir() {
					dirStack = append(dirStack, absName)
				}

			}
		}

		for len(dirStack) > 0 {
			dirName := dirStack[len(dirStack)-1]
			dirStack = dirStack[:len(dirStack)-1]
			if !sendFilename(ctxDone, ch, dirName) {
				return
			}

			for _, childInfo := range FReadDirAll(dirName) {
				if childInfo.IsDir() {
					dirStack = append(dirStack, filepath.Join(dirName, childInfo.Name()))
				}
			}
		}
	}()
	return ch
}

func sendFilename(ctxDone <-chan struct{}, ch chan<- string, fn string) bool {
	select {
	case <-ctxDone:
		return false
	case ch <- fn:
		return true
	}
}
