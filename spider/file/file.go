package main

import (
	"mypkg/fs"
)

type SpiderFile struct {
	fs.File
}

func (w *SpiderFile) Write() error {
	err := w.CreateFile()
	return err
}
