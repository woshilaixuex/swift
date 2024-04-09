package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	"xihu/common/orm"
)

type Config struct {
	zrpc.RpcServerConf
	SqlServerConf *orm.Config
}
