package main

import (
	"io"
	"io/ioutil"
	"mypkg/fs"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/wzshiming/fork"
	"gopkg.in/ffmt.v1"
)

var Seen map[string]bool
var DownLoadFile = []string{
	"img",
	"a",
	"link",
	"script",
}

var workList = make(chan string, 100)
var Wg sync.WaitGroup

func main() {
	Seen = make(map[string]bool, 10000)
	workList <- "https://m.tangeche.com"
	f := fork.NewForkBuf(50, 50)
	go func() {
		//		Wg.Wait()
		f.Join()
		close(workList)
	}()
	for newWork := range workList {
		//		Wg.Add(1)
		f.Push(func() {
			ScanLink(newWork)
		})
	}
}

func ScanLink(link string) {
	if Seen[link] {
		return
	}

	Seen[link] = true
	ffmt.Mark(link)
	file := fs.File{}
	URLAnalysis(link, &file)
	RequestLink(link, &file)
	file.CreateFile()
	//	Wg.Done()
}

func RequestLink(link string, file *fs.File) {
	up, err := url.Parse(link)
	if err != nil {
		ffmt.Mark(err)
		return
	}
	resp, err := http.Get(link)
	if err != nil {
		ffmt.Mark(err)
		return
	}
	allLinks := QueryAnalysis(up.Scheme, up.Host, link)
	for _, newLink := range allLinks {
		ffmt.Mark(newLink)
		workList <- newLink
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		if err != io.EOF {
			ffmt.Mark(err)
			return
		}
	}
	if len(b) == 0 {
		ffmt.Mark(link)
	}
	file.Content = b
	return
}

func URLAnalysis(link string, file *fs.File) {
	u, err := url.Parse(link)
	if err != nil {
		ffmt.Mark(err)
		return
	}
	if u.Path == "" || u.Path == "/" {
		u.Path = "/index.html"
	}
	index := strings.LastIndex(u.Path, "/")
	FileName := u.Path[index+1:]
	DirPath := "dist/" + u.Host + u.Path[:index+1]
	file.FileName = FileName
	file.Path = DirPath
	return
}
func QueryAnalysis(defaultScheme, defaultHost string, link string) (res []string) {
	res = []string{}
	doc, err := goquery.NewDocument(link)
	if err != nil {
		ffmt.Mark(err)
		return
	}
	for _, ele := range DownLoadFile {
		doc.Find(ele).Each(func(i int, contentSelection *goquery.Selection) {
			attr := ""
			switch ele {
			case "img", "script":
				attr = "src"
			case "a", "link":
				attr = "href"
			}
			if attr != "" {
				src, ok := contentSelection.Attr(attr)
				if ok {
					if src != "" {
						up, err := url.Parse(src)
						if err != nil {
							ffmt.Mark(err)
							return
						}
						Scheme, Host, Path := "", "", ""
						if up.Path != "" {
							if up.Path[:1] != "/" {
								Path = "/" + up.Path
							} else {
								Path = up.Path
							}
						}
						if up.Host == "" {
							Host = defaultHost
						} else {
							Host = up.Host
						}
						if up.Scheme == "" {
							Scheme = defaultScheme + "://"
						} else {
							Scheme = up.Scheme + "://"
						}
						wholePath := Scheme + Host + Path
						res = append(res, wholePath)
					}
				}
			}
		})
	}
	return
}
