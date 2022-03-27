package wraperror

import (
	"fmt"
	"go_advance/service/taskerror"
	"go_advance/util"
)

func GetDBData(){
	err := taskerror.GetData()
	if err == nil {
		return
	}

	fmt.Println(util.IsSqlNoRowsError(err))
	fmt.Printf("stack:%+v",err)
}