package sqlite

import "github.com/jimsmart/drivercaps"

// TODO These should all come from environment vars.

func NewTest(pkgURL, pkgName, drvName, connStr string) *drivercaps.DriverTest {
	t := &drivercaps.DriverTest{
		Database: "SQLite",
		PkgURL:   pkgURL,
		PkgName:  pkgName,
		DrvName:  drvName,
		ConnStr:  connStr,
		Columns: []*drivercaps.ColumnDefn{
			// https://www.sqlite.org/datatype3.html
			// INTEGER
			&drivercaps.ColumnDefn{DDL: "INT"},
			&drivercaps.ColumnDefn{DDL: "INTEGER"},
			&drivercaps.ColumnDefn{DDL: "TINYINT"},
			&drivercaps.ColumnDefn{DDL: "SMALLINT"},
			&drivercaps.ColumnDefn{DDL: "MEDIUMINT"},
			&drivercaps.ColumnDefn{DDL: "BIGINT"},
			&drivercaps.ColumnDefn{DDL: "UNSIGNED BIG INT"},
			&drivercaps.ColumnDefn{DDL: "INT2"},
			&drivercaps.ColumnDefn{DDL: "INT8"},
			// TEXT
			&drivercaps.ColumnDefn{DDL: "CHARACTER(20)"},
			&drivercaps.ColumnDefn{DDL: "VARCHAR(255)"},
			&drivercaps.ColumnDefn{DDL: "VARYING CHARACTER(255)"},
			&drivercaps.ColumnDefn{DDL: "NCHAR(55)"},
			&drivercaps.ColumnDefn{DDL: "NATIVE CHARACTER(70)"},
			&drivercaps.ColumnDefn{DDL: "NVARCHAR(100)"},
			&drivercaps.ColumnDefn{DDL: "TEXT"},
			&drivercaps.ColumnDefn{DDL: "CLOB"},
			// BLOB
			&drivercaps.ColumnDefn{DDL: "BLOB"},
			// REAL
			&drivercaps.ColumnDefn{DDL: "REAL"},
			&drivercaps.ColumnDefn{DDL: "DOUBLE"},
			&drivercaps.ColumnDefn{DDL: "DOUBLE PRECISION"},
			&drivercaps.ColumnDefn{DDL: "FLOAT"},
			// NUMERIC
			&drivercaps.ColumnDefn{DDL: "NUMERIC"},
			&drivercaps.ColumnDefn{DDL: "DECIMAL(10,5)"},
			&drivercaps.ColumnDefn{DDL: "BOOLEAN"},
			&drivercaps.ColumnDefn{DDL: "DATE"},
			&drivercaps.ColumnDefn{DDL: "DATETIME"},
		},
		CreateOpt: "WITHOUT ROWID",
	}
	return t
}
