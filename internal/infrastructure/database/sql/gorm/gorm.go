package gorm

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/gsabadini/go-bank-transfer/infrastructure/log"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gorm_log "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"hex-base/internal/appctx"
	"hex-base/internal/constant"
	"hex-base/internal/core/adapters/logger"
	"hex-base/internal/core/adapters/repo/sql_type/sql/gorm/repo/form"
	"hex-base/internal/core/adapters/repo/sql_type/sql/gorm/repo/todo"
	"hex-base/internal/infrastructure/database"
)

//type Adapter interface {
//	Engine() *gorm.DB
//}

type gormProvider struct {
	//Adapter
	db *gorm.DB
	formStore form.FormStorage
	todoStore todo.TodoStorage
}

func NewGormProvider(appCtx appctx.AppContext, config database.DatabaseConfig) *gormProvider {
	return &gormProvider{
		db: Connection(appCtx, config),
	}
}

func migration(connectionStr string, log logger.ILogger) {
	m, err := migrate.New(
		"file://./migrations",
		connectionStr)

	if err != nil {
		fmt.Println("migrate err: ", err)
		log.WithFields(logger.Fields(logrus.Fields{
			"migrate-issue": err.Error(),
		}))

		panic(any(1))
	}

	if err := m.Up(); err != nil {
		fmt.Println("migrate up err: ", err)
		log.WithFields(logger.Fields{
			"migrate-up-issue": err.Error(),
		})
	}

}

func Connection(appCtx appctx.AppContext, config database.DatabaseConfig) *gorm.DB {
	dbDriver := constant.ConvertDriver(config.Driver())
	var connectionStr = ""
	switch dbDriver {
	case constant.POSTGRES:
	case constant.MYSQL:
	default:
		connectionStr = getPostgresConnectionString(config)
	}

	// fmt.Printf("connection string: %s", connectionStr)
	log := appCtx.Logger()

	migration(connectionStr, log)

	dsn := connectionStr
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: gorm_log.Default.LogMode(gorm_log.Info),
	})

	if err != nil {
		log.WithFields(logger.Fields{
			"database":   "form-service postgres database",
			"connection": "disconnected",
			"issue":      "connection issue",
			"message":    err.Error(),
		}).Fatalln("form-service postgres db issue")
	} else {
		log.WithFields(logger.Fields{
			"database":   "form-service postgres database",
			"connection": "connected",
		}).Infof("form-service postgres db connected")
	}
	return db
}

func getPostgresConnectionString(config database.DatabaseConfig) string {
	dbHostEnv := config.Host()
	dbPortEnv := config.Port()
	dbUserEnv := config.User()
	dbPasswordEnv := config.Password()
	dbName := config.DBName()
	dbSSLMode := config.SslMode()

	//connectionStr := "postgresql://sykros:fqQ3nN4L@localhost:9001/zlp-demo?sslmode=disable"
	connectionStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s%s", dbUserEnv, dbPasswordEnv, dbHostEnv, dbPortEnv, dbName,
		func() string {
			if dbSSLMode == "" {
				return dbSSLMode
			}
			return fmt.Sprintf("?sslmode=%s", dbSSLMode)
		}())

	return connectionStr
}

func getMysqlConnectionString(config database.DatabaseConfig) string {
	return "Not implemented"
}
