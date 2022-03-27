package wraperror

import (
	"database/sql"
	"errors"
	"fmt"
	"testing"
)

type QueryError struct {
	err error
	msg string
}

func (q *QueryError) Error() string {
	return q.msg
}

func (q *QueryError) Unwrap() error {
	return q.err
}



func TestGetDBData(t *testing.T) {
	//basicErr := errors.New("test")
	//err := fmt.Errorf("new,%w", basicErr)
	//err1 := fmt.Errorf("new 2,%w",err)
	//err2 := fmt.Errorf("new 3")
	//fmt.Println(err2.Error())
	//if errors.Is(err1,err2) {
	//	fmt.Println("Is success")
	//}

	err3 := new(QueryError)
	err3.msg = "sdfsdf"
	err3.err = &QueryError{
		err: errors.New("qqwewqe"),
		msg: "query",
	}

	e := &QueryError{}
	e.msg = "dfsdfs"
	e.err = sql.ErrNoRows

	e2 := &QueryError{
		err:  errors.New("e2"),
		msg: "e2",
	}

	fmt.Println(err3.msg, err3.err, err3.err.Error())
	fmt.Println(e2.msg, e2.err, e2.err.Error())
	fmt.Println(e2.err.Error())

	e6 := errors.New("aaaaaaa")
	if errors.As(e6, &e2) {
		fmt.Println("as success")
		fmt.Println(err3.msg, err3.err, err3.err.Error())
		fmt.Println(e2.msg, e2.err, e2.err.Error())

		if err3.msg == e.msg {
			fmt.Println("equal")
		}
	}
}

func TestStructEqual(t *testing.T) {
	a := &QueryError{
		err: nil,
		msg: "1",
	}

	b := &QueryError{
		err: nil,
		msg: "1",
	}
	fmt.Printf("%p,%p\n", &a, &b)
	fmt.Printf("%p,%p\n", a, b)
	fmt.Println(a == b)
	//fmt.Println(errors.Is(a, b))
}
