package main

import (
	"context"
	"database/sql"
	"demo/app"
	"demo/config"
	"demo/database"
	"demo/logger"
	"demo/skill"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg := config.C(os.Getenv("ENV"))
	pg, cleanup := database.NewPostgres(cfg.Database)
	logs := logger.New()
	r := NewRouter(logs, cfg, pg)

	srv := http.Server{
		Addr:              cfg.Server.Host + ":" + cfg.Server.Port,
		Handler:           r,
		ReadHeaderTimeout: 5 * time.Second,
	}
	fmt.Println("Server is running on " + cfg.Server.Host + ":" + cfg.Server.Port)

	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM)
		<-sigint

		cleanup()
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			logs.Info("HTTP server Shutdown: " + err.Error())
		}
		close(idleConnsClosed)
	}()

	if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		log.Panic("HTTP server ListenAndServe: " + err.Error())
	}

	<-idleConnsClosed
}

func NewRouter(logs *slog.Logger, cfg config.Config, pg *sql.DB) *app.Router {
	r := app.NewRouter(logs)
	r.GET("/health", func(c app.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := pg.PingContext(ctx); err != nil {
			c.InternalServerError(fmt.Errorf("can not connect to database: %w", err))
			return
		}
		c.OK("demo api is ready and connected to database")
	})

	// TODO: add basic auth
	// r.Use(auth.Basic(""))

	{
		store := skill.NewStorage(pg)
		h := skill.NewHandler(store)
		r.GET("/skills", h.GetSkills)
		r.GET("/skills/:key", h.GetSkillByKey)
		r.PUT("/skills/:key", h.UpdateSkill)
		r.POST("/skills", h.CreateSkill)

		// actions
		r.PATCH("/skills/:key/actions/name", h.PatchName)
		r.PATCH("/skills/:key/actions/description", h.PatchDescription)
		r.PATCH("/skills/:key/actions/logo", h.PatchLogo)
		r.PATCH("/skills/:key/actions/levels", h.PatchLevels)
		r.PATCH("/skills/:key/actions/tags", h.PatchTags)
	}

	return r
}
