package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/serenefiregroup/ffa_server/internal/services/family"
)

func SetFIREGold(c *gin.Context) {
	spanCtx := getCtx(c)

	familyUUID := getFamilyULID(c)
	userUUID := GetRequestUserID(c)
	var req family.SetFIREGoldRequest
	if err := BindJSON(c, &req); err != nil {
		return
	}
	result := family.SetFIREGold(spanCtx, familyUUID, userUUID, &req)
	RenderJSON(c, result, nil)
}
