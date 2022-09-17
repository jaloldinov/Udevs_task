package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jaloldinov/Udevs_task/api_gateway/config"
	"github.com/jaloldinov/Udevs_task/api_gateway/pkg/logger"
	"github.com/jaloldinov/Udevs_task/api_gateway/services"

	// @Summary Book-Store
	// @Description requests to author-service, category-service and book-service.
	// @Produce json
	// @Param body body controllers.LoginParams true "body参数"
	// @Success 200 {string} string "ok" "返回用户信息"
	// @Failure 400 {string} string "err_code：10002 参数错误； err_code：10003 校验错误"
	// @Failure 401 {string} string "err_code：10001 登录失败"
	// @Failure 500 {string} string "err_code：20001 服务错误；err_code：20002 接口错误；err_code：20003 无数据错误；err_code：20004 数据库异常；err_code：20005 缓存异常"
	// @Router /user/person/login [post]
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/jaloldinov/Udevs_task/api_gateway/api/docs"
	v1 "github.com/jaloldinov/Udevs_task/api_gateway/api/handlers/v1"
)

type RouterOptions struct {
	Log      logger.Logger
	Cfg      config.Config
	Services services.ServiceManager
}

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func New(opt *RouterOptions) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AllowHeaders = append(config.AllowHeaders, "*")

	router.Use(cors.New(config))
	// router.Use(MaxAllowed(100))

	handlerV1 := v1.New(&v1.HandlerV1Options{
		Log:      opt.Log,
		Cfg:      opt.Cfg,
		Services: opt.Services,
	})

	apiV1 := router.Group("/v1")

	// author
	apiV1.POST("/author", handlerV1.CreateAuthor)
	apiV1.GET("/author", handlerV1.GetAllAuthor)
	apiV1.GET("/author/:author_id", handlerV1.GetAuthor)
	apiV1.PUT("/author/:author_id", handlerV1.UpdateAuthor)
	apiV1.DELETE("/author/:author_id", handlerV1.DeleteAuthor)

	// // category
	apiV1.POST("category", handlerV1.CreateCategory)
	apiV1.GET("category", handlerV1.GetAllCategory)
	apiV1.GET("category/:category_id", handlerV1.GetCategory)
	apiV1.PUT("category/:category_id", handlerV1.UpdateCategory)
	apiV1.DELETE("category/:category_id", handlerV1.DeleteCategory)

	// // book

	// swagger
	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}

func MaxAllowed(n int) gin.HandlerFunc {
	sem := make(chan struct{}, n)
	acquire := func() { sem <- struct{}{} }
	release := func() { <-sem }
	return func(c *gin.Context) {
		acquire()       // before request
		defer release() // after request
		c.Next()

	}
}
