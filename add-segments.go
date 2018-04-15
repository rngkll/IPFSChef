package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"

	"github.com/ipfs/go-ipfs-api"
)

var sh *shell.Shell

func GetOldestFile(path string, n int) string {

	type Fileinfo struct {
		Timestamp int64
		Filename  string
	}

	listfiles := []Fileinfo{}
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fileinfo, err := os.Stat(path + f.Name())
		if err != nil {
			log.Fatal(err)
		}
		filetimestamp := fileinfo.ModTime().Unix()
		filename := f.Name()
		listfiles = append(listfiles, Fileinfo{filetimestamp, filename})

	}

	//fmt.Println(listfiles)

	sort.Slice(listfiles, func(i, j int) bool {
		return listfiles[i].Timestamp < listfiles[j].Timestamp
	})
	//fmt.Println(listfiles)
	//fmt.Println(listfiles[0].Filename)

	return listfiles[n].Filename
}

func fileAdd(filename string) (string, string) {
	sh = shell.NewShell("localhost:5001")

	fi, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	hash, err := sh.Add(fi)
	if err != nil {
		log.Fatal(err)
	}

	// Temp code to remove file
	err = os.Remove(filename)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(filename, hash)
	return filename, hash

}

func main() {
	filename := GetOldestFile("stream/", 0)
	go fileAdd("stream/" + filename)

}
