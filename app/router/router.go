package router

import (
	"lapakUmkm/app/middlewares"
	"lapakUmkm/utils/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	middlewares.Cors(e)
	middlewares.BasicLogger(e)
	e.GET("/", index)
	AuthRouter(db, e)
	UserRouter(db, e)
	CategoryRouter(db, e)
	ProductRouter(db, e)
	ProductImagesRouter(db, e)
	CartRouter(db, e)
	FeedbackRouter(db, e)
	DiscussionRouter(db, e)
}

func index(c echo.Context) error {
	var data = map[string]interface{}{
		"message":       "Welcome to Starter Pack Echo",
		"documentation": "/swagger/index.html",
	}

	return c.JSON(http.StatusOK, helpers.ResponseSuccess("-", data))
}
