package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	_middleware "github.com/billysutomo/chocolate-waffle/internal/middleware"

	_deliveryHttp "github.com/billysutomo/chocolate-waffle/internal/delivery/http"
	_repositoryPostgre "github.com/billysutomo/chocolate-waffle/internal/repository/postgre"
	_usecase "github.com/billysutomo/chocolate-waffle/internal/usecase"
)

func init() {
	viper.SetConfigFile(`.env`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetString("MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
		log.Println("Service RUN on RELEASE mode")
	} else {
		log.Println("Service RUN on DEBUG mode")
	}
}

func initZapLog() *zap.Logger {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, _ := config.Build()
	return logger
}

func main() {
	r := gin.Default()

	dbHost := viper.GetString(`DB_HOST`)
	dbPort := viper.GetString(`DB_PORT`)
	dbUser := viper.GetString(`DB_USER`)
	dbPass := viper.GetString(`DB_PASS`)
	dbName := viper.GetString(`DB_NAME`)
	connection := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	log.Print(connection)
	dbConn, err := pgx.Connect(context.Background(), connection)

	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close(context.Background())

	/* setup logger */
	loggerMgr := initZapLog()
	defer loggerMgr.Sync()

	/* setup middleware */
	middl := _middleware.InitMiddleware()
	r.Use(middl.CORSMiddleware())

	/* setup domain user */
	userRepo := _repositoryPostgre.NewPosgtgreUserRepository(dbConn, loggerMgr)
	userUcase := _usecase.NewUserUsecase(userRepo, loggerMgr)
	_deliveryHttp.NewUserHandler(r, middl, userUcase)

	/* setup domain project */
	projectRepo := _repositoryPostgre.NewPosgtreProjectRepository(dbConn, loggerMgr)
	projectUcase := _usecase.NewProjectUsecase(projectRepo, loggerMgr)
	_deliveryHttp.NewProjectHandler(r, middl, projectUcase)

	/* setup domain block */
	blockRepo := _repositoryPostgre.NewPosgtreBlockRepository(dbConn, loggerMgr)
	blockUcase := _usecase.NewBlockUsecase(blockRepo, loggerMgr)
	_deliveryHttp.NewBlockHandler(r, middl, blockUcase)

	r.Run(":" + viper.GetString("PORT"))
}
