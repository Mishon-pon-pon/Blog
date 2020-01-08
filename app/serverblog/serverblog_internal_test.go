package serverblog

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServerBlog(t *testing.T) {
	s := newServer()
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	s.handleIndex().ServeHTTP(rec, req)
	index := template.Must(template.ParseFiles(
		"github.com/Mishon-pon-pon/Blog/web/index.html",
	))

	assert.Equal(t, rec.Body, index)
}
