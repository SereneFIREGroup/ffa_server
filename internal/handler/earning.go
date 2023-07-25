package handler

import (
	"github.com/gin-gonic/gin"
	earningServices "github.com/serenefiregroup/ffa_server/internal/services/earning"
)

func ListEarningCategory(c *gin.Context) {
	spanCtx := getCtx(c)

	obj, result := earningServices.ListEarningCategory(spanCtx)
	RenderJSON(c, result, obj)
}

func AddEarning(c *gin.Context) {
	spanCtx := getCtx(c)

	userUUID := GetRequestUserID(c)
	var req earningServices.AddEarningRequest
	if err := BindJSON(c, &req); err != nil {
		return
	}
	result := earningServices.AddEarning(spanCtx, userUUID, &req)
	RenderJSON(c, result, nil)
}

func ListEarning(c *gin.Context) {
	spanCtx := getCtx(c)

	userUUID := GetRequestUserID(c)
	var req earningServices.ListEarningRequest
	if err := BindJSON(c, &req); err != nil {
		return
	}
	obj, result := earningServices.ListEarning(spanCtx, userUUID, &req)
	RenderJSON(c, result, obj)
}
