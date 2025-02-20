// Copyright 2022 Juan Pablo Tosso
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package variables

import (
	"strings"
	"testing"
)

func TestNameToVariable(t *testing.T) {
	vars := []string{"URLENCODED_ERROR", "RESPONSE_CONTENT_TYPE", "UNIQUE_ID", "ARGS_COMBINED_SIZE", "AUTH_TYPE", "FILES_COMBINED_SIZE", "FULL_REQUEST", "FULL_REQUEST_LENGTH", "INBOUND_DATA_ERROR", "MATCHED_VAR", "MATCHED_VAR_NAME", "MULTIPART_BOUNDARY_QUOTED", "MULTIPART_BOUNDARY_WHITESPACE", "MULTIPART_CRLF_LF_LINES", "MULTIPART_DATA_AFTER", "MULTIPART_DATA_BEFORE", "MULTIPART_FILE_LIMIT_EXCEEDED", "MULTIPART_HEADER_FOLDING", "MULTIPART_INVALID_HEADER_FOLDING", "MULTIPART_INVALID_PART", "MULTIPART_INVALID_QUOTING", "MULTIPART_LF_LINE", "MULTIPART_MISSING_SEMICOLON", "MULTIPART_STRICT_ERROR", "MULTIPART_UNMATCHED_BOUNDARY", "OUTBOUND_DATA_ERROR", "PATH_INFO", "QUERY_STRING", "REMOTE_ADDR", "REMOTE_HOST", "REMOTE_PORT", "REQBODY_ERROR", "REQBODY_ERROR_MSG", "REQBODY_PROCESSOR_ERROR", "REQBODY_PROCESSOR_ERROR_MSG", "REQBODY_PROCESSOR", "REQUEST_BASENAME", "REQUEST_BODY", "REQUEST_BODY_LENGTH", "REQUEST_FILENAME", "REQUEST_LINE", "REQUEST_METHOD", "REQUEST_PROTOCOL", "REQUEST_URI", "REQUEST_URI_RAW", "RESPONSE_BODY", "RESPONSE_CONTENT_LENGTH", "RESPONSE_PROTOCOL", "RESPONSE_STATUS", "SERVER_ADDR", "SERVER_NAME", "SERVER_PORT", "SESSIONID", "RESPONSE_HEADERS_NAMES", "REQUEST_HEADERS_NAMES", "USERID", "ARGS", "ARGS_GET", "ARGS_POST", "FILES_SIZES", "FILES_NAMES", "FILES_TMP_CONTENT", "MULTIPART_FILENAME", "MULTIPART_NAME", "MATCHED_VARS_NAMES", "MATCHED_VARS", "FILES", "REQUEST_COOKIES", "REQUEST_HEADERS", "RESPONSE_HEADERS", "GEO", "REQUEST_COOKIES_NAMES", "FILES_TMPNAMES", "ARGS_NAMES", "ARGS_GET_NAMES", "ARGS_POST_NAMES", "RULE", "XML", "TX", "DURATION"}
	for _, v := range vars {
		_, err := Parse(v)
		if err != nil {
			t.Error("failed to test variable " + v)
		}
	}
}

func TestVariableToName(t *testing.T) {
	if len(rulemap) != len(rulemapRev) {
		t.Error("rulemap and rulemapRev are not equal")
	}
	for i := 0; i < len(rulemap); i++ {
		v := RuleVariable(i)
		if v.Name() != rulemap[v] || rulemapRev[v.Name()] != v {
			t.Errorf("failed to test variable %s (%d)", v.Name(), v)
		}
	}
}

func TestParseVariable(t *testing.T) {
	for i := 0; i < len(rulemap); i++ {
		v := RuleVariable(i)
		name := v.Name()
		// we create some lowercase cases
		if i%3 == 0 {
			name = strings.ToLower(name)
		}
		vv, err := Parse(name)
		if err != nil {
			t.Errorf("failed to test variable %s", name)
		}
		if v != vv {
			t.Errorf("failed to test variable %s", name)
		}
	}
}
