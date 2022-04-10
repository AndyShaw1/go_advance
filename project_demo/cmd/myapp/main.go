package main

import (
	_ "github.com/go-sql-driver/mysql"
	"go_advance/exercise/cocurrent"
)

func main() {
	cocurrent.ServerStart()
}
