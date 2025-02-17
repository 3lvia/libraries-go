package problemdetails

import "github.com/gin-gonic/gin/render"

const (
	UnsetType = "about:blank"
)

// D renders a ProblemDetails response.
func D(details ProblemDetails) JSON {
	if details.Type == "" {
		details.Type = UnsetType
	}
	return JSON{render.JSON{Data: details}}
}

// U renders a ProblemDetails response with the given status code, title, and detail.
func U(code int, title, detail string) JSON {
	return D(ProblemDetails{
		Status: code,
		Title:  title,
		Detail: detail,
	})
}

// P renders a ProblemDetails response with the given status code, type, title, and detail.
func P(code int, problemType, title, detail string) JSON {
	return D(ProblemDetails{
		Status: code,
		Type:   problemType,
		Title:  title,
		Detail: detail,
	})
}
