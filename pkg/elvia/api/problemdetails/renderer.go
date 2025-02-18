package problemdetails

import (
	"net/http"

	"github.com/gin-gonic/gin/render"
)

const (
	contentType = "application/problem+json; charset=utf-8"
)

// JSON is a helper to render a ProblemDetails response.
type JSON struct {
	render.JSON
}

// WriteContentType writes the Content-Type header for a ProblemDetails response.
func (r JSON) WriteContentType(w http.ResponseWriter) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = []string{contentType}
	}
}

// Render renders the content.
func (r JSON) Render(w http.ResponseWriter) error {
	r.WriteContentType(w)
	return r.JSON.Render(w)
}
