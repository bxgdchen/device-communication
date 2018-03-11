package db

import (
	"database/sql"
)

// DataAccessor ...
type DataAccessor struct {
	gatherDB *sql.DB
}
