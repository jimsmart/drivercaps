# drivercaps

[![BSD3](https://img.shields.io/badge/license-BSD3-blue.svg?style=flat)](LICENSE.md)

drivercaps is a [Go](https://golang.org) package to analyse and report database/sql drivers' column metadata capabilities.


## Why?

Different database drivers have varying levels of support for sql.ColumnType, this project shines some light on the subject.


## 'Under Construction'

TODO generate links to all reports

TODO docs

TODO Microsoft SQL Server support


Currently supporting the following database engines:

- SQLite
- Postgres
- Oracle
- MySQL


## Installation
```bash
$ go get github.com/jimsmart/drivercaps
```

## Usage

Change directory to the drivercaps root folder and execute:

```bash
$ go test -test.v ./...
```

## Output/Reports

After execution, capability reports for each driver can be found in the leaf folders of this package, e.g.

- [https://github.com/jimsmart/drivercaps/oracle/goracle](https://github.com/jimsmart/drivercaps/tree/master/oracle/goracle)
- [https://github.com/jimsmart/drivercaps/postgres/jackc](https://github.com/jimsmart/drivercaps/tree/master/postgres/jackc)
- etc.


## License

Package drivercaps is copyright 2018 by Jim Smart and released under the [BSD 3-Clause License](LICENSE.md)
