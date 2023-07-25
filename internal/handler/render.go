package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/serenefiregroup/ffa_server/pkg/errors"
	"github.com/serenefiregroup/ffa_server/pkg/log"
)

func RenderJSONAndStop(c *gin.Context, result error, obj interface{}) {
	errP := buildErrPayloadAndLog(c, result, true)
	if errP.HttpStatus == http.StatusOK {
		if obj == nil {
			c.JSON(errP.HttpStatus, errP)
		} else {
			c.JSON(errP.HttpStatus, obj)
		}
	} else {
		c.JSON(errP.HttpStatus, errP)
		_ = c.AbortWithError(errP.HttpStatus, result)
	}
	c.Next()
}

func RenderJSON(c *gin.Context, result error, obj interface{}) {
	errP := buildErrPayloadAndLog(c, result, true)
	if errP.HttpStatus == http.StatusOK {
		if obj == nil {
			c.JSON(errP.HttpStatus, errP)
		} else {
			c.JSON(errP.HttpStatus, obj)
		}
	} else {
		c.JSON(errP.HttpStatus, errP)
	}
	c.Next()
}

func buildErrPayloadAndLog(c *gin.Context, err error, shouldLog bool) (errp *errors.ErrPayload) {
	errp = errors.NewErrPayload(err)
	if shouldLog {
		// 根据状态码打印日志
		if errp.HttpStatus < 400 {
			// 不需要打印日志
		} else if errp.HttpStatus >= 500 && errp.HttpStatus < 600 {
			// 服务端错误
			log.Error("", err)
			// 对客户端隐藏详细信息
			errp.Code = errors.ServerError
			errp.HttpStatus = http.StatusInternalServerError
			errp.Desc = ""
			errp.Values = nil
		} else {
			// 客户端错误 & 自定义错误，Warn
			log.Warn("", err)
		}
	}
	return
}

func BindJSON(c *gin.Context, obj interface{}) (err error) {
	defer func() {
		if p := recover(); p != nil {
			switch p := p.(type) {
			case error:
				err = p
			default:
				err = fmt.Errorf("%s", p)
			}
			renderBindJSONError(c, err)
		}
	}()

	if err = c.BindJSON(obj); err != nil {
		renderBindJSONError(c, err)
		return
	}

	return
}

func renderBindJSONError(c *gin.Context, err error) {
	code := errors.Code(errors.Malformed, errors.JSON)
	err = errors.Wrapf(err, code, "error at parsing json: "+err.Error())
	RenderJSON(c, err, nil)
}
