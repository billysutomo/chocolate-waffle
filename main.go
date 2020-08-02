package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	_middleware "github.com/billysutomo/chocolate-waffle/internal/middleware"

	_authControllerHttp "github.com/billysutomo/chocolate-waffle/internal/domain/auth/controller/http"

	_urlControllerHttp "github.com/billysutomo/chocolate-waffle/internal/domain/url/controller/http"
	_urlRepository "github.com/billysutomo/chocolate-waffle/internal/domain/url/repository/postgre"
	_urlUcase "github.com/billysutomo/chocolate-waffle/internal/domain/url/usecase"
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
	connection := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	log.Print(connection)
	dbConn, err := pgx.Connect(context.Background(), connection)

	if err != nil {
		log.Fatal(err)
	}

	middl := _middleware.InitMiddleware()
	r.Use(middl.CORSMiddleware())

	_authControllerHttp.NewURLHandler(r)

	urlRepo := _urlRepository.NewPostgreURLRepository(dbConn)
	urlUcase := _urlUcase.NewURLUsecase(urlRepo)
	_urlControllerHttp.NewURLHandler(r, urlUcase)

	r.Run(":" + viper.GetString("PORT"))

	// translator := en.New()
	// uni := ut.New(translator, translator)

	// trans, found := uni.GetTranslator("en")
	// if !found {
	// 	log.Fatal("translator not found")
	// }

	// var v *validator.Validate
	// v = validator.New()

	// en_translations.RegisterDefaultTranslations(v, trans)

	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "Short Urls for your code garden",
	// 	})
	// })

	// r.GET("/url/:id", func(c *gin.Context) {
	// 	// TODO: get a short url by id
	// })

	// r.GET("/:id", func(c *gin.Context) {
	// 	// TODO: redirect to url
	// })

	// r.POST("/url", func(c *gin.Context) {
	// 	// TODO: create a short url
	// 	var url URLStruct
	// 	c.BindJSON(&url)

	// 	_ = v.RegisterValidation("slugcheck", func(fl validator.FieldLevel) bool {
	// 		pass, _ := regexp.MatchString("^[a-z0-9]+(?:-[a-z0-9]+)*$", fl.Field().String())
	// 		return pass
	// 	})

	// 	_ = v.RegisterTranslation("slugcheck", trans, func(ut ut.Translator) error {
	// 		return ut.Add("slugcheck", "{0} is not strong enough", true)
	// 	}, func(ut ut.Translator, fe validator.FieldError) string {
	// 		t, _ := ut.T("slugcheck", fe.Field())
	// 		return t
	// 	})

	// 	err := v.Struct(url)

	// 	if err != nil {
	// 		for _, e := range err.(validator.ValidationErrors) {
	// 			fmt.Println(e.Translate(trans))
	// 		}

	// 	}

	// 	c.JSON(200, gin.H{
	// 		"data": url,
	// 	})

	// })

}
