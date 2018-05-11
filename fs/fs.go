package fs

import (
	"io/ioutil"
	"os"

	"gopkg.in/ffmt.v1"
)

type File struct {
	FileName string
	Path     string
	Content  []byte
}

func (f *File) CreateFile() error {
	//	info := f.ReadInfo()
	//	if f.FileName ==
	err := MakeDir(f.Path)
	if err != nil {
		ffmt.Mark(err)
		return err
	}
	fp := f.Path + f.FileName
	//	if len(f.Content) == 0 {
	//		//		ffmt.Mark("内容为空")
	//		return errors.New("内容为空")
	//	}
	err = ioutil.WriteFile(fp, f.Content, 0644)
	if err != nil {
		ffmt.Mark(err)
		return err
	}
	return nil
}

func MakeDir(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	return err
}
