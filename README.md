# drivercaps

[![BSD3](https://img.shields.io/badge/license-BSD3-blue.svg?style=flat)](LICENSE.md)

drivercaps is a [Go](https://golang.org) package to analyse and report database/sql drivers' column metadata capabilities.

TODO more docs

## Why?

Different database drivers have varying levels of support for [sql.ColumnType](https://golang.org/pkg/database/sql/#ColumnType), this project shines some light on the subject.

## What?

Currently supporting the following database engines:

- Microsoft SQL Server
- MySQL
- Oracle
- Postgres
- SQLite

## Driver sql.ColumnType Capability Reports

TODO generate report table (vs manual maintenance)

Driver | Database | Report | .Name<br/>Support | .DBTypeName<br/>Support | .Nullable<br/>Support | .DecimalSize<br/>Support | .Length<br/>Support | .ScanType<br/>Support
---|---|---|---|---|---|---|---|---
github.com/denisenkom/go-mssqldb | MS SQL Server | [View](https://github.com/jimsmart/drivercaps/tree/master/mssql/denisenkom) | Yes | Yes | Yes | Yes | Yes | Typed
github.com/minus5/gofreetds | MS SQL Server | [View](https://github.com/jimsmart/drivercaps/tree/master/mssql/minus5) | Yes | No | No | No | No | Default &#91;1&#93;
github.com/go-sql-driver/mysql | MySQL | [View](https://github.com/jimsmart/drivercaps/tree/master/mysql/gosqldriver) | Yes | Yes | Yes | Mostly &#91;2&#93; | No | Typed
github.com/ziutek/mymysql | MySQL | [View](https://github.com/jimsmart/drivercaps/tree/master/mysql/ziutek) | Yes | No | No | No | No | Default &#91;1&#93;
github.com/go-goracle/goracle | Oracle | [View](https://github.com/jimsmart/drivercaps/tree/master/oracle/goracle) | Yes | Yes | Yes | Mostly &#91;3&#93; | Yes | Typed
github.com/mattn/go-oci8 | Oracle | [View](https://github.com/jimsmart/drivercaps/tree/master/oracle/mattn) | Yes | Invalid &#91;4&#93; | No | No | Yes | Invalid &#91;5&#93;
github.com/rana/ora | Oracle | [View](https://github.com/jimsmart/drivercaps/tree/master/oracle/rana) | Yes | Yes | Invalid &#91;6&#93; | Kinda &#91;7&#93; | Yes | Typed &#91;8&#93;
github.com/jackc/pgx | Postgres | [View](https://github.com/jimsmart/drivercaps/tree/master/postgres/jackc) | Yes | Yes | No | Kinda &#91;9&#93; | Yes | Typed
github.com/jbarham/gopgsqldriver | Postgres | [View](https://github.com/jimsmart/drivercaps/tree/master/postgres/jbarham) | Yes | No | No | No | No | Default &#91;1&#93;
github.com/lib/pq | Postgres | [View](https://github.com/jimsmart/drivercaps/tree/master/postgres/lib) | Yes | Yes | No | Kinda &#91;9&#93; | Yes | Typed
github.com/gwenn/gosqlite | SQLite | [View](https://github.com/jimsmart/drivercaps/tree/master/sqlite/gwenn) | Yes | Yes &#91;10&#93; | No | No | No | Invalid &#91;11&#93;
github.com/mattn/go-sqlite3 | SQLite | [View](https://github.com/jimsmart/drivercaps/tree/master/sqlite/mattn) | Yes | Yes &#91;10&#93; | Invalid &#91;6&#93; | No | No | Invalid &#91;11&#93;
github.com/mxk/go-sqlite | SQLite | [View](https://github.com/jimsmart/drivercaps/tree/master/sqlite/mxk) | Yes | No | No | No | No | Default &#91;1&#93;

&#91;1&#93; interface{} only<br/>
&#91;2&#93; float and double types have invalid precision and scale value MaxInt64<br/>
&#91;3&#93; number and float have invalid default scale value -127, float type has invalid? default precision value 126<br/>
&#91;4&#93; exposes only internal type codenames<br/>
&#91;5&#93; invalid scan type []string<br/>
&#91;6&#93; reports everything as nullable<br/>
&#91;7&#93; number and float types have invalid scale value -127<br/>
&#91;8&#93; binary_float and binary_double have invalid scan type nil<br/>
&#91;9&#93; numeric and decimal types have invalid default precision and scale values 65535, 65531<br/>
&#91;10&#93; type name includes length and precision values<br/>
&#91;11&#93; invalid scan type nil

## Installation

To install and produce your own reports locally:

```bash
$ go get github.com/jimsmart/drivercaps
```

## Usage

Change directory to the drivercaps root folder and execute:

```bash
$ go test -test.v ./...
```

## Output

After execution, capability reports for each driver can be found in the leaf folders of the package, e.g.

- drivercaps/oracle/goracle/report.csv
- drivercaps/oracle/goracle/report.txt
- drivercaps/oracle/goracle/README.md
- drivercaps/postgres/jackc/report.csv
- etc.

## License

Package drivercaps is copyright 2018 by Jim Smart and released under the [BSD 3-Clause License](LICENSE.md)
