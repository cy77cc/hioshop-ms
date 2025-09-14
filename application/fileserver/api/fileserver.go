package main

import (
	"flag"
	"fmt"

	"github.com/cy77cc/hioshop_ms/application/fileserver/api/internal/config"
	"github.com/cy77cc/hioshop_ms/application/fileserver/api/internal/handler"
	"github.com/cy77cc/hioshop_ms/application/fileserver/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/fileserver.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
