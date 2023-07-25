package handler

import (
	"github.com/gin-gonic/gin"
	fourMoneyServices "github.com/serenefiregroup/ffa_server/internal/services/four_money"
)

func ListFourMoneyCategory(c *gin.Context) {
	spanCtx := getCtx(c)

	obj, result := fourMoneyServices.ListFourMoneyCategory(spanCtx)
	RenderJSON(c, result, obj)
}

func AddFourMoney(c *gin.Context) {
	spanCtx := getCtx(c)

	familyUUID := getFamilyULID(c)
	userUUID := GetRequestUserID(c)
	var req fourMoneyServices.AddFourMoneyRequest
	if err := BindJSON(c, &req); err != nil {
		return
	}
	result := fourMoneyServices.AddFourMoney(spanCtx, familyUUID, userUUID, &req)
	RenderJSON(c, result, nil)
}
