package database

import (
	"context"
	"database/sql"
)

// Db object
var Db *sql.DB

// Ctx object
var Ctx context.Context