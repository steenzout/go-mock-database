// Mock implementation of the database/sql package.
//
// Copyright 2016 Pedro Salgado
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
package sql

import (
	"database/sql"
	"database/sql/driver"
	"time"

	"github.com/stretchr/testify/mock"
)

// DB mock.
type DB struct {
	mock.Mock
}

// Begin starts a transaction.
// The isolation level is dependent on the driver.
func (db *DB) Begin() (*sql.Tx, error) {
	out := db.Called()
	return out.Get(0).(*sql.Tx), out.Error(1)
}

// Close mock.
func (db *DB) Close() error {
	out := db.Called()
	return out.Error(0)
}

// Driver mock.
func (db *DB) Driver() driver.Driver {
	out := db.Called()
	return out.Get(0).(driver.Driver)
}

// Query mock.
func (db *DB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	slice := make([]interface{}, 1)
	slice[0] = query
	out := db.Called(append(slice, args...))
	return out.Get(0).(*sql.Rows), out.Error(1)
}

// Ping mock.
func (db *DB) Ping() error {
	out := db.Called()
	return out.Error(0)
}

// Prepare mock.
func (db *DB) Prepare(query string) (*sql.Stmt, error) {
	out := db.Called(query)
	return out.Get(0).(*sql.Stmt), out.Error(1)
}

// QueryRow mock.
func (db *DB) QueryRow(query string, args ...interface{}) *sql.Row {
	slice := make([]interface{}, 1)
	slice[0] = query
	out := db.Called(append(slice, args...))
	return out.Get(0).(*sql.Row)
}

// SetConnMaxLifetime mock.
func (db *DB) SetConnMaxLifetime(d time.Duration) {
	db.Called(d)
}

// SetMaxIdleConns mock.
func (db *DB) SetMaxIdleConns(n int) {
	db.Called(n)
}

// Stats mock.
func (db *DB) Stats() sql.DBStats {
	out := db.Called()
	return out.Get(0).(sql.DBStats)
}
