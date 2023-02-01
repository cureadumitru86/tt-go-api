package handlers

import (
	"strings"

	"github.com/cureadumitru86/tt-go-api/src/data"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

/*
	ServiceContext - Custom context object that includes additional session infomation

To use, add cc := c.(*ServiceContext) to type cast correctly
*/
type ServiceContext struct {
	echo.Context
}

// DB - Creates database object for the context of the handler handling the request.
func (c *ServiceContext) DB() (*gorm.DB, error) {
	dsn := viper.GetString("database_url")
	db, err := data.NewDB(dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// InitServer - Initializes a server instance with config.
func InitServer() (*echo.Echo, error) {
	e := echo.New()
	e.HideBanner = true
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &ServiceContext{c}
			return next(cc)
		}
	})
	e.Use(middleware.Recover())
	logger, _ := zap.NewProduction()
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info("request",
				zap.String("URI", v.URI),
				zap.Int("status", v.Status),
			)

			return nil
		},
	}))
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Skipper: func(c echo.Context) bool {
			if strings.Contains(c.Request().URL.Path, "swagger") {
				return true
			}
			return false
		},
	}))
	e.Use(middleware.RequestID())
	e.Use(middleware.Secure())
	// e.Use(middleware.CSRF())
	e.Use(middleware.CORS())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(100)))
	registerRoutes(e)
	return e, nil
}
