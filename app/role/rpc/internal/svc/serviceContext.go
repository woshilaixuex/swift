package svc

import (
	"xihu/app/role/rpc/internal/config"
	"xihu/common/orm"
)

type ServiceContext struct {
	Config    config.Config
	RoleModel *orm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		RoleModel: orm.MustNewSqlServer(c.SqlServerConf),
	}
}
