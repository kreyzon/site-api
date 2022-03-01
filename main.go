package main

import (
	"database/sql"
	"os"
	"site/api/config/database"
	"site/api/config/env"
	"site/api/config/server"
	errorMiddleware "site/api/internal/error"
	"site/api/internal/logger"
	portifolioRoute "site/api/pkg/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() { // create external dependencies
	colorfulLogger, err := logger.NewColorful(logger.White, logger.Yellow, logger.Red)
	if err != nil {
		colorfulLogger.Fatal("fatal error initializing logger: %v\n", err)
	}
	var postgresDB *gorm.DB
	if os.Getenv("DATABASE_URL") != "" {
		sqlDB, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
		if err != nil {
			colorfulLogger.Fatal("Error opening database: %q", err)
		}
		postgresDB, err = gorm.Open(postgres.New(postgres.Config{
			Conn: sqlDB,
		}), &gorm.Config{})
		if err != nil {
			colorfulLogger.Fatal("fatal error initializing database: %v\n", err)
		}
	} else {
		postgresDB, err = database.New(database.ConfigFromEnv)
		if err != nil {
			colorfulLogger.Fatal("fatal error initializing database: %v\n", err)
		}
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
