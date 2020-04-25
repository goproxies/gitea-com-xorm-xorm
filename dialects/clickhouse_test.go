// Copyright 2020 The Xorm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dialects

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"xorm.io/xorm/schemas"
)

func TestParseClickHouse(t *testing.T) {
	uri, err := ParseClickHouse("tcp://host1:9000?username=user&password=qwerty&database=clicks&read_timeout=10&write_timeout=20&alt_hosts=host2:9000,host3:9000")
	assert.NoError(t, err)

	assert.EqualValues(t, &URI{
		DBType: schemas.CLICKHOUSE,
		Proto:  "tcp",
		Host:   "host1",
		Port:   "9000",
		DBName: "clicks",
		User:   "user",
		Passwd: "qwerty",
	}, uri)
}
