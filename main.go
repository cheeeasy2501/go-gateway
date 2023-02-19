package main

import (
	"net/http"

	"github.com/cheeeasy2501/go-gateway/internal/service"
	"github.com/cheeeasy2501/go-gateway/internal/transport/http/v1"
	"github.com/cheeeasy2501/go-gateway/pkg"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Тестовый код, порефакторить это г.
func main() {
	authConn := pkg.NewGRPCConnection(":1000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer authConn.CloseConnection()

	services := service.NewServices(authConn)
	controllers := v1.NewController(services)

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		controllers.Auth.
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
