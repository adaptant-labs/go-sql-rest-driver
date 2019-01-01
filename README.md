# Go-SQL-REST-Driver

A simple driver Go's [database/sql](https://golang.org/pkg/database/sql/)
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

## License

go-sql-rest-driver is licensed under the terms of the Apache 2.0 license, the full
version of which can be found in the [LICENSE](https://raw.githubusercontent.com/adaptant-labs/go-sql-rest-driver/master/LICENSE)
file included in the distribution.
