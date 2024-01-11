package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/mashbens/restfull2/app/modules"
	"github.com/mashbens/restfull2/config"

	"github.com/mashbens/restfull2/api"

	"github.com/mashbens/restfull2/util"
)

func main() {

	conf := config.GetConfig()
	dbCon := util.NewConnectionDatabase(conf)

	controllers := modules.RegisterModules(dbCon, conf)
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	api.RegisterRoutes(e, &controllers)

	go func() {
		if err := e.Start(":8088"); err != nil {
			log.Fatal(err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	defer dbCon.CloseConnection()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
