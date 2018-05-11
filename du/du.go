package main

import (
	"flag"
	"io/ioutil"
	"path/filepath"
	"sync"
	"time"

	"gopkg.in/ffmt.v1"
	"gopkg.in/wzshiming/fork.v2"
)

var wg sync.WaitGroup
var GR = fork.NewForkBuf(100, 100000)

func main() {
	flag.Parse()
	roots := flag.Args()
	filesizes := make(chan int64, 10)
	exit := make(chan struct{}) //接收退出指令
	ticker := time.NewTicker(10 * time.Millisecond)
	var nbytes, nfiles int64
	for _, root := range roots {
		wg.Add(1)
		//		go WalkDir(root, filesizes)
		GR.Push(func() {
			WalkDir(root, filesizes)
		})
	}
	go func() {
		wg.Wait()
		GR.Join()
		exit <- struct{}{}
		close(filesizes)
	}()
loop:
	for {
		select {
		case size := <-filesizes:
			nbytes += size
			nfiles++
		case <-ticker.C:
			PrintFile(nbytes, nfiles)
		case <-exit:
			close(exit)
			break loop
		default:
		}

	}
	PrintFile(nbytes, nfiles)
	ffmt.Mark("执行完毕")
}

/*
	遍历目录文件
*/
func WalkDir(dir string, filesizes chan<- int64) {
	defer wg.Done()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		ffmt.Mark(err)
		return
	}
	for _, entry := range entries {
		if entry.IsDir() {
			subpath := filepath.Join(dir, entry.Name())
			wg.Add(1)
			GR.Push(func() {
				WalkDir(subpath, filesizes)
			})
		} else {
			filesizes <- entry.Size()
		}
	}
}

/*
	打印文件
*/

func PrintFile(nbytes, nfiles int64) {
	ffmt.Printf("%d files and %.2f GB\n", nfiles, float64(nbytes)/(1024*1024*1024))
	return
}
