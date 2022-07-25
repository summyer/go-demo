package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var data, err = ioutil.ReadFile("e://b.txt")
	if err != nil {
		log.Fatal(err)
	}
	var content = string(data)
	fmt.Println(content)
	var bytes = []byte("this is new line")
	ioutil.WriteFile("e://b.txt", bytes, fs.ModePerm)
	file, err := os.OpenFile("e://b.txt", os.O_APPEND, fs.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	if _, err = file.WriteString("new data that wasn't there originally\n"); err != nil {
		panic(err)
	}
	f1, err := ioutil.TempDir("", "go-file-*")
	//C:\Users\summyer\AppData\Local\Temp\go-file-1777516050C
	//dir如果指定不为空，则指定的目录需要存在
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(f1)
	//f2, err := ioutil.TempFile("", "a-*")
	//C:\Users\summyer\AppData\Local\Temp\a-1946641848

	f2, err := ioutil.TempFile(f1, "a-*")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(f2.Name())
}
