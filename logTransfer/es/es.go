package es

import (
	"fmt"
	"strings"

	"github.com/olivere/elastic/v7"
)

var (
	client *elastic.Client
)

// Init 初始化连接
func Init(addr string) (err error) {
	if !strings.HasPrefix(addr, "http://") {
		addr = "http://" + addr
	}
	client, err = elastic.NewClient(elastic.SetURL(addr))
	if err != nil {
		fmt.Printf("es connect failed, err:%v\n", err)
		return
	}
	return
}
