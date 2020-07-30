package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	_urlHttpController "github.com/billysutomo/chocolate-waffle/internal/url/controller/http"
	_urlRepository "github.com/billysutomo/chocolate-waffle/internal/url/repository/postgre"
	_urlUcase "github.com/billysutomo/chocolate-waffle/internal/url/usecase"
)

// URLStruct comment
type URLStruct struct {
	Slug string `json:"slug" validate:"required,slugcheck"`
	URL  string `json:"url" validate:"required"`
}

func main() {
	r := gin.Default()

	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sql.Open(`mysql`, dsn)

	if err != nil {
		log.Fatal(err)
	}

	urlRepo := _urlRepository.NewPostgreURLRepository(dbConn)

	urlUcase := _urlUcase.NewURLService(urlRepo)
	_urlHttpController.NewURLHandler(r, urlUcase)

	r.Run()

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
