package form

import (
	"github.com/gin-gonic/gin"
	"hex-base/internal/constant"
	"hex-base/internal/core/adapters/logger"
	"hex-base/internal/core/adapters/repo/sql_type/sql"
	"hex-base/internal/core/adapters/validator"
	"hex-base/internal/core/domain"
	"time"
)

type FormAPI interface {
	ApiBuilder
	ApiInfo() string
}

type ApiBuilder interface {
	InjectRepo(adapter sql.SqlAdapter) ApiBuilder
	WithTimeout(ctxTimeout time.Duration) ApiBuilder
	InjectLogger(logger logger.ILogger) ApiBuilder
	InjectValidator(validator validator.ValidatorAdapter) ApiBuilder
	BelongToRouter(router *gin.RouterGroup) ApiBuilder
	ContextPath(contextPath string) ApiBuilder
	Setup() FormAPI
}

type formApi struct {
	ApiBuilder
	contextPath string
	group       *gin.RouterGroup
	repository  domain.FormRepository
	ctxTimeout  time.Duration
	logger      logger.ILogger
	validator   validator.ValidatorAdapter
}

func (api *formApi) InjectRepo(sqlAdapter sql.SqlAdapter) ApiBuilder {
	switch sqlAdapter.GetFramework() {
	case constant.GORM:
		api.repository = sqlAdapter.Gorm().FormStorage()
	case constant.ENT:
		api.repository = sqlAdapter.Ent().FormRepository()
	}

	return api
}

func (api *formApi) WithTimeout(ctxTimeout time.Duration) ApiBuilder {
	api.ctxTimeout = ctxTimeout
	return api
}

func (api *formApi) InjectLogger(logger logger.ILogger) ApiBuilder {
	api.logger = logger
	return api

}

func (api *formApi) ContextPath(contextPath string) ApiBuilder {
	api.contextPath = contextPath
	return api
}

func (api *formApi) InjectValidator(validator validator.ValidatorAdapter) ApiBuilder {
	api.validator = validator
	return api
}

func (api *formApi) BelongToRouter(router *gin.RouterGroup) ApiBuilder {
	api.group = router
	return api
}

func (api *formApi) Setup() FormAPI {
	formRouter := api.group.Group(api.contextPath)
	formRouter.POST("", api.buildCreateFormAction())
	formRouter.GET("/", api.buildListFormByTodoIDAction())
	return api
}

func (api *formApi) ApiInfo() string{
	return "gin-v1"
}
func NewFormAPIBuilder() ApiBuilder {
	return &formApi{}
}
