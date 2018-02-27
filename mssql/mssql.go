package mssql

import (
	"github.com/jimsmart/drivercaps"
)

var (
	User string = "sa"
	Pass string
	Host string = "localhost"
	Port string = "1433"
)

func NewTest(pkgURL, pkgName, drvName, connStr string) *drivercaps.DriverTest {
	t := &drivercaps.DriverTest{
		Database: "Microsoft SQL Server",
		PkgURL:   pkgURL,
		PkgName:  pkgName,
		DrvName:  drvName,
		ConnStr:  connStr,
		// Columns: []*drivercaps.ColumnDefn{
		// 	//
		// 	&drivercaps.ColumnDefn{DDL: "NUMBER"},
		// },
	}
	return t
}
