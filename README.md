# drivercaps

[![BSD3](https://img.shields.io/badge/license-BSD3-blue.svg?style=flat)](LICENSE.md)

drivercaps is a [Go](https://golang.org) package to analyse and report database/sql drivers' column metadata capabilities.


## Why?

Different database drivers have varying levels of support for sql.ColumnType, this project shines some light on the subject.


## 'Under Construction'

TODO skip to Output/Reports, below, if you just want to view the capability reports.

TODO generate report links list

TODO summarise reports

TODO more docs


Currently supporting the following database engines:

- SQLite
- Postgres
- Microsoft SQL Server
- Oracle
- MySQL

## Reports

Driver | Database | Report
---|---|---
github.com/denisenkom/go-mssqldb | MS SQL | [View Report](https://github.com/jimsmart/drivercaps/tree/master/mssql/denisenkom)
github.com/minus5/gofreetds | MS SQL | [View Report](https://github.com/jimsmart/drivercaps/tree/master/mssql/minus5)
github.com/go-sql-driver/mysql | MySQL | [View Report](https://github.com/jimsmart/drivercaps/tree/master/mysql/gosqldriver)
github.com/ziutek/mymysql | MySQL | [View Report](https://github.com/jimsmart/drivercaps/tree/master/mysql/ziutek)
github.com/go-goracle/goracle | Oracle | [View Report](https://github.com/jimsmart/drivercaps/tree/master/oracle/goracle)
github.com/mattn/go-oci8 | Oracle | [View Report](https://github.com/jimsmart/drivercaps/tree/master/oracle/mattn)
github.com/rana/ora | Oracle | [View Report](https://github.com/jimsmart/drivercaps/tree/master/oracle/rana)
github.com/jackc/pgx | Postgres | [View Report](https://github.com/jimsmart/drivercaps/tree/master/postgres/jackc)
github.com/jbarham/gopgsqldriver | Postgres | [View Report](https://github.com/jimsmart/drivercaps/tree/master/postgres/jbarham)
github.com/lib/pq | Postgres | [View Report](https://github.com/jimsmart/drivercaps/tree/master/postgres/lib)
github.com/gwenn/gosqlite | SQLite | [View Report](https://github.com/jimsmart/drivercaps/tree/master/sqlite/gwenn)
github.com/mattn/go-sqlite3 | SQLite | [View Report](https://github.com/jimsmart/drivercaps/tree/master/sqlite/mattn)
github.com/mxk/go-sqlite | SQLite | [View Report](https://github.com/jimsmart/drivercaps/tree/master/sqlite/mxk)


## Installation

To install and produce your own reports:

```bash
$ go get github.com/jimsmart/drivercaps
```

## Usage

Change directory to the drivercaps root folder and execute:

```bash
$ go test -test.v ./...
```

## Output

After execution, capability reports for each driver can be found in the leaf folders of this package, e.g.

- drivercaps/oracle/goracle/report.csv
- drivercaps/oracle/goracle/report.txt
- drivercaps/oracle/goracle/README.md
- drivercaps/postgres/jackc/report.csv
- etc.


## License

Package drivercaps is copyright 2018 by Jim Smart and released under the [BSD 3-Clause License](LICENSE.md)
