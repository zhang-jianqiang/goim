package main

import (
	"context"
	"github.com/Terry-Mao/goim/api/biz"
	"google.golang.org/grpc"
	"log"
)

func main() {
	dial, err := grpc.Dial(":3119", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer dial.Close()

	client := biz.NewBizClient(dial)
	res, err := client.PushMids(context.TODO(), &biz.PushMidsReq{
		Op:   1000,
		Mids: []int64{123},
		Msg:  []byte("我是rpc客户端发送的消息"),
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)
}
