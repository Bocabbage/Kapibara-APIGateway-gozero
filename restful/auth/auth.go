package main

import (
	"flag"
	"fmt"

	kerrors "kapibara-apigateway-gozero/internal/errors"
	"kapibara-apigateway-gozero/restful/auth/internal/config"
	"kapibara-apigateway-gozero/restful/auth/internal/handler"
	"kapibara-apigateway-gozero/restful/auth/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/auth-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// Add Error handler
	httpx.SetErrorHandler(kerrors.HttpErrorHandler)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
