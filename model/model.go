package model

import (
	"context"

	"github.com/uptrace/bun"
) // import()

var (
	Db *bun.DB
	Ctx context.Context
) // var()

func Init( _db *bun.DB ) {
	Db = _db
	Ctx = context.Background()
} // Init()