package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/wanhello/omgind/internal/app/contextx"
	"github.com/wanhello/omgind/internal/app/ginx"
	"github.com/wanhello/omgind/pkg/auth"
	"github.com/wanhello/omgind/pkg/errors"
	"github.com/wanhello/omgind/pkg/global"
	"github.com/wanhello/omgind/pkg/logger"
)

func wrapUserAuthContext(c *gin.Context, userID string) {
	ginx.SetUserID(c, userID)
	ctx := contextx.NewUserID(c.Request.Context(), userID)
	ctx = logger.NewUserIDContext(ctx, userID)
	c.Request = c.Request.WithContext(ctx)
}

// UserAuthMiddleware 用户授权中间件
func UserAuthMiddleware(a auth.Auther, skippers ...SkipperFunc) gin.HandlerFunc {
	if !global.C.JWTAuth.Enable {
		return func(c *gin.Context) {
			wrapUserAuthContext(c, global.C.Root.UserName)
			c.Next()
		}
	}

	return func(c *gin.Context) {
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}

		userID, err := a.ParseUserID(c.Request.Context(), ginx.GetToken(c))
		if err != nil {
			if err == auth.ErrInvalidToken {
				if global.C.IsDebugMode() {
					wrapUserAuthContext(c, global.C.Root.UserName)
					c.Next()
					return
				}
				ginx.ResError(c, errors.ErrInvalidToken)
				return
			}
			ginx.ResError(c, errors.WithStack(err))
			return
		}

		wrapUserAuthContext(c, userID)
		c.Next()
	}
}
