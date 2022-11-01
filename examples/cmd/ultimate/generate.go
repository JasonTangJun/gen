package main

import (
	"github.com/JasonTangJun/gen"
	"github.com/JasonTangJun/gen/examples/conf"
	"github.com/JasonTangJun/gen/examples/dal"
	"github.com/JasonTangJun/gen/examples/dal/model"
)

func init() {
	dal.DB = dal.ConnectDB(conf.MySQLDSN).Debug()

	prepare(dal.DB) // prepare table for generate
}

var dataMap = map[string]func(detailType string) (dataType string){
	"int":  func(detailType string) (dataType string) { return "int64" },
	"json": func(string) string { return "json.RawMessage" },
}

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "../../dal/query",
		Mode:    gen.WithDefaultQuery,

		WithUnitTest: true,

		FieldNullable:     true,
		FieldCoverable:    true,
		FieldWithIndexTag: true,
	})

	g.UseDB(dal.DB)

	g.WithDataTypeMap(dataMap)
	g.WithJSONTagNameStrategy(func(c string) string { return "-" })

	g.ApplyBasic(model.Customer{})
	g.ApplyBasic(g.GenerateAllTable()...)

	g.Execute()
}
