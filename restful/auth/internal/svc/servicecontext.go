package svc

import (
	"kapibara-apigateway-gozero/restful/auth/internal/config"
	kmMysql "kapibara-apigateway-gozero/restful/auth/internal/models/mysql"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config     config.Config
	UsersModel kmMysql.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.MySQLConnStr)
	return &ServiceContext{
		Config:     c,
		UsersModel: kmMysql.NewUsersModel(sqlConn),
	}
}
