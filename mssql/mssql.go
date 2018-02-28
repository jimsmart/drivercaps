package mssql

import (
	"github.com/jimsmart/drivercaps"
)

var (
	User string = "sa"
	// Pass string
	Pass string = "Password123"
	Host string = "localhost"
	// Port string = "1433"
	Port string = "32789"
)

func NewTest(pkgURL, pkgName, drvName, connStr string) *drivercaps.DriverTest {
	t := &drivercaps.DriverTest{
		Database: "Microsoft SQL Server",
		PkgURL:   pkgURL,
		PkgName:  pkgName,
		DrvName:  drvName,
		ConnStr:  connStr,
		Columns: []*drivercaps.ColumnDefn{
			// https://docs.microsoft.com/en-us/sql/t-sql/data-types/data-types-transact-sql
			// Exact numerics
			&drivercaps.ColumnDefn{DDL: "bigint"},
			&drivercaps.ColumnDefn{DDL: "bit"},
			&drivercaps.ColumnDefn{DDL: "decimal"},
			&drivercaps.ColumnDefn{DDL: "int"},
			// &drivercaps.ColumnDefn{DDL: "money"}, // TODO(js) Will panic.
			&drivercaps.ColumnDefn{DDL: "numeric"},
			&drivercaps.ColumnDefn{DDL: "smallint"},
			// &drivercaps.ColumnDefn{DDL: "smallmoney"}, // TODO(js) Probably panics also?
			&drivercaps.ColumnDefn{DDL: "tinyint"},
			// Approximate numerics
			&drivercaps.ColumnDefn{DDL: "float"},
			&drivercaps.ColumnDefn{DDL: "real"},
			// Date and time
			&drivercaps.ColumnDefn{DDL: "date"},
			&drivercaps.ColumnDefn{DDL: "datetime2"},
			&drivercaps.ColumnDefn{DDL: "datetime"},
			&drivercaps.ColumnDefn{DDL: "datetimeoffset"},
			&drivercaps.ColumnDefn{DDL: "smalldatetime"},
			&drivercaps.ColumnDefn{DDL: "time"},
			// Character strings
			&drivercaps.ColumnDefn{DDL: "char"},
			&drivercaps.ColumnDefn{DDL: "text"},
			&drivercaps.ColumnDefn{DDL: "varchar"},
			// Unicode character strings
			&drivercaps.ColumnDefn{DDL: "nchar"},
			&drivercaps.ColumnDefn{DDL: "ntext"},
			&drivercaps.ColumnDefn{DDL: "nvarchar"},
			// Binary strings
			&drivercaps.ColumnDefn{DDL: "binary"},
			&drivercaps.ColumnDefn{DDL: "image"},
			&drivercaps.ColumnDefn{DDL: "varbinary"},
			// Other
			// // &drivercaps.ColumnDefn{DDL: "cursor"}, // TODO(js)
			// &drivercaps.ColumnDefn{DDL: "hierarchyid"},
			&drivercaps.ColumnDefn{DDL: "sql_variant"},
			// //			&drivercaps.ColumnDefn{DDL: ""}, // TODO(js) Spatial geometry types.
			// // &drivercaps.ColumnDefn{DDL: "table"}, // TODO(js)
			&drivercaps.ColumnDefn{DDL: "rowversion", OnePerTable: true},
			&drivercaps.ColumnDefn{DDL: "uniqueidentifier"},
			&drivercaps.ColumnDefn{DDL: "xml"},
			// //			&drivercaps.ColumnDefn{DDL: ""}, // TODO(js) Spatial geography types.
		},
	}
	return t
}
