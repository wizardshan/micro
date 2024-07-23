package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"user/internal/config"
	"user/model"
)

type ServiceContext struct {
	Config     config.Config
	UsersModel model.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DB.DNS)
	return &ServiceContext{
		Config:     c,
		UsersModel: model.NewUsersModel(conn, c.Cache),
	}
}
