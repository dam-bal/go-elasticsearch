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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/scriptlanguage"
)

// StoredScript type.
//
// https://github.com/elastic/elasticsearch-specification/blob/e279583a47508af40eb07b84694c5aae7885aa09/specification/_types/Scripting.ts#L47-L57
type StoredScript struct {
	// Lang Specifies the language the script is written in.
	Lang    scriptlanguage.ScriptLanguage `json:"lang"`
	Options map[string]string             `json:"options,omitempty"`
	// Source The script source.
	Source string `json:"source"`
}

func (s *StoredScript) UnmarshalJSON(data []byte) error {

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

		case "lang":
			if err := dec.Decode(&s.Lang); err != nil {
				return err
			}

		case "options":
			if s.Options == nil {
				s.Options = make(map[string]string, 0)
			}
			if err := dec.Decode(&s.Options); err != nil {
				return err
			}

		case "source":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Source = o

		}
	}
	return nil
}

// NewStoredScript returns a StoredScript.
func NewStoredScript() *StoredScript {
	r := &StoredScript{
		Options: make(map[string]string, 0),
	}

	return r
}
