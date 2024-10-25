package main

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gen"
	"gorm.io/gorm"
)

// generate code
func main() {

	// specify the output directory (default: "./query")
	// ### if you want to query without context constrain, set mode gen.WithoutContext ###
	g := gen.NewGenerator(gen.Config{
		OutPath: "./test/dao",
		Mode:    gen.WithoutContext,
		// if you want the nullable field generation property to be pointer type, set FieldNullable true
		// FieldNullable: true,
		// if you want to generate index tags from database, set FieldWithIndexTag true
		FieldWithIndexTag: true,
		// if you want to generate type tags from database, set FieldWithTypeTag true
		FieldWithTypeTag: true,
		// if you need unit tests for query code, set WithUnitTest true
		// WithUnitTest: true,
	})
	// Specify the data type mapping for SQLite
	dataMap := map[string]func(columnType gorm.ColumnType) (dataType string){
		"INTEGER": func(columnType gorm.ColumnType) (dataType string) { return "int" },
		"TEXT":    func(columnType gorm.ColumnType) (dataType string) { return "string" },
		"BOOLEAN": func(columnType gorm.ColumnType) (dataType string) { return "bool" },
		// Add other SQLite data type mappings as needed
	}
	g.WithDataTypeMap(dataMap)

	// SQLite connection string (use in-memory database for demonstration)
	dbUrl := "/Users/Lee/Library/Application Support/XiuXianHelpTool/XiaLou.db"
	db, err := gorm.Open(sqlite.Open(dbUrl), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	g.UseDB(db)

	// apply basic crud api on structs or table models which is specified by table name with function
	// GenerateModel/GenerateModelAs. And generator will generate table models' code when calling Execute.
	g.ApplyBasic(
		g.GenerateAllTable()...,
	)

	// apply diy interfaces on structs or table models
	g.ApplyInterface(func() {})

	// execute the action of code generation
	g.Execute()
}
