package main

import (
	"fmt"
	"github.com/ipfs/go-ipfs/core"
	"github.com/ipfs/go-ipfs/core/coreunix"
	"io/ioutil"
	"log"
	"os"
	"sort"
)

func GetOldestFile(path string) string {

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

	return listfiles[0].Filename
}

func AddFile(ipfs *core.IpfsNode, file string) (string, error) {
	fi, err := os.Open(file)
	if err != nil {
		return "", err
	}

	return coreunix.Add(ipfs, fi)
}

func main() {
	filename := GetOldestFile("stream/")
	fmt.Println(filename)
	AddFile(filename)

}
