package main

import (
	"flag"
	"github.com/zeromicro/go-zero/core/logx"

	"wczero/services/user/api/internal/config"
	"wczero/services/user/api/internal/handler"
	"wczero/services/user/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/user-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	logx.MustSetup(c.Log)
	defer logx.Close()

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	// fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	logx.Infof("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
