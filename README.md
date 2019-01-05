# Go-SQL-REST-Driver

[![godoc](https://godoc.org/github.com/adaptant-labs/go-sql-rest-driver?status.svg)](http://godoc.org/github.com/adaptant-labs/go-sql-rest-driver)
[![Build Status](https://travis-ci.com/adaptant-labs/go-sql-rest-driver.svg?branch=master)](https://travis-ci.com/adaptant-labs/go-sql-rest-driver)
[![Go Report Card](https://goreportcard.com/badge/github.com/adaptant-labs/go-sql-rest-driver)](https://goreportcard.com/report/github.com/adaptant-labs/go-sql-rest-driver)

A simple driver for Go's [database/sql](https://golang.org/pkg/database/sql/)
package for SQL over REST API endpoints, such as those found within the
RestAssured Query Gateway.

## Installation

Install:

```shell
go get -u github.com/adaptant-labs/go-sql-rest-driver
```

## Usage
_Go SQL REST Driver_ is a limited implementation of Go's `database/sql/driver` interface. You only need to import the driver and can then use the [`database/sql`](https://golang.org/pkg/database/sql/) API.

Use `mysql` as `driverName` and a valid REST API Endpoint:
```go
import "database/sql"
import _ "github.com/adaptant-labs/go-sql-rest-driver"

db, err := sql.Open("restsql", "http://localhost:9000/query/v1/")
```

## Online Documentation

Limited API documentation for the driver interface itself is provided through godoc, this can be accessed
directly on the [package entry](https://godoc.org/github.com/adaptant-labs/go-sql-rest-driver)
in the godoc package repository. In general, however, the [`database/sql`](https://golang.org/pkg/database/sql/)
API should always be the main method of invocation.

## Acknowledgements

This project has received funding from the European Union’s Horizon 2020 research and innovation programme under grant agreement No 731678.

## License

go-sql-rest-driver is licensed under the terms of the Apache 2.0 license, the full
version of which can be found in the [LICENSE](https://raw.githubusercontent.com/adaptant-labs/go-sql-rest-driver/master/LICENSE)
file included in the distribution.
