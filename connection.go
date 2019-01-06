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
	"context"
	"database/sql/driver"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type restsqlConn struct {
	client *http.Client
	url string
}

// Begin is stubbed out, but is necessary to satisfy the interface requirements
func (rc *restsqlConn) Begin() (driver.Tx, error) {
	return nil, nil
}

// Prepare is stubbed out, but is necessary to satisfy the interface
// requirements
func (rc *restsqlConn) Prepare(query string) (driver.Stmt, error) {
	return nil, nil
}

// Ping checks connectivity to the host http server by issuing a HEAD request.
func (rc *restsqlConn) Ping(ctx context.Context) (err error) {
	req, err := http.NewRequest("HEAD", rc.url, nil)
	if err != nil {
		return err
	}

	resp, err := rc.client.Do(req.WithContext(ctx))
	if err != nil {
		return err
	}

	resp.Body.Close()

	// Note that we do not particularly care about the response code, as we
	// only care about whether the endpoint is responsive or not.
	return nil
}

// Query carries out a basic SQL query on a REST API endpoint. This presently
// takes the form of the raw query being appended to the base path.
func (rc *restsqlConn) Query(query string, args []driver.Value) (driver.Rows, error) {
	res, err := http.Get(rc.url + query)
	if err != nil {
		return nil, driver.ErrBadConn
	}

	defer res.Body.Close()

	rows := new(jsonRows)
	rows.rc = rc

	dec := json.NewDecoder(res.Body)
	for {
		var data interface{}

		err = dec.Decode(&data)
		if err == io.EOF {
			return rows, nil
		} else if err != nil {
			log.Fatal("failed decoding JSON message")
		}

		rows.rs.data = append(rows.rs.data, data)
	}

	return rows, err
}

// Close is stubbed out, but is necessary to satisfy the interface requirements
func (rc *restsqlConn) Close() (err error) {
	rc.client = nil
	return nil
}
