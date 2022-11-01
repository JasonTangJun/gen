package main

import (
	"context"
	"fmt"

	"JasonTangJun/gen/examples/biz"
	"JasonTangJun/gen/examples/conf"
	"JasonTangJun/gen/examples/dal"
	"JasonTangJun/gen/examples/dal/query"
)

func init() {
	dal.DB = dal.ConnectDB(conf.MySQLDSN).Debug()
}

func main() {
	// start your project here
	fmt.Println("hello world")
	defer fmt.Println("bye~")

	query.SetDefault(dal.DB)
	biz.Query(context.Background())
}
