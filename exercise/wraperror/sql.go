package wraperror

import (
	"fmt"
	"go_advance/service/taskerror"
	"go_advance/util"
)

//我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
//应该抛出wrap后的error，这样会携带stack信息，方法定位问题，封装一个util方法，利用方法来替代err == sql.ErrNoRows的判断即可。

func GetDBData(){
	err := taskerror.GetData()
	if err == nil {
		return
	}

	fmt.Println(util.IsSqlNoRowsError(err))
	fmt.Printf("stack:%+v",err)
}