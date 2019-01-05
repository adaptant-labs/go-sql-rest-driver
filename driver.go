// Copyright 2018-2019 Adaptant Labs
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

// Package restsql provides a driver for Go's database/sql package supporting
// SQL over REST API endpoints.
package restsql

import (
	"database/sql"
	"database/sql/driver"
)

// RESTSQLDriver is exported to make the driver directly accessible where
// needed. General usage is expected to be constrained to the database/sql
// APIs.
type RESTSQLDriver struct {
}

// Open prepares a destination URL endpoint that raw queries are appended to.
// The actual establishment of the connection is deferred until Query time.
func (d RESTSQLDriver) Open(url string) (driver.Conn, error) {
	var err error

	rc := &restsqlConn{url: url}

	return rc, err
}

func init() {
	sql.Register("restsql", &RESTSQLDriver{})
}
