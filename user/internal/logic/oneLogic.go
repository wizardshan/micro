package logic

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"user/internal/svc"
	"user/internal/types"
)

type OneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OneLogic {
	return &OneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OneLogic) One(req *types.RequestUser) (resp *types.ResponseUser, err error) {

	user, err := l.svcCtx.UsersModel.FindOne(l.ctx, req.Id)
	fmt.Println(err)
	resp = new(types.ResponseUser)
	resp.Id = user.Id
	resp.Nickname = user.Nickname
	return
}
