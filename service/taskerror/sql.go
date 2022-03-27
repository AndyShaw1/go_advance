package taskerror

import (
	"database/sql"
	"github.com/pkg/errors"
	"go_advance/store/db"
)

type sqlStruct struct {
}

func GetData() error {
	var err error
	query := "select * from tabel_1 limit 1"
	tmp := new(sqlStruct)
	err = db.MainDB.Unsafe().Get(tmp, query)
	err = sql.ErrNoRows
	err = errors.Wrap(err, "get mysql error")
	//err = errors.WithMessage(err,"mysql 2")
	//err = errors.WithMessage(err,"mysql 3")
	return err
}
