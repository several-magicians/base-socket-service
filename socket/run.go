package socket

import (
	"base-socket-service/conf"
	"github.com/gofiber/fiber/v2"
	"github.com/zishang520/engine.io/v2/log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/zishang520/engine.io/v2/types"
	"github.com/zishang520/socket.io/v2/socket"
)

func Run() {
	log.DEBUG = true
	c := socket.DefaultServerOptions()
	c.SetServeClient(true)
	// c.SetConnectionStateRecovery(&socket.ConnectionStateRecovery{})
	// c.SetAllowEIO3(true)
	c.SetPingInterval(300 * time.Millisecond)
	c.SetPingTimeout(200 * time.Millisecond)
	c.SetMaxHttpBufferSize(1000000)
	c.SetConnectTimeout(1000 * time.Millisecond)
	c.SetCors(&types.Cors{
		Origin:      "*",
		Credentials: true,
	})
	socketIo := NewSocketIo()

	app := fiber.New()

	// app.Put("/socket.io", adaptor.HTTPHandler(socketIo.ServeHandler(c))) // test
	app.Get("/socket.io", adaptor.HTTPHandler(socketIo.ServeHandler(c)))
	app.Post("/socket.io", adaptor.HTTPHandler(socketIo.ServeHandler(c)))

	go app.Listen(":" + conf.Port)

	exit := make(chan struct{})
	SignalC := make(chan os.Signal)

	signal.Notify(SignalC, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		for s := range SignalC {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				close(exit)
				return
			}
		}
	}()

	<-exit
	socketIo.Close(nil)
	os.Exit(0)
}
