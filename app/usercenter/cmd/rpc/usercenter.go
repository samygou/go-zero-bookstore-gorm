package main

import (
	"flag"
	"fmt"
	"go-zero-bookstore/common/sdk/db/mdb/mysqlx"

	"go-zero-bookstore/app/usercenter/cmd/rpc/internal/config"
	"go-zero-bookstore/app/usercenter/cmd/rpc/internal/server"
	"go-zero-bookstore/app/usercenter/cmd/rpc/internal/svc"
	"go-zero-bookstore/app/usercenter/cmd/rpc/usercenter"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/usercenter.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	mysqlx.Sess = mysqlx.New(c.DB.DataSource, map[string]interface{}{
		"maxOpenConns":    c.Mysql.MaxOpenConns,
		"maxIdleConns":    c.Mysql.MaxIdleConns,
		"maxConnLifeTime": c.Mysql.MaxConnLifeTime,
	})

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		usercenter.RegisterUsercenterServer(grpcServer, server.NewUsercenterServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
