package form

import (
	"github.com/gin-gonic/gin"
	"hex-base/internal/core/adapters/logger"
	"hex-base/internal/core/adapters/validator"
	"hex-base/internal/core/domain"
	"time"
)

type ApiBuilder interface {
	InjectRepo(repository domain.FormRepository) ApiBuilder
	WithTimeout(ctxTimeout time.Duration) ApiBuilder
	InjectLogger(logger logger.ILogger) ApiBuilder
	InjectValidator(validator validator.ValidatorAdapter) ApiBuilder
	BelongToRouter(router *gin.RouterGroup) ApiBuilder
	ContextPath(contextPath string) ApiBuilder
	Setup()
}

type formApi struct {
	ApiBuilder
	contextPath string
	group      *gin.RouterGroup
	repository domain.FormRepository
	ctxTimeout time.Duration
	logger     logger.ILogger
	validator  validator.ValidatorAdapter
}

func (api *formApi) InjectRepo(repository domain.FormRepository) ApiBuilder{
	api.repository = repository
	return api
}

func (api *formApi) WithTimeout(ctxTimeout time.Duration) ApiBuilder{
	api.ctxTimeout = ctxTimeout
	return api
}

func (api *formApi) InjectLogger(logger logger.ILogger) ApiBuilder{
	api.logger = logger
	return api

}

func (api *formApi)ContextPath(contextPath string) ApiBuilder{
	api.contextPath = contextPath
	return api
}

func (api *formApi) InjectValidator(validator validator.ValidatorAdapter) ApiBuilder{
	api.validator = validator
	return api
}

func (api *formApi) Setup(){
	formRouter := api.group.Group(api.contextPath)
	formRouter.POST("",api.buildCreateFormAction())
	formRouter.GET("/",api.buildCreateFormAction())
}

func NewFormAPIBuilder(router *gin.RouterGroup,
) ApiBuilder {
	return &formApi{}
}
