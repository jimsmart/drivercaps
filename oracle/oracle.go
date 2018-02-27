package oracle

import (
	"github.com/jimsmart/drivercaps"
)

var (
	// User   string = "test_user"
	User string = "system"
	Pass string = "oracle"
	Host string = "localhost"
	// Port   string = "1521"
	Port   string = "32782"
	Schema string = "xe" // TODO(js) Schema or DB?
)

func NewTest(pkgURL, pkgName, drvName, connStr string) *drivercaps.DriverTest {
	t := &drivercaps.DriverTest{
		Database: "Oracle",
		PkgURL:   pkgURL,
		PkgName:  pkgName,
		DrvName:  drvName,
		ConnStr:  connStr,
		Columns: []*drivercaps.ColumnDefn{
			// https://docs.oracle.com/database/121/SQLQR/sqlqr06002.htm#SQLQR959
			&drivercaps.ColumnDefn{DDL: "CHAR"},
			&drivercaps.ColumnDefn{DDL: "CHAR (8)"},
			&drivercaps.ColumnDefn{DDL: "VARCHAR2 (8)"},
			&drivercaps.ColumnDefn{DDL: "NCHAR"},
			&drivercaps.ColumnDefn{DDL: "NCHAR (8)"},
			&drivercaps.ColumnDefn{DDL: "NVARCHAR2 (8)"},
			&drivercaps.ColumnDefn{DDL: "DATE"},
			&drivercaps.ColumnDefn{DDL: "TIMESTAMP"}, // TODO(js) TIMESTAMP [ (fractional_seconds_precision) ] [ WITH [ LOCAL ] TIME ZONE ]
			&drivercaps.ColumnDefn{DDL: "TIMESTAMP WITH TIME ZONE"},
			&drivercaps.ColumnDefn{DDL: "INTERVAL YEAR TO MONTH"}, // TODO(js) INTERVAL YEAR [ (year_precision) ] TO MONTH
			&drivercaps.ColumnDefn{DDL: "INTERVAL DAY TO SECOND"}, // TODO(js) INTERVAL DAY [ (day_precision) ] TO SECOND [ (fractional_seconds_precision) ]
			&drivercaps.ColumnDefn{DDL: "BLOB"},
			&drivercaps.ColumnDefn{DDL: "CLOB"},
			&drivercaps.ColumnDefn{DDL: "CLOB"},
			&drivercaps.ColumnDefn{DDL: "BFILE"},
			// &drivercaps.ColumnDefn{DDL: "LONG"}, // TODO(js) Bad? Perhaps only bad on only one driver?
			// &drivercaps.ColumnDefn{DDL: "LONG RAW"}, // TODO(js) Bad? Perhaps only bad on only one driver?
			&drivercaps.ColumnDefn{DDL: "RAW (8)"},
			&drivercaps.ColumnDefn{DDL: "NUMBER"},
			&drivercaps.ColumnDefn{DDL: "NUMBER (8)"},
			&drivercaps.ColumnDefn{DDL: "NUMBER (8,4)"},
			&drivercaps.ColumnDefn{DDL: "FLOAT"},
			&drivercaps.ColumnDefn{DDL: "FLOAT (30)"},
			&drivercaps.ColumnDefn{DDL: "BINARY_FLOAT"},
			&drivercaps.ColumnDefn{DDL: "BINARY_DOUBLE"},
			// TODO(js) None of these seem to work with goracle?
			// &drivercaps.ColumnDefn{DDL: "ROWID", OnePerTable: true},
			// // &drivercaps.ColumnDefn{DDL: "UROWID (8)", OnePerTable: true},
			// // https://docs.oracle.com/database/121/SQLQR/sqlqr06003.htm#SQLQR967
			// &drivercaps.ColumnDefn{DDL: "SYS.AnyData"},
			// &drivercaps.ColumnDefn{DDL: "SYS.AnyType"},
			// &drivercaps.ColumnDefn{DDL: "SYS.AnyDataSet"},
			// &drivercaps.ColumnDefn{DDL: "ORDAudio"},
			// &drivercaps.ColumnDefn{DDL: "ORDImage"},
			// &drivercaps.ColumnDefn{DDL: "ORDVideo"},
			// &drivercaps.ColumnDefn{DDL: "ORDDoc"},
			// &drivercaps.ColumnDefn{DDL: "ORDDicom"},
			// &drivercaps.ColumnDefn{DDL: "SI_StillImage"},
			// &drivercaps.ColumnDefn{DDL: "SI_AverageColor"},
			// &drivercaps.ColumnDefn{DDL: "SI_PositionalColor"},
			// &drivercaps.ColumnDefn{DDL: "SI_ColorHistogram"},
			// &drivercaps.ColumnDefn{DDL: "SI_Texture"},
			// &drivercaps.ColumnDefn{DDL: "SI_FeatureList"},
			// &drivercaps.ColumnDefn{DDL: "SI_Color"},
			// &drivercaps.ColumnDefn{DDL: "SDO_Geometry"},
			// &drivercaps.ColumnDefn{DDL: "SDO_Topo_Geometry"},
			// // &drivercaps.ColumnDefn{DDL: "SDO_GeoRaster"},
			// &drivercaps.ColumnDefn{DDL: "XMLType"},
			// &drivercaps.ColumnDefn{DDL: "URIType"},
		},
	}
	return t
}
