package server

import (
	"context"
	"fmt"
	"net/http"

	"delta.nitt.edu/dion/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Params struct {
	fx.In
	Config *config.Config
}

type RouterParams struct {
	Params
	R  *gin.Engine
	Db *gorm.DB
}

var db = make(map[string]string)

func newLogger(p Params) (*zap.Logger, error) {
	if p.Config.IsProd {
		return zap.NewProduction()
	} else {
		return zap.NewDevelopment()
	}
}

func SetupRouter(p RouterParams) {
	if p.Config.IsProd {
		gin.SetMode(gin.ReleaseMode)
	}
	fmt.Println(p.Db)

	// Ping test
	p.R.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get user value
	p.R.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := p.R.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})
}

func NewServer(lc fx.Lifecycle, g *gin.Engine, logger *zap.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Sugar().Infow("Starting Server")
			return g.Run(":8000")
		},
		OnStop: func(ctx context.Context) error {
			logger.Sugar().Infow("Stopping Server")
			return nil
		},
	})
}

var Module = fx.Options(
	fx.Provide(
		gin.Default,
		newLogger,
	),
)
