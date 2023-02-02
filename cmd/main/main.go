package main

import (
	"hex-base/internal/constant"
	"hex-base/internal/infrastructure"
	"hex-base/internal/infrastructure/logger"
	"hex-base/internal/infrastructure/router"
	"hex-base/internal/infrastructure/validator"
	"os"
	"time"
)

func main() {
	var app = infrastructure.NewConfig().
		Name(os.Getenv("APP_NAME")).
		ContextTimeout(10 * time.Second).
		Logger(logger.ZapType).
		Validator(validator.PLAYGROUND_INSTANCE).
		DbSQL(constant.ENT).
		DbNoSql()

	app.WebServerPort(os.Getenv("APP_PORT")).
		WebServer(router.InstanceGin).
		Start()
}