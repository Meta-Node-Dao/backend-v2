package main

import (
	"flag"
	"fmt"

	"metaLand/app/api/internal/config"
	"metaLand/app/api/internal/handler"
	"metaLand/app/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

// If you want to run on window and debugger on local env, pls make sure you locate your dir : F:\go-Lin\backend-v2\app\api\etc (change to your own)
var configFile = flag.String("f", "F:\\go-Lin\\backend-v2\\app\\api\\etc\\meta_land.yaml", "the config file")

// var configFile = flag.String("f", "etc/meta_land.yaml", "the config file")

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
