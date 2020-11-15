package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	_ "github.com/lib/pq"

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

	middl := _middleware.InitMiddleware()
	r.Use(middl.CORSMiddleware())

	userRepo := _repositoryPostgre.NewPosgtgreUserRepository(dbConn)
	userUcase := _usecase.NewUserUsecase(userRepo)
	_deliveryHttp.NewUserHandler(r, middl, userUcase)

	projectRepo := _repositoryPostgre.NewPosgtreProjectRepository(dbConn)
	projectUcase := _usecase.NewProjectUsecase(projectRepo)
	_deliveryHttp.NewProjectHandler(r, middl, projectUcase)

	blockRepo := _repositoryPostgre.NewPosgtreBlockRepository(dbConn)
	blockUcase := _usecase.NewBlockUsecase(blockRepo)
	_deliveryHttp.NewBlockHandler(r, middl, blockUcase)

	r.Run(":" + viper.GetString("PORT"))
}
