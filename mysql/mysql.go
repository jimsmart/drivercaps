package mysql

import (
	"github.com/jimsmart/drivercaps"
)

// Setup script:
//
// CREATE DATABASE test_db;
// CREATE USER test_user IDENTIFIED BY 'password';
// GRANT ALL ON test_db.* TO 'test_user';

// Or define docker vars:
// MYSQL_ALLOW_EMPTY_PASSWORD yes
// MYSQL_USER test_user
// MYSQL_PASSWORD password
// MYSQL_DATABASE test_db

var (
	User string = "test_user"
	Pass string = "password"
	Host string = "localhost"
	// Port string = "3306"
	Port string = "32784"
	DB   string = "test_db"
)

func NewTest(pkgURL, pkgName, drvName, connStr string) *drivercaps.DriverTest {
	t := &drivercaps.DriverTest{
		Database: "MySQL",
		PkgURL:   pkgURL,
		PkgName:  pkgName,
		DrvName:  drvName,
		ConnStr:  connStr,
		Columns: []*drivercaps.ColumnDefn{
			// https://dev.mysql.com/doc/refman/5.7/en/data-types.html
			// https://dev.mysql.com/doc/refman/5.7/en/numeric-types.html
			&drivercaps.ColumnDefn{DDL: "BIT"},
			&drivercaps.ColumnDefn{DDL: "BIT (8)"},
			&drivercaps.ColumnDefn{DDL: "TINYINT"}, // TODO(js)  TINYINT[(M)] [UNSIGNED] [ZEROFILL]
			// &drivercaps.ColumnDefn{DDL: "TINYINT (8)"}, // TODO(js) What does this even mean?
			&drivercaps.ColumnDefn{DDL: "TINYINT UNSIGNED"},
			&drivercaps.ColumnDefn{DDL: "BOOL"},
			&drivercaps.ColumnDefn{DDL: "BOOLEAN"},
			&drivercaps.ColumnDefn{DDL: "SMALLINT"}, // TODO(js) SMALLINT[(M)] [UNSIGNED] [ZEROFILL]
			// &drivercaps.ColumnDefn{DDL: "SMALLINT (8)"}, // TODO(js) What does this even mean?
			&drivercaps.ColumnDefn{DDL: "SMALLINT UNSIGNED"},
			&drivercaps.ColumnDefn{DDL: "MEDIUMINT"}, // TODO(js) MEDIUMINT[(M)] [UNSIGNED] [ZEROFILL]
			// &drivercaps.ColumnDefn{DDL: "MEDIUMINT (8)"}, // TODO(js) What does this even mean?
			&drivercaps.ColumnDefn{DDL: "MEDIUMINT UNSIGNED"},
			&drivercaps.ColumnDefn{DDL: "INT"}, // TODO(js) INT[(M)] [UNSIGNED] [ZEROFILL]
			// &drivercaps.ColumnDefn{DDL: "INT (8)"}, // TODO(js) What does this even mean?
			&drivercaps.ColumnDefn{DDL: "INT UNSIGNED"},
			&drivercaps.ColumnDefn{DDL: "INTEGER"}, // TODO(js) INTEGER[(M)] [UNSIGNED] [ZEROFILL]
			// &drivercaps.ColumnDefn{DDL: "INTEGER (8)"}, // TODO(js) What does this even mean?
			&drivercaps.ColumnDefn{DDL: "INTEGER UNSIGNED"},
			&drivercaps.ColumnDefn{DDL: "BIGINT"}, // TODO(js) BIGINT[(M)] [UNSIGNED] [ZEROFILL]
			// &drivercaps.ColumnDefn{DDL: "BIGINT (8)"}, // TODO(js) What does this even mean?
			&drivercaps.ColumnDefn{DDL: "BIGINT UNSIGNED"},
			&drivercaps.ColumnDefn{DDL: "DECIMAL"}, // TODO(js) DECIMAL[(M[,D])] [UNSIGNED] [ZEROFILL]
			&drivercaps.ColumnDefn{DDL: "DECIMAL UNSIGNED"},
			&drivercaps.ColumnDefn{DDL: "DECIMAL (8,4)"},
			&drivercaps.ColumnDefn{DDL: "DECIMAL (8,4) UNSIGNED"},
			&drivercaps.ColumnDefn{DDL: "DEC"}, // TODO(js) DEC[(M[,D])] [UNSIGNED] [ZEROFILL]
			&drivercaps.ColumnDefn{DDL: "DEC UNSIGNED"},
			&drivercaps.ColumnDefn{DDL: "DEC (8,4)"},
			&drivercaps.ColumnDefn{DDL: "DEC (8,4) UNSIGNED"},
			&drivercaps.ColumnDefn{DDL: "NUMERIC"}, // TODO(js) NUMERIC[(M[,D])] [UNSIGNED] [ZEROFILL]
			&drivercaps.ColumnDefn{DDL: "NUMERIC UNSIGNED"},
			&drivercaps.ColumnDefn{DDL: "NUMERIC (8,4)"},
			&drivercaps.ColumnDefn{DDL: "NUMERIC (8,4) UNSIGNED"},
			&drivercaps.ColumnDefn{DDL: "FIXED"}, // TODO(js) FIXED[(M[,D])] [UNSIGNED] [ZEROFILL]
			&drivercaps.ColumnDefn{DDL: "FIXED UNSIGNED"},
			&drivercaps.ColumnDefn{DDL: "FIXED (8,4)"},
			&drivercaps.ColumnDefn{DDL: "FIXED (8,4) UNSIGNED"},
			&drivercaps.ColumnDefn{DDL: "FLOAT"}, // TODO(js) FLOAT[(M,D)] [UNSIGNED] [ZEROFILL]
			&drivercaps.ColumnDefn{DDL: "FLOAT (5,2)"},
			&drivercaps.ColumnDefn{DDL: "FLOAT (100,30)"},
			&drivercaps.ColumnDefn{DDL: "FLOAT UNSIGNED"},
			&drivercaps.ColumnDefn{DDL: "FLOAT (5,2) UNSIGNED"},
			&drivercaps.ColumnDefn{DDL: "FLOAT (100,30) UNSIGNED"},
			&drivercaps.ColumnDefn{DDL: "DOUBLE"}, // TODO(js) DOUBLE[(M,D)] [UNSIGNED] [ZEROFILL]
			&drivercaps.ColumnDefn{DDL: "DOUBLE (5,2)"},
			&drivercaps.ColumnDefn{DDL: "DOUBLE (100,30)"},
			&drivercaps.ColumnDefn{DDL: "DOUBLE UNSIGNED"},
			&drivercaps.ColumnDefn{DDL: "DOUBLE (5,2) UNSIGNED"},
			&drivercaps.ColumnDefn{DDL: "DOUBLE (100,30) UNSIGNED"},
			&drivercaps.ColumnDefn{DDL: "DOUBLE PRECISION"}, // TODO(js) DOUBLE PRECISION[(M,D)] [UNSIGNED] [ZEROFILL]
			&drivercaps.ColumnDefn{DDL: "DOUBLE PRECISION (5,2)"},
			&drivercaps.ColumnDefn{DDL: "DOUBLE PRECISION (100,30)"},
			&drivercaps.ColumnDefn{DDL: "DOUBLE PRECISION UNSIGNED"},
			&drivercaps.ColumnDefn{DDL: "DOUBLE PRECISION (5,2) UNSIGNED"},
			&drivercaps.ColumnDefn{DDL: "DOUBLE PRECISION (100,30) UNSIGNED"},
			&drivercaps.ColumnDefn{DDL: "REAL"}, // TODO(js) REAL[(M,D)] [UNSIGNED] [ZEROFILL]
			&drivercaps.ColumnDefn{DDL: "REAL (5,2)"},
			&drivercaps.ColumnDefn{DDL: "REAL (100,30)"},
			&drivercaps.ColumnDefn{DDL: "REAL UNSIGNED"},
			&drivercaps.ColumnDefn{DDL: "REAL (5,2) UNSIGNED"},
			&drivercaps.ColumnDefn{DDL: "REAL (100,30) UNSIGNED"},
			&drivercaps.ColumnDefn{DDL: "FLOAT (10)"}, // TODO(js) FLOAT(P) [UNSIGNED] [ZEROFILL]
			&drivercaps.ColumnDefn{DDL: "FLOAT (30)"},
			&drivercaps.ColumnDefn{DDL: "FLOAT (10) UNSIGNED"},
			&drivercaps.ColumnDefn{DDL: "FLOAT (30) UNSIGNED"},
			// https://dev.mysql.com/doc/refman/5.7/en/date-and-time-types.html
			&drivercaps.ColumnDefn{DDL: "DATE"},
			&drivercaps.ColumnDefn{DDL: "DATETIME"},
			// &drivercaps.ColumnDefn{DDL: "TIMESTAMP"}, // TODO(js) Can only have one of these per table, we get two due to auto NOT NULL.
			// TODO(js) Test this next line:
			// &drivercaps.ColumnDefn{DDL: "TIMESTAMP DEFAULT 0", OnePerTable: true},
			&drivercaps.ColumnDefn{DDL: "TIME"},
			&drivercaps.ColumnDefn{DDL: "YEAR"},
			// &drivercaps.ColumnDefn{DDL: "YEAR (2)"},
			&drivercaps.ColumnDefn{DDL: "YEAR (4)"},
			// https://dev.mysql.com/doc/refman/5.7/en/string-types.html
			&drivercaps.ColumnDefn{DDL: "CHAR"},
			&drivercaps.ColumnDefn{DDL: "CHAR (8)"},
			&drivercaps.ColumnDefn{DDL: "VARCHAR (8)"},
			&drivercaps.ColumnDefn{DDL: "BINARY"},
			&drivercaps.ColumnDefn{DDL: "BINARY (8)"},
			&drivercaps.ColumnDefn{DDL: "VARBINARY (8)"},
			&drivercaps.ColumnDefn{DDL: "TINYBLOB"},
			&drivercaps.ColumnDefn{DDL: "BLOB"},
			&drivercaps.ColumnDefn{DDL: "BLOB (8)"},
			&drivercaps.ColumnDefn{DDL: "MEDIUMBLOB"},
			&drivercaps.ColumnDefn{DDL: "LONGBLOB"},
			&drivercaps.ColumnDefn{DDL: "TINYTEXT"},
			&drivercaps.ColumnDefn{DDL: "TEXT"},
			&drivercaps.ColumnDefn{DDL: "TEXT (8)"},
			&drivercaps.ColumnDefn{DDL: "MEDIUMTEXT"},
			&drivercaps.ColumnDefn{DDL: "LONGTEXT"},
			&drivercaps.ColumnDefn{DDL: "LONG"},
			&drivercaps.ColumnDefn{DDL: "LONG VARCHAR"},
			&drivercaps.ColumnDefn{DDL: "ENUM ('a','b','c')"},
			&drivercaps.ColumnDefn{DDL: "SET ('a','b','c')"},
			// https://dev.mysql.com/doc/refman/5.7/en/json.html
			&drivercaps.ColumnDefn{DDL: "JSON"},
		},
		PKType: "INTEGER",
	}
	return t
}
