package ent

import (
	"context"
	"entgo.io/ent/dialect"
	sql "database/sql"
	"fmt"
	entsql "entgo.io/ent/dialect/sql"

	"hex-base/internal/appctx"
	"hex-base/internal/constant"
	"hex-base/internal/core/adapters/repo/sql_type/sql/ent/lib/ent"
	"hex-base/internal/core/adapters/repo/sql_type/sql/ent/repo/form"
	"hex-base/internal/core/adapters/repo/sql_type/sql/ent/repo/todo"
	"hex-base/internal/infrastructure/database"
	"log"
	"time"
)

type entProvider struct {
	client *ent.Client
	formStore form.FormStorage
	todoStore todo.TodoStorage
}

func (provider *entProvider) Engine() *ent.Client {
	return provider.client
}

func NewEntProvider(appCtx appctx.AppContext, config database.DatabaseConfig) *entProvider {
	client, err :=  Connection(appCtx,config)
	if err != nil{
		panic(any("error when initiate database"))
	}
	return &entProvider{
		client:    client,
		formStore: form.NewFormStorage(client),
		todoStore: todo.NewTodoStorage(client),
	}
}


func Connection(appCtx appctx.AppContext, config database.DatabaseConfig) (*ent.Client, error) {

	//migration(connectionStr)
	dbDriver := constant.ConvertDriver(config.Driver())
	var connectionStr = ""
	switch dbDriver {
	case constant.POSTGRES:
	case constant.MYSQL:
	default:
		connectionStr = getPostgresConnectionString(config)
	}

	dsn := connectionStr
	db, err := sql.Open(dialect.Postgres, dsn)

	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	drv := entsql.OpenDB(dialect.Postgres, db)

	client := ent.NewClient(ent.Driver(drv), ent.Debug(), ent.Log(func(i ...interface{}) {
		for _, v := range i {
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"), v)
			fmt.Print("\n")
		}
	}))

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	if err != nil {
		log.Fatalf("failed: %v", err)
	}
	// Run the auto migration tool.

	return ent.NewClient(ent.Driver(drv), ent.Debug(), ent.Log(func(i ...interface{}) {
		for _, v := range i {
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"), v)
			fmt.Print("\n")
		}
	})), err
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

// get the core
func ConnectDb() any {
	return any(1)
}

