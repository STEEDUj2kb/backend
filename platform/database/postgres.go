package database

// Reference:
//	- https://entgo.io/docs/sql-integration/#use-pgx-with-postgresql
//	- https://blog.logrocket.com/an-introduction-to-ent/
//	- https://github.com/create-go-app/fiber-go-template/blob/master/platform/database/postgres.go

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/STEEDUj2kb/v1/pkg/utils"
	"github.com/STEEDUj2kb/v1/platform/ent"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
)

var EntClient *ent.Client

// Open new connection
func open(databaseUrl string) (*ent.Client, error) {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Define database connection settings.
	maxConn, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))
	maxIdleConn, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
	maxLifetimeConn, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME_CONNECTIONS"))

	// Set database connection settings:
	// 	- SetMaxOpenConns: the default is 0 (unlimited)
	// 	- SetMaxIdleConns: defaultMaxIdleConns = 2
	// 	- SetConnMaxLifetime: 0, connections are reused forever
	db.SetMaxOpenConns(maxConn)
	db.SetMaxIdleConns(maxIdleConn)
	db.SetConnMaxLifetime(time.Duration(maxLifetimeConn))

	// Try to ping database.
	if err := db.Ping(); err != nil {
		defer db.Close() // close database connection
		return nil, fmt.Errorf("error, not sent ping to database, %w", err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv)), nil
}

func init() {
	// Build PostgreSQL connection URL.
	postgresConnURL, err := utils.ConnectionURLBuilder("postgres")
	if err != nil {
		log.Fatal(err)
	}
	//Open a connection to the database
	Client, err := open(postgresConnURL)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to database successfully")
	defer Client.Close()
	// AutoMigration with ENT
	if err := Client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	EntClient = Client
}
