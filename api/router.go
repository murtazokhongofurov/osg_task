package api

import (
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/osg_task/api/docs"
	v1 "github.com/osg_task/api/handler/v1"
	"github.com/osg_task/api/handler/v1/middleware"
	"github.com/osg_task/api/handler/v1/tokens"
	"github.com/osg_task/internal/controller/storage"
	"github.com/osg_task/internal/pkg/config"
	"github.com/osg_task/internal/pkg/logger"
	swaggerfile "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Options struct {
	Conf           config.Config
	Logger         logger.Logger
	Storage        storage.StorageI
	CasbinEnforcer *casbin.Enforcer
}

// New ...
// @title           Online Service Group Task api
// @version         1.0
// @description     This is Task api server
// @termsOfService  2 term OSG

// @contact.name   	Murtazoxon
// @contact.url    	https://t.me/murtazokhon_gofurov
// @contact.email  	gofurovmurtazoxon@gmail.com

// @host      		localhost:7777
// @BasePath  		/v1

// @securityDefinitions.apikey BearerAuth
// @in 		header
// @name 	Authorization
func New(option *Options) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	corConfig := cors.DefaultConfig()
	corConfig.AllowAllOrigins = true
	corConfig.AllowCredentials = true
	corConfig.AllowHeaders = []string{"*"}
	corConfig.AllowBrowserExtensions = true
	corConfig.AllowMethods = []string{"*"}
	router.Use(cors.New(corConfig))

	jwtHandler := tokens.JWTHandler{
		SigninKey: option.Conf.SigninKey,
		Log:       option.Logger,
	}
	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:     option.Logger,
		Storage:    option.Storage,
		Cfg:        option.Conf,
		JwtHandler: jwtHandler,
	})

	router.Use(middleware.NewAuth(option.CasbinEnforcer, jwtHandler, config.Load()))

	router.MaxMultipartMemory = 8 << 20 // 8 Mib

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "App is running",
		})
	})

	api := router.Group("/v1")

	api.GET("login/:phone", handlerV1.Login)

	api.POST("/employee", handlerV1.CreateEmployee)
	api.GET("/employee/:id", handlerV1.GetEmployee)
	api.GET("/employees", handlerV1.GetAllEmployees)
	api.PUT("/employee", handlerV1.UpdateEmployee)
	api.DELETE("/employee/:id", handlerV1.DeleteEmployee)
	api.POST("/developer", handlerV1.CreateDeveloper)
	// media
	api.POST("/image-upload", handlerV1.UploadMedia)

	// project
	api.POST("/project", handlerV1.CreateProject)
	api.GET("/project/:id", handlerV1.GetProject)
	api.PUT("/project", handlerV1.UpdateProject)
	api.GET("/projects", handlerV1.GetAllProjects)
	api.PUT("/project/status", handlerV1.UpdateProjectStatus)
	api.DELETE("project/:id", handlerV1.DeleteProject)

	//task
	api.POST("/task", handlerV1.CreateTask)
	api.GET("/task/:id", handlerV1.GetTask)
	api.GET("/tasks", handlerV1.GetAllTask)
	api.PUT("/task", handlerV1.UpdateTask)
	api.PUT("/task/status", handlerV1.UpdateTaskStatus)
	api.DELETE("/task/:id", )
	url := ginSwagger.URL("swagger/doc.json")
	api.GET("swagger/*any", ginSwagger.WrapHandler(swaggerfile.Handler, url))

	return router
}
