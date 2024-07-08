package app

import (
	gcontext "context"
	"log/slog"
	"net/http"
	"time"

	"demo/logger"

	"github.com/go-playground/validator/v10"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type context struct {
	*gin.Context
	logger    *slog.Logger
	validator *validator.Validate
}

type Context interface {
	Bind(v any) error
	OK(v any)
	OkWithPagination(v any, page, perPage, totalItems uint)
	CREATED(v any)
	UPDATED()
	BadRequest(err error)
	InternalServerError(err error)
	NotFound(err error)
	JSON(code int, v any)
	Ctx() gcontext.Context
	GetString(key string) string
	Param(key string) string
	Query(key string) string
	GetQuery(key string) (string, bool)
	DefaultQuery(key, defaultValue string) string
}

func NewContext(c *gin.Context, logger *slog.Logger) Context {
	validate := validator.New()
	return &context{Context: c, logger: logger, validator: validate}
}

func (c *context) Bind(v any) error {
	return c.Context.ShouldBindJSON(v)
}

func (c *context) OK(v any) {
	c.Context.JSON(http.StatusOK, Response{
		Status: Success,
		Data:   v,
	})
}

func (c *context) OkWithPagination(v any, page, perPage, totalItems uint) {
	c.Context.JSON(http.StatusOK, Pagination{
		Response{
			Status: Success,
			Data:   v,
		},
		paging(page, perPage, totalItems),
	})
}

func (c *context) CREATED(v any) {
	c.Context.JSON(http.StatusCreated, Response{
		Status: Success,
		Data:   v,
	})
}

func (c *context) UPDATED() {
	c.Context.JSON(http.StatusNoContent, nil)
}

func (c *context) BadRequest(err error) {
	logger.AppErrorf(c.logger.Handler(), err.Error())
	c.Context.JSON(http.StatusBadRequest, Response{
		Status:  Fail,
		Message: err.Error(),
	})
}

func (c *context) InternalServerError(err error) {
	logger.AppErrorf(c.logger.Handler(), err.Error())
	c.Context.JSON(http.StatusInternalServerError, Response{
		Status:  Fail,
		Message: err.Error(),
	})
}

func (c *context) NotFound(err error) {
	logger.AppErrorf(c.logger.Handler(), err.Error())
	c.Context.JSON(http.StatusNotFound, Response{
		Status:  Fail,
		Message: err.Error(),
	})
}

func (c *context) JSON(code int, v any) {
	c.Context.JSON(code, v)
}

func (c *context) Ctx() gcontext.Context {
	return c.Context.Request.Context()
}

func (c *context) GetString(key string) string {
	return c.Context.GetString(key)
}

func (c *context) Param(key string) string {
	return c.Context.Param(key)
}

func (c *context) Query(key string) string {
	return c.Context.Query(key)
}

func (c *context) GetQuery(key string) (string, bool) {
	return c.Context.GetQuery(key)
}

func (c *context) DefaultQuery(key, defaultValue string) string {
	return c.Context.DefaultQuery(key, defaultValue)
}

func NewGinHandler(handler func(Context), logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(NewContext(c, logger.With(slog.String("transaction-id", c.Request.Header.Get("transaction-id")))))
	}
}

type Router struct {
	*gin.Engine
	logger *slog.Logger
}

func NewRouter(logger *slog.Logger) *Router {
	r := gin.Default()

	config := cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"X-Requested-With", "Authorization", "Origin", "Content-Length", "Content-Type", "TransactionID"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}

	r.Use(cors.New(config))

	return &Router{Engine: r, logger: logger}
}

func (r *Router) GET(path string, handler func(Context)) {
	r.Engine.GET(path, NewGinHandler(handler, r.logger))
}

func (r *Router) POST(path string, handler func(Context)) {
	r.Engine.POST(path, NewGinHandler(handler, r.logger))
}

func (r *Router) PUT(path string, handler func(Context)) {
	r.Engine.PUT(path, NewGinHandler(handler, r.logger))
}

func (r *Router) PATCH(path string, handler func(Context)) {
	r.Engine.PATCH(path, NewGinHandler(handler, r.logger))
}

func (r *Router) DELETE(path string, handler func(Context)) {
	r.Engine.DELETE(path, NewGinHandler(handler, r.logger))
}
