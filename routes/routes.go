package routes

import (
	"github.com/douglasmaicon/api-go-gin/controllers"
	"github.com/gin-gonic/gin"

	docs "github.com/douglasmaicon/api-go-gin/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func HandleRequests() {
	/*definir a configuração padrão e levantar o serviço gin*/
	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/"

	r.POST("/search", controllers.Search)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run()
}
