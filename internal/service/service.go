package service

import (
	"context"
	"go-blog/global"
	"go-blog/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{
		ctx: ctx,
		dao: dao.New(global.DBEngine),
	}
	return svc
}
