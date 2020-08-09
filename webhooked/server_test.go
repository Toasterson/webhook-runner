package webhooked

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/go-playground/webhooks.v5/github"
	"net/http"
	"os"
	"reflect"
	"testing"
)

func TestWebhook(t *testing.T) {
	tests := []struct {
		name     string
		event    github.Event
		typ      interface{}
		filename string
		headers  http.Header
	}{
		{
			name:     "PushEvent",
			event:    github.PushEvent,
			typ:      github.PushPayload{},
			filename: "../testdata/github/push.json",
			headers: http.Header{
				"X-Github-Event":  []string{"push"},
				"X-Hub-Signature": []string{"sha1=0534736f52c2fc5896ef1bd5a043127b20d233ba"},
			},
		},
	}

	for _, tt := range tests {
		tc := tt
		client := &http.Client{}
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			payload, err := os.Open(tc.filename)
			assert.NoError(err)
			defer func() {
				_ = payload.Close()
			}()

			var parseError error
			var results interface{}
			//TODO fix with correct server
			server := newServer(func(w http.ResponseWriter, r *http.Request) {
				results, parseError = hook.Parse(r, tc.event)
			})
			defer server.Close()
			req, err := http.NewRequest(http.MethodPost, server.URL+path, payload)
			assert.NoError(err)
			req.Header = tc.headers
			req.Header.Set("Content-Type", "application/json")

			resp, err := client.Do(req)
			assert.NoError(err)
			assert.Equal(http.StatusOK, resp.StatusCode)
			assert.NoError(parseError)
			assert.Equal(reflect.TypeOf(tc.typ), reflect.TypeOf(results))
		})
	}
}
