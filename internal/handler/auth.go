package handler

import (
	"github.com/gin-gonic/gin"
	userServices "github.com/serenefiregroup/ffa_server/internal/services/user"
)

func VerifyPhone(c *gin.Context) {
	spanCtx := getCtx(c)

	ip := c.RemoteIP()
	var req userServices.VerifyPhoneRequest
	if err := BindJSON(c, &req); err != nil {
		return
	}
	result := userServices.VerifyPhone(spanCtx, &req, ip)
	RenderJSON(c, result, nil)
}

func SignUp(c *gin.Context) {
	spanCtx := getCtx(c)

	var req userServices.SignUpRequest
	if err := BindJSON(c, &req); err != nil {
		return
	}
	obj, result := userServices.SignUp(spanCtx, &req)
	RenderJSON(c, result, obj)
}

func SignIn(c *gin.Context) {
	spanCtx := getCtx(c)

	var req userServices.SignInRequest
	if err := BindJSON(c, &req); err != nil {
		return
	}
	obj, result := userServices.SignIn(spanCtx, &req)
	RenderJSON(c, result, obj)
}

func SignOut(c *gin.Context) {
	spanCtx := getCtx(c)

	userUUID := GetRequestUserID(c)
	result := userServices.SignOut(spanCtx, userUUID)
	RenderJSON(c, result, nil)
}

func Me(c *gin.Context) {
	spanCtx := getCtx(c)

	userUUID := GetRequestUserID(c)
	token := GetRequestToken(c)
	obj, result := userServices.Me(spanCtx, userUUID, token)
	RenderJSON(c, result, obj)
}
