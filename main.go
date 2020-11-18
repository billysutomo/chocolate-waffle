package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	_ "github.com/lib/pq"

	_middleware "github.com/billysutomo/chocolate-waffle/internal/middleware"

	_projectDeliveryHttp "github.com/billysutomo/chocolate-waffle/internal/domain/project/delivery/http"
	_projectRepositoryPostgre "github.com/billysutomo/chocolate-waffle/internal/domain/project/repository/postgre"
	_projectUsecase "github.com/billysutomo/chocolate-waffle/internal/domain/project/usecase"

	_userDeliveryHttp "github.com/billysutomo/chocolate-waffle/internal/domain/user/delivery/http"
	_userRepositoryPostgre "github.com/billysutomo/chocolate-waffle/internal/domain/user/repository/postgre"
	_userUsecase "github.com/billysutomo/chocolate-waffle/internal/domain/user/usecase"

	_blockDeliveryHttp "github.com/billysutomo/chocolate-waffle/internal/domain/block/delivery/http"
	_blockRepositoryPostgre "github.com/billysutomo/chocolate-waffle/internal/domain/block/repository/postgre"
	_blockUsecase "github.com/billysutomo/chocolate-waffle/internal/domain/block/usecase"
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
	connection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)
	log.Print(connection)
	dbConn, err := sql.Open("postgres", connection)

	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	/* setup logger */
	loggerMgr := initZapLog()
	defer loggerMgr.Sync()

	/* setup middleware */
	middl := _middleware.InitMiddleware()
	r.Use(middl.CORSMiddleware())

	/* setup domain user */
	userRepo := _userRepositoryPostgre.NewPosgtgreUserRepository(dbConn, loggerMgr)
	userUcase := _userUsecase.NewUserUsecase(userRepo, loggerMgr)
	_userDeliveryHttp.NewUserHandler(r, middl, userUcase)

	/* setup domain project */
	projectRepo := _projectRepositoryPostgre.NewPosgtreProjectRepository(dbConn, loggerMgr)
	projectUcase := _projectUsecase.NewProjectUsecase(projectRepo, loggerMgr)
	_projectDeliveryHttp.NewProjectHandler(r, middl, projectUcase)

	/* setup domain block */
	blockRepo := _blockRepositoryPostgre.NewPosgtreBlockRepository(dbConn, loggerMgr)
	blockUcase := _blockUsecase.NewBlockUsecase(blockRepo, loggerMgr)
	_blockDeliveryHttp.NewBlockHandler(r, middl, blockUcase)

	r.Run(":" + viper.GetString("PORT"))
}
