package grpc

import (
	"context"
	"github.com/Terry-Mao/goim/api/biz"
	"github.com/Terry-Mao/goim/internal/logic"
	log "github.com/golang/glog"
)

type bizServer struct {
	srv *logic.Logic
}

func (b *bizServer) PushMids(ctx context.Context, req *biz.PushMidsReq) (res *biz.PushMidsRes, err error) {
	log.Info("收到rpc消息，", req)
	if err = b.srv.PushMids(ctx, req.Op, req.Mids, req.Msg); err != nil {
		return
	}
	return &biz.PushMidsRes{}, nil
}
