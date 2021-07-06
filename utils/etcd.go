package utils

import (
	"context"
	"time"

	"github.com/astaxie/beego"
	"github.com/coreos/etcd/clientv3"
)

// EtcdConn 连接etcd
func etcdConn() (client *clientv3.Client, err error) {
	client, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{GEtcdHost + ":" + GEtcdPort},
		DialTimeout: time.Second})
	if err != nil {
		return
	}

	return
}

// PutConf 增加/修改etcd项
func PutConf(key, val string) (rsp *clientv3.PutResponse, err error) {

	conn, err := etcdConn()
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	rsp, err = conn.Put(ctx, key, val)
	cancel()
	if err != nil {
		beego.Info("存入etcd失败", err)
		return
	}
	return
}

// DelConf 删除etcd配置项
func DelConf(key string) (rsp *clientv3.DeleteResponse, err error) {
	conn, err := etcdConn()
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	rsp, err = conn.Delete(ctx, key)
	cancel()
	if err != nil {
		beego.Info("删除etcd失败", err)
		return
	}
	return
}
