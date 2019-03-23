package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"go.etcd.io/etcd/etcdserver/etcdserverpb"

	"github.com/golang/protobuf/proto"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s file\n", os.Args[0])
		return
	}

	b, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	var m etcdserverpb.MemberListResponse
	if err := proto.Unmarshal(b, &m); err != nil {
		log.Fatalf("failed to unmarshal: %v", err)
	}

	fmt.Fprintf(os.Stdout, "%+v", m)
}
