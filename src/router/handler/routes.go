package handler

import (
	"net/http"

	"golang-101-master/src/service/info"
	_ "golang-101-master/src/swagger/docs"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	//"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type route struct {
	Name        string
	Description string
	Method      string
	Pattern     string
	Endpoint    gin.HandlerFunc
	//Validation  gin.HandlerFunc
}

type Routes struct {
	transaction []route
}

// @titleAlumni-Kmitl
// @vertion 1.0A
// @description This is a My Shop API document from Backend-Shortcourse workshop

// @host localhost:8080
// @scheme http

// @tag.name Find My Address
// @tag.description create read update delete product

func (r Routes) InitTransactionRoute() http.Handler {

	info := info.NewEndpoint()

	r.transaction = []route{
		{
			Name:        "InsertUserInfo",
			Description: "Insert : User",
			Method:      http.MethodPost,
			Pattern:     "/createUser",
			Endpoint:    info.CreateUserEndpoint,
			//Validation:  v.none,
		},
		{
			Name:        "SearchInfo",
			Description: "Search : Info",
			Method:      http.MethodPost,
			Pattern:     "/info",
			Endpoint:    info.InfoEndpoint,
			//Validation:  v.none,
		},
		{
			Name:        "showGeography",
			Description: "Search : Geography",
			Method:      http.MethodPost,
			Pattern:     "/showGeography",
			Endpoint:    info.ShowGeographyEndpoint,
			//Validation:  v.none,
		},
	}
	ro := gin.New()
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	ro.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	// ro.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//ro.Use(cors.New(cors.Config{
	//	AllowOrigins:     []string{"*"},
	//	AllowMethods:     []string{"POST", "GET"},
	//	AllowHeaders:     []string{"Content-Type", "Authorization"},
	//	AllowCredentials: true,
	//}))

	store := ro.Group("/api")

	for _, e := range r.transaction {
		store.Handle(e.Method, e.Pattern, e.Endpoint)
	}

	return ro
}
