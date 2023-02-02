package form

import (
	"github.com/gin-gonic/gin"
	create3 "hex-base/internal/core/adapters/api/action/form/create"
	action_get "hex-base/internal/core/adapters/api/action/form/get"
	create2 "hex-base/internal/core/adapters/presenter/form/create"
	get_presenter "hex-base/internal/core/adapters/presenter/form/get"
	"hex-base/internal/core/usecase/form/create"
	"hex-base/internal/core/usecase/form/get"
)

func (api *formApi) buildCreateFormAction()  gin.HandlerFunc{
	return func(c *gin.Context) {
		var (
			usecase = create.NewCreateFormInteractor(
				api.repository,
				create2.NewCreateFormPresenter(),
				api.ctxTimeout,
				)
			action = create3.NewCreateFormAction(usecase, api.logger, api.validator)
			)
		action.Execute(c.Writer, c.Request)
	}
}


func (api *formApi) buildListFormByTodoIDAction() gin.HandlerFunc{
	return func(c *gin.Context) {
		var (
			usecase = get.NewListFormByTodoIDInteractor(
				api.repository,
				get_presenter.NewListFormByTodoIdPresenter(),
				api.ctxTimeout,
			)
			action = action_get.NewListFormByTodoIdAction(usecase, api.logger)
		)
		action.Execute(c.Writer, c.Request)
	}
}