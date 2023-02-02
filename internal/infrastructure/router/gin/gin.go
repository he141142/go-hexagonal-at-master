package gin

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"hex-base/internal/core/adapters/logger"
	"hex-base/internal/core/adapters/repo/sql_type/sql"
	"hex-base/internal/core/adapters/validator"
	router2 "hex-base/internal/infrastructure/router"
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
	port       router2.Port
}

func NewGinServer(
	log logger.ILogger,
	db sql.SqlAdapter,
	validator validator.ValidatorAdapter,
	port router2.Port,
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
			g.log.WithError(err).Fatal("Error starting HTTP server")
		}
	}()

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := server.Shutdown(ctx); err != nil {
		g.log.WithError(err).Fatal("Server Shutdown Failed")
	}

	g.log.Infof("Service down")
}

/* TODO ADD MIDDLEWARE */
func (g ginEngine) setAppHandlers(router *gin.Engine) {

	router.Group("/api/v1")

	//router.POST("/v1/transfers", g.buildCreateTransferAction())
	//router.GET("/v1/transfers", g.buildFindAllTransferAction())
	//
	//router.GET("/v1/accounts/:account_id/balance", g.buildFindBalanceAccountAction())
	//router.POST("/v1/accounts", g.buildCreateAccountAction())
	//router.GET("/v1/accounts", g.buildFindAllAccountAction())
	//
	//router.GET("/v1/health", g.healthcheck())
}
