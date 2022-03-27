package util

import (
	"database/sql"
	"github.com/pkg/errors"
)

func IsSqlNoRowsError(err error) bool {
	return errors.Is(err, sql.ErrNoRows)
}
