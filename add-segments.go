package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"

	core "github.com/ipfs/go-ipfs/core"
	coreapi "github.com/ipfs/go-ipfs/core/coreapi"
	coreiface "github.com/ipfs/go-ipfs/core/coreapi/interface"
	//coreunix "github.com/ipfs/go-ipfs/core/coreunix"
	keystore "github.com/ipfs/go-ipfs/keystore"
	//mdag "github.com/ipfs/go-ipfs/merkledag"
	repo "github.com/ipfs/go-ipfs/repo"
	config "github.com/ipfs/go-ipfs/repo/config"
	//unixfs "github.com/ipfs/go-ipfs/unixfs"

	datastore "gx/ipfs/QmXRKBQA4wXP7xWbFiZsR1GP4HV6wMDQ1aWFxZZ4uBcPX9/go-datastore"
	syncds "gx/ipfs/QmXRKBQA4wXP7xWbFiZsR1GP4HV6wMDQ1aWFxZZ4uBcPX9/go-datastore/sync"
	peer "gx/ipfs/QmZoWKhxUmZ2seW4BzX6fJkNR8hh9PsGModr7q171yq2SS/go-libp2p-peer"
	ci "gx/ipfs/QmaPbCnUMBohSGo3KnxEa2bHqyJVVeEEcwtqJAYxerieBo/go-libp2p-crypto"
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

/*func AddFile(ipfs *core.IpfsNode, file string) (string, error) {
	fi, err := os.Open(file)
	if err != nil {
		return "", err
	}

	return coreunix.Add(ipfs, fi)
}*/

func makeAPIIdent(ctx context.Context) (*core.IpfsNode, coreiface.CoreAPI, error) {
	var ident config.Identity

	sk, pk, err := ci.GenerateKeyPair(ci.RSA, 512)
	if err != nil {
		return nil, nil, err
	}

	id, err := peer.IDFromPublicKey(pk)
	if err != nil {
		return nil, nil, err
	}

	kbytes, err := sk.Bytes()
	if err != nil {
		return nil, nil, err
	}

	ident = config.Identity{
		PeerID:  id.Pretty(),
		PrivKey: base64.StdEncoding.EncodeToString(kbytes),
	}

	r := &repo.Mock{
		C: config.Config{
			Identity: ident,
		},
		D: syncds.MutexWrap(datastore.NewMapDatastore()),
		K: keystore.NewMemKeystore(),
	}
	node, err := core.NewNode(ctx, &core.BuildCfg{Repo: r})
	if err != nil {
		return nil, nil, err
	}
	api := coreapi.NewCoreAPI(node)
	return node, api, nil
}

func makeAPI(ctx context.Context) (*core.IpfsNode, coreiface.CoreAPI, error) {
	return makeAPIIdent(ctx)
}

func FileAdd(path string) (int, error) {
	ctx := context.Background()
	_, api, err := makeAPI(ctx)
	if err != nil {
		return 1, err
	}

	fi, err := os.Open(path)
	if err != nil {
		return 1, err
	}

	p, err := api.Unixfs().Add(ctx, fi)
	if err != nil {
		return 1, err
	}

	fmt.Println(p)
	return 0, err

}

func main() {
	filename := GetOldestFile("stream/")
	fmt.Println(filename)
	FileAdd(filename)

}
