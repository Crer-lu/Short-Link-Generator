package controller

import (
	"errors"
	"go-svc-tpl/api/dto"
	"go-svc-tpl/internal/dao"
	"go-svc-tpl/internal/dao/model"
	"go-svc-tpl/utils/stacktrace"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func RedirectLink() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取请求参数
		short := ctx.Param("short")
		db := dao.DB(ctx).Table(model.LinkTable)
		var link model.Link
		result := db.Where("short = ?", short).Limit(1).Find(&link)
		if result.RowsAffected == 0 {
			// not found
			dto.ResponseFail(ctx, stacktrace.PropagateWithCode(errors.New("ErrNoShortLink"), dto.ErrNoShortLink, "ErrNoShortLink"))
		} else {
			if strings.HasPrefix(link.Origin, "http") {
			} else {
				link.Origin = "http://" + link.Origin
			}
			ctx.Redirect(http.StatusTemporaryRedirect, link.Origin)
		}
	}
}
