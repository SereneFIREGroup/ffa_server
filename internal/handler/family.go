package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/serenefiregroup/ffa_server/internal/services/family"
)

func FamilyInfo(c *gin.Context) {
	spanCtx := getCtx(c)

	familyUUID := getFamilyID(c)
	userUUID := GetRequestUserID(c)
	obj, result := family.Info(spanCtx, familyUUID, userUUID)
	RenderJSON(c, result, obj)
}

func SetFIREGold(c *gin.Context) {
	spanCtx := getCtx(c)

	familyUUID := getFamilyID(c)
	userUUID := GetRequestUserID(c)
	var req family.SetFIREGoldRequest
	if err := BindJSON(c, &req); err != nil {
		return
	}
	result := family.SetFIREGold(spanCtx, familyUUID, userUUID, &req)
	RenderJSON(c, result, nil)
}
