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

package restsql

import (
	"database/sql/driver"
	"io"
)

type resultSet struct {
	// As we deal with dynamic JSON responses, leave this as an interface
	// type for client-side unmarshalling and struct mapping.
	data	interface{}
}

type jsonRows struct {
	rc	*restsqlConn
	rs	resultSet
}

// By default we return a single column, which embodies the entire row response from the query.
func (rows *jsonRows) Columns() []string {
	return []string{""}
}

func (rows *jsonRows) NextResultSet() error {
	return io.EOF
}

func (rows *jsonRows) Next(dest []driver.Value) error {
	dest[0] = rows.rs.data;
	return nil
}

func (rows *jsonRows) Close() (err error) {
	return
}
