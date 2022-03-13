package main

import (
	"database/sql"
	"os"
	"site/api/config/database"
	"site/api/config/env"
	"site/api/config/server"
	"site/api/internal/auth"
	errorMiddleware "site/api/internal/error"
	"site/api/internal/logger"
	"site/api/pkg/routes"

	jwt "github.com/appleboy/gin-jwt/v2"
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
	// active authentication configuration
	authMiddleware := auth.AuthHandler(colorfulLogger, postgresDB)
	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		colorfulLogger.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}
	// configure application
	app := server.Application{
		Port:           env.GetDefaultInt("PORT", 8080),
		Engine:         gin.New(),
		Routers:        make(map[string]*gin.RouterGroup),
		BasePath:       server.BasePath,
		DatabaseClient: postgresDB,
		Logger:         colorfulLogger,
		AuthModule:     authMiddleware,
	}
	app.GlobalMiddleware(errorMiddleware.ErrorHandler(colorfulLogger))
	// app.GlobalMiddleware(authMiddleware.MiddlewareFunc())
	app.Engine.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		colorfulLogger.Info("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	// setup routes
	portifolioRoute, err := routes.NewRoute(app)
	if err != nil {
		colorfulLogger.Error("portifolioRoute returned with: %v", err)
	}
	portifolioRoute.InitRoutes()
	authRoute, err := routes.NewRouteAuth(app)
	if err != nil {
		colorfulLogger.Error("authRoute returned with: %v", err)
	}
	authRoute.InitRoutes()
	// run the server
	err = app.Run()
	if err != nil {
		colorfulLogger.Error("application returned with: %v", err)
	}
}
