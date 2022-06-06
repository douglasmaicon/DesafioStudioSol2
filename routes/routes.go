package routes

import (
	"github.com/douglasmaicon/api-go-gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	/*definir a configuração padrão e levantar o serviço gin*/
	r := gin.Default()

	r.POST("/search", controllers.Search)

	r.Run()
}
