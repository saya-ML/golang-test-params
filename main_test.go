package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getsWithIDInThePath(t *testing.T) {
	tests := []struct {
		name    string
		handler func(http.ResponseWriter, *http.Request)
		method  string
		param   string
		want    string
	}{
		{
			name:    "id mail@mail.com",
			handler: getWithIDInThePath,
			method:  http.MethodGet,
			param:   "mail@mail.com",
			want:    "get: 'mail@mail.com'",
		},
		{
			name:    "id mail@mail.com",
			handler: postWithIDInThePath,
			method:  http.MethodPost,
			param:   "mail@mail.com",
			want:    "post: 'mail@mail.com'",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := http.NewRequest(tt.method, "/", nil)
			r.SetPathValue("id", tt.param)
			assert.Nil(t, err)
			w := httptest.NewRecorder()
			tt.handler(w, r)
			assert.Equal(t, tt.want, w.Body.String())
		})
	}
}
