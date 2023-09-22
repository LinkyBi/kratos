package server

import (
	//修改v1 "base/api/helloworld/v1"
	v1 "base/api/base/v1"
	"base/internal/conf"
	"base/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// 修改greeter *service.GreeterService
func NewHTTPServer(c *conf.Server, base *service.BaseService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterBaseHTTPServer(srv, base) //修改RegisterGreeterHTTPServer
	return srv
}
