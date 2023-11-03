package route

import (
	"go-svc-tpl/api/dto"
	"go-svc-tpl/internal/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, dto.Resp{
		Code: http.StatusOK,
		Msg:  "success",
		Data: "pong~",
	})
}

func SetupRouter(r *gin.RouterGroup) {
	r.GET("/:short", controller.RedirectLink())
	api := r.Group("/api")
	{
		//api.Use(controller.HandleCors())
		api.GET("/ping", Ping)
		setupLinkController(api)
		setupUserController(api)
	}
}
