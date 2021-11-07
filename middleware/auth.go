/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-11-07 19:17:00
 */
package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func UserAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userInfoJson := session.Get("userInfo")

		if userInfoJson == nil {
			// 取不到就是没有登录
			c.Header("Content-Type", "text/html; charset=utf-8")
			c.String(200, `<script type="text/javascript">top.location.href="/login"</script>`)
			return
		}
		c.Set("userInfo", userInfoJson)
		c.Next()
	}
}
