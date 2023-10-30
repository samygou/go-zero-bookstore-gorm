package main

import (
	"flag"
	"fmt"
	"go-zero-bookstore/common/sdk/db/mdb/mysqlx"

	"go-zero-bookstore/app/bookstore/cmd/rpc/internal/config"
	"go-zero-bookstore/app/bookstore/cmd/rpc/internal/server"
	"go-zero-bookstore/app/bookstore/cmd/rpc/internal/svc"
	"go-zero-bookstore/app/bookstore/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/bookstore.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	// 初始化gorm
	mysqlx.Sess = mysqlx.New(c.DB.DataSource, map[string]interface{}{
		"maxOpenConns":    c.Mysql.MaxOpenConns,
		"maxIdleConns":    c.Mysql.MaxIdleConns,
		"maxConnLifeTime": c.Mysql.MaxConnLifeTime,
	})

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterBookstoreServer(grpcServer, server.NewBookstoreServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
