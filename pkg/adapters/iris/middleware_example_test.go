package iris

import (
	"github.com/kataras/iris"
)

func Example() {
	r := iris.New()
	r.Use(
		SentinelMiddleware(
			// customize resource extractor if required
			// method_path by default
			WithResourceExtractor(func(ctx *iris.Context) string {
				return ctx.GetHeader("X-Real-IP")
			}),
			// customize block fallback if required
			// abort with status 429 by default
			WithBlockFallback(func(ctx *iris.Context) {
				ctx.AbortWithStatusJSON(400, map[string]interface{}{
					"err":  "too many request; the quota used up",
					"code": 10222,
				})
			}),
		),
	)

	r.GET("/test", func(c *iris.Context) {})
	_ = r.Run(":0")
}
