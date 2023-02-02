package gin

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"hex-base/internal/appctx"
	"hex-base/internal/core/adapters/logger"
	"hex-base/internal/core/adapters/repo/sql_type/sql"
	"hex-base/internal/core/adapters/validator"
	"hex-base/internal/infrastructure/router/gin/handlers/form"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type ginEngine struct {
	router     *gin.Engine
	log        logger.ILogger
	db         sql.SqlAdapter
	validator  validator.ValidatorAdapter
	ctxTimeout time.Duration
	port       appctx.Port
}

func NewGinServer(
	log logger.ILogger,
	db sql.SqlAdapter,
	validator validator.ValidatorAdapter,
	port appctx.Port,
	t time.Duration,
) *ginEngine {
	return &ginEngine{
		router:     gin.New(),
		log:        log,
		db:         db,
		validator:  validator,
		port:       port,
		ctxTimeout: t,
	}
}

func (g ginEngine) Listen() {
	gin.Recovery()

	g.setAppHandlers(g.router)

	server := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 15 * time.Second,
		Addr:         fmt.Sprintf(":%d", g.port),
		Handler:      g.router,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		g.log.WithFields(logger.Fields{"port": g.port}).Infof("Starting HTTP Server")
		if err := server.ListenAndServe(); err != nil {
			g.log.WithError(err).Fatalln("Error starting HTTP server")
		}
	}()

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := server.Shutdown(ctx); err != nil {
		g.log.WithError(err).Fatalln("Server Shutdown Failed")
	}

	g.log.Infof("Service down")
}

/* TODO ADD MIDDLEWARE */
func (g ginEngine) setAppHandlers(router *gin.Engine) {

	v1:=router.Group("/api/v1")

	 formApi:= form.NewFormAPIBuilder().
		BelongToRouter(v1).
		InjectLogger(g.log).
		InjectRepo(g.db).
		WithTimeout(g.ctxTimeout).
		InjectValidator(g.validator).
		ContextPath("/form").
		Setup()

	 fmt.Println(formApi.ApiInfo())

}
