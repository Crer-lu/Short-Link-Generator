package route

import (
	"go-svc-tpl/api/dto"
	"go-svc-tpl/internal/controller"

	"github.com/gin-gonic/gin"
)

func setupUserController(r *gin.RouterGroup) {
	lcw := UserCtlWrapper{
		ctl: controller.NewUserController(), // Factory method of ILinkController.
		// Implemented in controller package.
	}

	p := r.Group("/user")
	p.POST("/register", lcw.Register)                                // POST /user/register
	p.GET("/captcha", lcw.Captcha)                                   // GET /user/captcha
	p.POST("/logout", controller.AuthMiddleware(), lcw.Logout)       // POST /user/logout
	p.POST("/login", lcw.Login)                                      // POST /user/login
	p.GET("/info", controller.AuthMiddleware(), lcw.GetInfo)         // GET /user/info
	p.POST("/info", controller.AuthMiddleware(), lcw.UpdateInfo)     // POST /user/info
	p.POST("/passwd", controller.AuthMiddleware(), lcw.UpdatePasswd) // POST /user/passwd

}

type UserCtlWrapper struct {
	ctl controller.IUserController
}

func (w *UserCtlWrapper) Register(c *gin.Context) {
	var req dto.UserRegisterReq
	if err := dto.BindReq(c, &req); err != nil {
		dto.ResponseFail(c, err)
		return
	}
	resp, err := w.ctl.Register(c, &req)
	if err != nil {
		dto.ResponseFail(c, err)
		return
	}
	dto.ResponseSuccess(c, resp)
}

func (w *UserCtlWrapper) Captcha(c *gin.Context) {
	var req dto.UserCaptchaReq
	if err := dto.BindReq(c, &req); err != nil {
		dto.ResponseFail(c, err)
		return
	}
	resp, err := w.ctl.Captcha(c, &req)
	if err != nil {
		dto.ResponseFail(c, err)
		return
	}
	dto.ResponseSuccess(c, resp)
}

func (w *UserCtlWrapper) Logout(c *gin.Context) {
	var req dto.UserLogoutReq
	if err := dto.BindReq(c, &req); err != nil {
		dto.ResponseFail(c, err)
		return
	}
	resp, err := w.ctl.Logout(c, &req)
	if err != nil {
		dto.ResponseFail(c, err)
		return
	}
	dto.ResponseSuccess(c, resp)
}
func (w *UserCtlWrapper) Login(c *gin.Context) {
	var req dto.UserLoginReq
	if err := dto.BindReq(c, &req); err != nil {
		dto.ResponseFail(c, err)
		return
	}
	resp, err := w.ctl.Login(c, &req)
	if err != nil {
		dto.ResponseFail(c, err)
		return
	}
	dto.ResponseSuccess(c, resp)
}

func (w *UserCtlWrapper) GetInfo(c *gin.Context) {
	var req dto.UserGetInfoReq
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

func (w *UserCtlWrapper) UpdateInfo(c *gin.Context) {
	var req dto.UserUpdateInfoReq
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

func (w *UserCtlWrapper) UpdatePasswd(c *gin.Context) {
	var req dto.UserUpdatePasswdReq
	if err := dto.BindReq(c, &req); err != nil {
		dto.ResponseFail(c, err)
		return
	}
	resp, err := w.ctl.UpdatePasswd(c, &req)
	if err != nil {
		dto.ResponseFail(c, err)
		return
	}
	dto.ResponseSuccess(c, resp)
}
