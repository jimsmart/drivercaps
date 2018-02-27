package postgres

import (
	"github.com/jimsmart/drivercaps"
)

var (
	User string = "postgres"
	Pass string
	Host string = "localhost"
	// Port string = "5432"
	Port string = "32785"
)

func NewTest(pkgURL, pkgName, drvName, connStr string) *drivercaps.DriverTest {
	t := &drivercaps.DriverTest{
		Database: "Postgres",
		PkgURL:   pkgURL,
		PkgName:  pkgName,
		DrvName:  drvName,
		ConnStr:  connStr,
		Columns: []*drivercaps.ColumnDefn{
			// https://www.postgresql.org/docs/9.5/static/datatype.html
			&drivercaps.ColumnDefn{DDL: "bigint"},
			&drivercaps.ColumnDefn{DDL: "int8"},
			&drivercaps.ColumnDefn{DDL: "bigserial"},
			// TODO(js) Look into this - maybe we can only have one serial field per table?
			// duplicate key value violates unique constraint.
			// DETAIL:  Key (typname, typnamespace)=(test_table_column_2_seq, 2200) already exists.
			// &drivercaps.ColumnDefn{DDL: "serial8"},
			&drivercaps.ColumnDefn{DDL: "bit"},
			&drivercaps.ColumnDefn{DDL: "bit (8)"},
			&drivercaps.ColumnDefn{DDL: "bit varying"},
			&drivercaps.ColumnDefn{DDL: "bit varying (8)"},
			&drivercaps.ColumnDefn{DDL: "varbit"},
			&drivercaps.ColumnDefn{DDL: "varbit (8)"},
			&drivercaps.ColumnDefn{DDL: "boolean"},
			&drivercaps.ColumnDefn{DDL: "bool"},
			&drivercaps.ColumnDefn{DDL: "box"},
			&drivercaps.ColumnDefn{DDL: "bytea"},
			&drivercaps.ColumnDefn{DDL: "character"},
			&drivercaps.ColumnDefn{DDL: "character (8)"},
			&drivercaps.ColumnDefn{DDL: "char"},
			&drivercaps.ColumnDefn{DDL: "char (8)"},
			&drivercaps.ColumnDefn{DDL: "character varying"},
			&drivercaps.ColumnDefn{DDL: "character varying (8)"},
			&drivercaps.ColumnDefn{DDL: "varchar"},
			&drivercaps.ColumnDefn{DDL: "varchar (8)"},
			&drivercaps.ColumnDefn{DDL: "cidr"},
			&drivercaps.ColumnDefn{DDL: "circle"},
			&drivercaps.ColumnDefn{DDL: "date"},
			&drivercaps.ColumnDefn{DDL: "double precision"},
			&drivercaps.ColumnDefn{DDL: "float8"},
			&drivercaps.ColumnDefn{DDL: "inet"},
			&drivercaps.ColumnDefn{DDL: "integer"},
			&drivercaps.ColumnDefn{DDL: "int"},
			&drivercaps.ColumnDefn{DDL: "int4"},
			&drivercaps.ColumnDefn{DDL: "interval"}, // TODO(js) interval [ fields ] [ (p) ]
			&drivercaps.ColumnDefn{DDL: "json"},
			&drivercaps.ColumnDefn{DDL: "jsonb"},
			&drivercaps.ColumnDefn{DDL: "line"},
			&drivercaps.ColumnDefn{DDL: "lseg"},
			&drivercaps.ColumnDefn{DDL: "macaddr"},
			&drivercaps.ColumnDefn{DDL: "money"},
			&drivercaps.ColumnDefn{DDL: "numeric"},
			&drivercaps.ColumnDefn{DDL: "numeric (8,4)"},
			&drivercaps.ColumnDefn{DDL: "decimal"},
			&drivercaps.ColumnDefn{DDL: "decimal (8,4)"},
			&drivercaps.ColumnDefn{DDL: "path"},
			&drivercaps.ColumnDefn{DDL: "pg_lsn"},
			&drivercaps.ColumnDefn{DDL: "point"},
			&drivercaps.ColumnDefn{DDL: "polygon"},
			&drivercaps.ColumnDefn{DDL: "real"},
			&drivercaps.ColumnDefn{DDL: "float4"},
			&drivercaps.ColumnDefn{DDL: "smallint"},
			&drivercaps.ColumnDefn{DDL: "int2"},
			&drivercaps.ColumnDefn{DDL: "smallserial"},
			&drivercaps.ColumnDefn{DDL: "serial2"},
			&drivercaps.ColumnDefn{DDL: "serial"},
			&drivercaps.ColumnDefn{DDL: "serial4"},
			&drivercaps.ColumnDefn{DDL: "text"},
			&drivercaps.ColumnDefn{DDL: "time"},                     // TODO(js) time [ (p) ] [ without time zone ]
			&drivercaps.ColumnDefn{DDL: "time with time zone"},      // TODO(js) time [ (p) ] with time zone
			&drivercaps.ColumnDefn{DDL: "timetz"},                   // TODO(js) timetz [ (p) ]
			&drivercaps.ColumnDefn{DDL: "timestamp"},                // TODO(js) timestamp [ (p) ] [ without time zone ]
			&drivercaps.ColumnDefn{DDL: "timestamp with time zone"}, // TODO(js) timestamp [ (p) ] without time zone
			&drivercaps.ColumnDefn{DDL: "timestamptz"},              // TODO(js) timestamptz [ (p) ]
			&drivercaps.ColumnDefn{DDL: "tsquery"},
			&drivercaps.ColumnDefn{DDL: "tsvector"},
			&drivercaps.ColumnDefn{DDL: "txid_snapshot"},
			&drivercaps.ColumnDefn{DDL: "uuid"},
			&drivercaps.ColumnDefn{DDL: "xml"},
		},
	}
	return t
}
