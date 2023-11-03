package route

import (
	"go-svc-tpl/api/dto"
	"go-svc-tpl/internal/controller"

	"github.com/gin-gonic/gin"
)

func setupLinkController(r *gin.RouterGroup) {
	lcw := LinkCtlWrapper{
		ctl: controller.NewLinkController(), // Factory method of ILinkController.
		// Implemented in controller package.
	}

	p := r.Group("/link")
	p.POST("/create", lcw.Create)
	p.POST("/delete", controller.AuthMiddleware(), lcw.Delete)
	p.POST("/info", controller.AuthMiddleware(), lcw.Update)
	p.GET("/info", lcw.GetInfo)
	p.GET("/list", controller.AuthMiddleware(), lcw.GetList)
}

type LinkCtlWrapper struct {
	ctl controller.ILinkController
}

func (w *LinkCtlWrapper) Create(c *gin.Context) {
	var req dto.LinkCreateReq
	if err := dto.BindReq(c, &req); err != nil {
		dto.ResponseFail(c, err)
		return
	}
	resp, err := w.ctl.Create(c, &req)
	if err != nil {
		dto.ResponseFail(c, err)
		return
	}
	dto.ResponseSuccess(c, resp)
}

func (w *LinkCtlWrapper) Delete(c *gin.Context) {
	var req dto.LinkDeleteReq
	if err := dto.BindReq(c, &req); err != nil {
		dto.ResponseFail(c, err)
		return
	}
	resp, err := w.ctl.Delete(c, &req)
	if err != nil {
		dto.ResponseFail(c, err)
		return
	}
	dto.ResponseSuccess(c, resp)
}

func (w *LinkCtlWrapper) Update(c *gin.Context) {
	var req dto.UpdateLinkInfoReq
	if err := dto.BindReq(c, &req); err != nil {
		dto.ResponseFail(c, err)
		return
	}
	resp, err := w.ctl.UpdateInfo(c, &req)
	if err != nil {
		dto.ResponseFail(c, err)
		return
	}
	dto.ResponseSuccess(c, resp)
}

func (w *LinkCtlWrapper) GetInfo(c *gin.Context) {
	var req dto.GetLinkInfoReq
	if err := dto.BindReq(c, &req); err != nil {
		dto.ResponseFail(c, err)
		return
	}
	resp, err := w.ctl.GetInfo(c, &req)
	if err != nil {
		dto.ResponseFail(c, err)
		return
	}
	dto.ResponseSuccess(c, resp)
}

func (w *LinkCtlWrapper) GetList(c *gin.Context) {
	var req dto.GetLinkListReq
	if err := dto.BindReq(c, &req); err != nil {
		dto.ResponseFail(c, err)
		return
	}
	resp, err := w.ctl.GetList(c, &req)
	if err != nil {
		dto.ResponseFail(c, err)
		return
	}
	dto.ResponseSuccess(c, resp)
}
