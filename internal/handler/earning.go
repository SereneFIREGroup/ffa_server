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

	userID := GetRequestUserID(c)
	var req earningServices.AddEarningRequest
	if err := BindJSON(c, &req); err != nil {
		return
	}
	result := earningServices.AddEarning(spanCtx, userID, &req)
	RenderJSON(c, result, nil)
}

func ListEarning(c *gin.Context) {
	spanCtx := getCtx(c)

	userID := GetRequestUserID(c)
	var req earningServices.ListEarningRequest
	if err := BindJSON(c, &req); err != nil {
		return
	}
	obj, result := earningServices.ListEarning(spanCtx, userID, &req)
	RenderJSON(c, result, obj)
}

func Aggregation(c *gin.Context) {
	spanCtx := getCtx(c)

	userID := GetRequestUserID(c)
	obj, result := earningServices.AggregationEarning(spanCtx, userID)
	RenderJSON(c, result, obj)
}
