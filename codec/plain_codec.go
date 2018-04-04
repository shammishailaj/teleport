// Copyright 2015-2017 HenryLee. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package codec

import (
	"fmt"

	"github.com/henrylee2cn/goutil"
)

// plain codec name and id
const (
	NAME_PLAIN = "plain"
	ID_PLAIN   = 's'
)

func init() {
	Reg(new(PlainCodec))
}

// PlainCodec plain codec
type PlainCodec struct{}

// Name returns codec name.
func (PlainCodec) Name() string {
	return NAME_PLAIN
}

// Id returns codec id.
func (PlainCodec) Id() byte {
	return ID_PLAIN
}

// Marshal returns the string encoding of v.
func (PlainCodec) Marshal(v interface{}) ([]byte, error) {
	var b []byte
	switch s := v.(type) {
	case nil:
	case string:
		b = goutil.StringToBytes(s)
	case *string:
		b = goutil.StringToBytes(*s)
	case []byte:
		b = s
	case *[]byte:
		b = *s
	default:
		return nil, fmt.Errorf("string codec: %T can not be directly converted to []byte type", v)
	}
	return b, nil
}

// Unmarshal parses the string-encoded data and stores the result
// in the value pointed to by v.
func (PlainCodec) Unmarshal(data []byte, v interface{}) error {
	switch s := v.(type) {
	case nil:
		return nil
	case *string:
		*s = string(data)
	case []byte:
		copy(s, data)
	case *[]byte:
		*s = make([]byte, len(data))
		copy(*s, data)
	default:
		return fmt.Errorf("string codec: []byte can not be directly converted to %T type", v)
	}
	return nil
}