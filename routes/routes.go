package routes

import (
	"goRabbitMQ/controller"

	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/offers")
	{
		grp1.POST("Publisher", func(c *gin.Context) {
			controller.PublishOffers(c)
		}) //declare queue
	}
	return r
}
