package main

import (
	"site/api/config/database"
	"site/api/config/env"
	"site/api/config/server"
	errorMiddleware "site/api/internal/errors"
	"site/api/internal/logger"
	portifolioRoute "site/api/pkg/routes"

	"github.com/gin-gonic/gin"
)

func main() { // create external dependencies
	colorfulLogger, err := logger.NewColorful(logger.White, logger.Yellow, logger.Red)
	if err != nil {
		colorfulLogger.Fatal("fatal error initializing logger: %v\n", err)
	}
	postgresDB, err := database.New(database.ConfigFromEnv)
	if err != nil {
		colorfulLogger.Fatal("fatal error initializing database: %v\n", err)
	}
	// configure application
	app := server.Application{
		Port:           env.GetDefaultInt("SERVER_PORT", 8080),
		Engine:         gin.New(),
		Routers:        make(map[string]*gin.RouterGroup),
		BasePath:       server.BasePath,
		DatabaseClient: postgresDB,
		Logger:         colorfulLogger,
	}
	app.GlobalMiddleware(errorMiddleware.ErrorHandler(colorfulLogger))

	portifolioRoute, err := portifolioRoute.NewRoute(app)
	if err != nil {
		colorfulLogger.Error("application returned with: %v", err)
	}
	portifolioRoute.InitRoutes()
	// run the server
	err = app.Run()
	if err != nil {
		colorfulLogger.Error("application returned with: %v", err)
	}
}
