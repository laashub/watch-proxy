// Copyright 2018 Heptio
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

package emitter

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/glog"
	"github.com/heptio/quartermaster/config"
)

func TestEmitChanges(t *testing.T) {

	reqBody := `{"data":[{"asset_type_id":"","data":{"metadata":{"creationTimestamp":"1970-01-01T00:00:00Z","name":"test","resourceVersion":"0000","selfLink":"/api/v1/namespaces/test","uid":"00000000-0000-0000-0000-000000000000"},"spec":{"finalizers":["kubernetes"]},"status":{"phase":"Active"}},"uniqueId":"00000000-0000-0000-0000-000000000000","event":"add"}],"meta":null}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contents, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Errorf("Test HTTP Server had error reading body")
		}
		glog.Infof("%s\n", string(contents))
		if string(contents) != reqBody {
			t.Errorf("Request body is wrong, got: %v, want: %v.",
				string(contents), reqBody)
		}
	}))

	payloadJson := `{
"metadata": {
	"name":"test",
	"selfLink":"/api/v1/namespaces/test",
	"uid":"00000000-0000-0000-0000-000000000000",
	"resourceVersion":"0000",
	"creationTimestamp":"1970-01-01T00:00:00Z"},
	"spec": {
		"finalizers": [
			"kubernetes"
		]
	},
	"status": {
		"phase":"Active"
	}
}
`
	payload := make(map[string]interface{})
	json.Unmarshal([]byte(payloadJson), &payload)

	loadCache(config.Config{EmitCacheDuration: "60m"})

	tests := []struct {
		name     string
		emission Emission
	}{
		{
			name: "test-namespace",
			emission: Emission{
				EmitType: "http",
				Client:   http.Client{Timeout: time.Second * 2},
				HttpUrl:  ts.URL,
				EmittableList: []EmitObject{
					EmitObject{
						Payload:   payload,
						ObjType:   "namespace",
						Key:       "test",
						UID:       "00000000-0000-0000-0000-000000000000",
						EventType: "add",
					},
				},
			},
		},
	}

	defer ts.Close()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			EmitChanges(tt.emission)
		})
	}
}
