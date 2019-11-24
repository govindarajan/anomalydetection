package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/govindarajan/anomalydetection/anomix/cmd/http/middlewares"
	"github.com/govindarajan/anomalydetection/anomix/internal/configmanager"
	"github.com/govindarajan/anomalydetection/anomix/pkg/types"
	logger "github.com/govindarajan/anomalydetection/log"
	"github.com/govindarajan/anomalydetection/store"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	host = flag.String("host", "0.0.0.0", "Host ip")
	port = flag.String("port", "8080", "Host port")
	// ServiceDirectory the directory where the service runs
	ServiceDirectory string
)

func main() {
	flag.Parse()
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.BodyLimit("1024KB"))
	e.Use(middleware.Secure())
	e.Use(middleware.RemoveTrailingSlash())
	e.Use(middlewares.RequestID)
	e.Use(middlewares.Method)
	// initialize empty context
	ctx := types.NewContext(nil, fmt.Sprint(os.Getpid()))
	// initialize configuration manager
	if err := configmanager.InitConfig(ServiceDirectory); err != nil {
		log.Println("Failed initializing configmanager. Error = ", err)
		os.Exit(1)
	}
	config := configmanager.GetConfig()
	// Initialize Logger
	err := logger.InitLogger(config.ProcessName, config.LogLevel)
	if err != nil {
		log.Println("Error initializing Logger. Error = ", err)
		os.Exit(1)
	}

	// Initialize Store
	db, err := store.NewSQLite(config.SqliteFile)
	if err != nil {
		logger.Error(err, "Exiting process")
		os.Exit(2)
	}
	store.InitStore(db)

	logger.Info(ctx, "Starting Server...")
	// adding routes
	AddRoutes(e)

	if err := e.Start(fmt.Sprintf("%s:%s", *host, *port)); err != nil {
		logger.Critical(ctx, "Failed to start server!", err)
		os.Exit(1)
	}
}
