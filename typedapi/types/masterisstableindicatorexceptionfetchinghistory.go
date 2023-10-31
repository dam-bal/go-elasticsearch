// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/e279583a47508af40eb07b84694c5aae7885aa09

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// MasterIsStableIndicatorExceptionFetchingHistory type.
//
// https://github.com/elastic/elasticsearch-specification/blob/e279583a47508af40eb07b84694c5aae7885aa09/specification/_global/health_report/types.ts#L94-L97
type MasterIsStableIndicatorExceptionFetchingHistory struct {
	Message    string `json:"message"`
	StackTrace string `json:"stack_trace"`
}

func (s *MasterIsStableIndicatorExceptionFetchingHistory) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "message":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Message = o

		case "stack_trace":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.StackTrace = o

		}
	}
	return nil
}

// NewMasterIsStableIndicatorExceptionFetchingHistory returns a MasterIsStableIndicatorExceptionFetchingHistory.
func NewMasterIsStableIndicatorExceptionFetchingHistory() *MasterIsStableIndicatorExceptionFetchingHistory {
	r := &MasterIsStableIndicatorExceptionFetchingHistory{}

	return r
}
