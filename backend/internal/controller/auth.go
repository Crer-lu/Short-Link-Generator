package controller

import (
	"errors"
	"go-svc-tpl/api/dto"
	"go-svc-tpl/utils/stacktrace"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func HandleCors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8000")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding")
		logrus.Println(c.Request.Method)
		logrus.Println(c.Request.Header.Get("Origin"))
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		}
		c.Next()
	}

}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求头
		token, err := c.Cookie("token")
		if err != nil {
			dto.ResponseFail(c, stacktrace.PropagateWithCode(err, dto.ErrPrivilege, "Please Login First"))
			c.Abort()
		}
		userID_str, err := c.Cookie("userID")
		if err != nil {
			dto.ResponseFail(c, stacktrace.PropagateWithCode(err, dto.ErrPrivilege, "Please Login First"))
			c.Abort()
		}
		userID, _ := strconv.Atoi(userID_str)
		if token == token_mapping[userID] {
			// auth success!
			c.Set("userID", userID)
			c.Set("token", token)
			c.Next()
		} else {
			dto.ResponseFail(c, stacktrace.PropagateWithCode(errors.New("AuthFailed"), dto.ErrPrivilege, "AuthFailed"))
			c.Abort()
		}
	}

}
