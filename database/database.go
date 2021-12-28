package database

import (
	"database/sql"
	"log"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
) // import()

func ConnectBun() ( db *bun.DB, err error ) {
	dsn := "postgres://library:123@localhost:5432/testlibrary?sslmode=disable"
    // dsn := "unix://user:pass@dbname/var/run/postgresql/.s.PGSQL.5432"
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db = bun.NewDB(sqldb, pgdialect.New())
	if db == nil {
		log.Println( "failed to connect to database " )
		return
	} // if()

	db.SetMaxIdleConns( 2 )
	db.SetMaxOpenConns( 5 )
	db.SetConnMaxLifetime( time.Hour )
	return
} // ConnectBun()