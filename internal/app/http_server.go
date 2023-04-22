package app

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	v1_activityGroup "github.com/kholiqcode/go-todolist/internal/activityGroup/delivery/http/v1"
	v1_todo "github.com/kholiqcode/go-todolist/internal/todo/delivery/http/v1"
	"github.com/kholiqcode/go-todolist/pkg/logger"
	"github.com/kholiqcode/go-todolist/pkg/middleware"
	"github.com/kholiqcode/go-todolist/utils"
)

type HttpServer interface {
	ListenAndServe()
}

type httpServerImpl struct {
	route           *chi.Mux
	config          *utils.BaseConfig
	activityHandler v1_activityGroup.ActivityGroupHandler
	todoHandler     v1_todo.TodoHandler
	startAt         time.Time
}

func (s *httpServerImpl) ListenAndServe() {

	middleware.SetupMiddleware(s.route, s.config)

	s.activityHandler.MapRoutes()
	s.todoHandler.MapRoutes()

	s.runHealthCheck()

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", s.config.ServerPort),
		Handler: s.route,
	}

	go func() {
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			logger.LogFatal(fmt.Sprintf("HTTP server error: %v", err))
		}

		log.Println("Stopped serving new connections.")
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()

	err := server.Shutdown(shutdownCtx)
	utils.LogIfError(err)

	logger.LogInfo("HTTP server stopped.")

}
