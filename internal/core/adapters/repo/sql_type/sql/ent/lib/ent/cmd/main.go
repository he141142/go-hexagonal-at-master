//package main
//
//import (
//	"entgo.io/ent/entc"
//	"entgo.io/ent/entc/gen"
//	"log"
//)

//import (
//	"ariga.io/atlas/sql/sqltool"
//	"context"
//	"entgo.io/ent/dialect"
//	entSql "entgo.io/ent/dialect/sql"
//	"entgo.io/ent/dialect/sql/schema"
//	_ "github.com/lib/pq"
//	"gitlab.com/varadise-ltd/backend/iot-hub/goiot/internal/config"
//	"gitlab.com/varadise-ltd/backend/iot-hub/goiot/internal/ent"
//	"gitlab.com/varadise-ltd/backend/iot-hub/goiot/internal/utils"
//	schema_helper "hex-base/internal/core/utils"
//	"log"
//	"os"
//	"path/filepath"
//)
//
//func main() {
//	if len(os.Args) != 2 {
//		log.Fatalf("Must provide migration name!")
//		return
//	}
//	_ = Migrate(os.Args[1],schema_helper.GetRoot())
//}
//
//func OpenPostgresConnection(connectionString string) *entSql.Driver {
//	driver, err := entSql.Open(dialect.Postgres, connectionString)
//	if err != nil {
//		log.Fatalf("failed to connect to the database. %s", err)
//		return nil
//	}
//	return driver
//}
//
//func Migrate(migrationName string,targetDir string)error{
//	ctx := context.Background()
//	// Create a local migration directory able to understand golang-migrate.go migration file format for replay.
//	dir, err := sqltool.NewGolangMigrateDir(filepath.Join(utils.GetRoot(),targetDir))
//
//	driver := OpenPostgresConnection(config.NewDBConfig().
//		GetConnectionString())
//
//	versionedClient := ent.NewClient(ent.Driver(driver))
//
//	if err != nil {
//		log.Fatalf("failed creating atlas migration directory: %v", err)
//	}
//
//	// Migrate diff options.
//	opts := []schema.MigrateOption{
//		schema.WithDir(dir),                         // provide migration directory
//		schema.WithMigrationMode(schema.ModeReplay), // provide migration mode
//		//schema.WithFormatter(formatter),
//		schema.WithDropColumn(true),
//		schema.WithDropIndex(true),
//	}
//	// Generate migrations using Atlas support for Postgres (note the Ent dialect option passed above).
//
//	err = versionedClient.Schema.NamedDiff(ctx, migrationName, opts...)
//
//	if err != nil {
//		log.Fatalf("failed generating migration file: %v", err)
//		return err
//	}else {
//		return nil
//	}
//}

package main

import (
	"fmt"
	schema_helper "hex-base/internal/core/utils"
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	if err := entc.Generate(fmt.Sprintf("%s/core/adapters/repo/sql_type/sql/ent/lib/ent/schema",schema_helper.GetRoot()), &gen.Config{}); err != nil {
		log.Fatal("running ent codegen:", err)
	}
}

