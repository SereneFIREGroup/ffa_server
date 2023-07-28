package handler

import (
	"context"
	"net/http"
	"net/http/httputil"
	"runtime/debug"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/serenefiregroup/ffa_server/internal/model/constants"
	userModel "github.com/serenefiregroup/ffa_server/internal/model/user"
	"github.com/serenefiregroup/ffa_server/pkg/db"
	"github.com/serenefiregroup/ffa_server/pkg/errors"
	jaegerUtils "github.com/serenefiregroup/ffa_server/pkg/jaeger"
	"github.com/serenefiregroup/ffa_server/pkg/log"
	"go.uber.org/ratelimit"
)

const (
	HeaderUserID = "X-User-ID"
	HeaderToken  = "X-Token"
)

var PanicHandler = gin.HandlerFunc(func(ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			httpRequest, _ := httputil.DumpRequest(ctx.Request, false)
			log.Error("%s\n%s\n%s\n%s\n", err, ctx.Request.RemoteAddr, string(httpRequest), debug.Stack())
			ctx.AbortWithStatus(http.StatusInternalServerError)
		}
	}()
	ctx.Next()
})

var Jaeger = gin.HandlerFunc(func(c *gin.Context) {
	traceId := c.GetHeader("uber-trace-id")
	var span opentracing.Span
	if traceId != "" {
		var err error
		span, err = jaegerUtils.GetParentSpan(c.FullPath(), traceId, c.Request.Header)
		if err != nil {
			return
		}
	} else {
		span = jaegerUtils.StartSpan(opentracing.GlobalTracer(), c.FullPath())
	}
	defer span.Finish()

	c.Set(jaegerUtils.SpanCTX, opentracing.ContextWithSpan(c, span))
	c.Next()
})

func LeakBucket() gin.HandlerFunc {
	limiter := ratelimit.New(1)
	return func(c *gin.Context) {
		if time.Now().Sub(limiter.Take()) > 0 {
			c.AbortWithStatus(http.StatusTooManyRequests)
		}
		c.Next()
	}
}

func ParseHeadOrCookie(c *gin.Context, k string) string {
	v := c.GetHeader(k)
	if len(v) == 0 {
		v, _ = c.Cookie(k)
	}
	return v
}

func CheckLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := getCtx(c)

		userID := GetRequestUserID(c)
		token := GetRequestToken(c)
		if len(userID) != constants.ULIDLen {
			RenderJSONAndStop(c, errors.InvalidParameterError(errors.User, errors.ID, errors.InvalidParameter), nil)
			return
		}
		if len(token) != constants.TokenLen {
			RenderJSONAndStop(c, errors.InvalidParameterError(errors.User, errors.Token, errors.InvalidParameter), nil)
			return
		}
		s, err := userModel.GetSession(ctx, db.DB, userID, token)
		if err != nil {
			RenderJSONAndStop(c, errors.Trace(err), nil)
			return
		}
		if s == nil {
			RenderJSONAndStop(c, errors.AuthFailureError(errors.InvalidToken), nil)
			return
		}
		c.Next()
	}
}

func CheckFamily() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := getCtx(c)

		familyID := getFamilyID(c)
		userID := GetRequestUserID(c)
		if len(userID) != constants.ULIDLen {
			RenderJSONAndStop(c, errors.InvalidParameterError(errors.User, errors.ID, errors.InvalidParameter), nil)
			return
		}
		if len(familyID) != constants.ULIDLen {
			RenderJSONAndStop(c, errors.InvalidParameterError(errors.Family, errors.ID, errors.InvalidParameter), nil)
			return
		}
		u, err := userModel.GetUserByFamilyAndID(ctx, db.DB, familyID, userID)
		if err != nil {
			RenderJSONAndStop(c, errors.Trace(err), nil)
			return
		}
		if u == nil {
			RenderJSONAndStop(c, errors.NotFoundError(errors.User), nil)
			return
		}
		c.Next()
	}
}

func getCtx(c *gin.Context) context.Context {
	spanCtxInterface, _ := c.Get(jaegerUtils.SpanCTX)
	var spanCtx context.Context
	spanCtx = spanCtxInterface.(context.Context)
	return spanCtx
}

func getFamilyID(c *gin.Context) string {
	familyID := c.Param("familyID")
	return familyID
}

func getUserID(c *gin.Context) string {
	userID := c.Param("userID")
	return userID
}

func GetRequestUserID(c *gin.Context) string {
	return ParseHeadOrCookie(c, HeaderUserID)
}

func GetRequestToken(c *gin.Context) string {
	return ParseHeadOrCookie(c, HeaderToken)
}
