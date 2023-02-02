package main

import (
	"hex-base/internal/appctx"
	"hex-base/internal/constant"
	"hex-base/internal/infrastructure"
	"hex-base/internal/infrastructure/logger"
	"hex-base/internal/infrastructure/router"
	"hex-base/internal/infrastructure/validator"
	"os"
	"time"
)

func main() {
	baseLog,_ := logger.NewZapLogger()
	var app = infrastructure.NewConfig(appctx.NewAppContext(baseLog)).
		Name(os.Getenv("APP_NAME")).
		ContextTimeout(10 * time.Second).
		Logger(logger.ZapType).
		Validator(validator.PLAYGROUND_INSTANCE).
		DbSQL(constant.ENT)
		//DbNoSql()

	app.WebServerPort("8009").
		WebServer(router.InstanceGin).
		Start()
}