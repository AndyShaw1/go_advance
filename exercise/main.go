package main

import (
	"go_advance/exercise/cocurrent"
	"go_advance/exercise/redis"
)

func main() {
	//db.ConnectDB()
	//wraperror.GetDBData()
	cocurrent.ServerStart()
	redis.Compare()
}
