package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/djimenez/iconv-go"
	"os"
	"path/filepath"
	"strings"
)


func check(e error) {
	if e != nil {
		println(e.Error())
	}
}


func main() {
	var files []string

	pathSrc := "e:/Downloads/help/HTML/prog"
	err := filepath.Walk(pathSrc, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".htm" {
			files = append(files, path)
			return nil

		} else {
			return nil
		}

	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Files:")
	for _, file := range files {
		fmt.Println(file)
		dat, err := os.Open(file)
		check(err)
		doc, _ := goquery.NewDocumentFromReader(dat)
		body, _ := doc.Find("body").Html()  // .Text() for plain text
		h3 := doc.Find("h3").First().Text()
		h4 := doc.Find("h4").First().Text()
		h5 := doc.Find("h5").First().Text()
		dat.Close()
		h := ""
		if h3 != "" {
			h = h3
		} else if h4 != "" {
			h = h4
		} else if h5 != "" {
			h = h5
		}
		if h != "" {
			fmt.Println(h)
			dirname, _ := iconv.ConvertString(h, "windows-1251", "utf-8")

			// Replace wrong filesystem symbols (maybe regexp to do)
			dirname = strings.ReplaceAll(dirname, "\n", "")
			dirname = strings.ReplaceAll(dirname, ":", "")

			err = os.Mkdir(dirname, os.ModePerm)
			check(err)
			filename := dirname + ".txt"
			f, err := os.Create(filepath.Join(dirname, filename))
			check(err)
			body, _ = iconv.ConvertString(body, "windows-1251", "utf-8")
			_, err = f.WriteString(body)
			check(err)

			f.Sync()
			f.Close()
		} else {
			fmt.Println("Can't get html header.")
		}


	}
}
