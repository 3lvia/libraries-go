package problemdetails

// ProblemDetails is a struct that represents a problem details object
// https://datatracker.ietf.org/doc/html/rfc9457
type ProblemDetails struct {
	// Type is a member containing a URI reference that identifies the problem type.
	// Consumers MUST use the "type" URI as the problem type's primary identifier.
	// It is RECOMMENDED that the URI reference be an absolute URI and that it provide human-readable documentation.
	// When this member is not present, its value is assumed to be "about:blank".
	Type string `json:"type"`

	// Status is a member containing an integer HTTP status code for this occurrence of the problem.
	// The "status" member, if present, is only advisory; it conveys the HTTP status code used for the convenience of the consumer.
	Status int `json:"status"`

	// Title is a member containing a short, human-readable summary of the problem type.
	// It SHOULD NOT change from occurrence to occurrence of the problem, except for purposes of localization.
	Title string `json:"title"`

	// Detail is a member containing a human-readable explanation specific to this occurrence of the problem.
	// Like "title", this member's value can be localized.
	// It ought to focus on helping the client correct the problem, rather than giving debugging information.
	Detail string `json:"detail,omitempty"`

	// Instance is a member containing a URI reference that identifies the specific occurrence of the problem.
	// When the "instance" URI is dereferenceable, it can provide additional information about the problem for consumers.
	// When the "instance" URI is not dereferenceable, it serves as a unique identifier for the problem occurrence that may be
	// of significance to the server but is opaque to the consumer.
	Instance string `json:"instance,omitempty"`
}
