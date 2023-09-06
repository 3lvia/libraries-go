package hashivault

import (
	"go.opentelemetry.io/otel/trace"
	"log"
)

const defaultTracerName = "go.opentelemetry.io/otel"

var tracerName string

func traceError(span trace.Span, err error, l *log.Logger) {
	if err != nil {
		l.Printf("error: %s", err.Error())
		span.RecordError(err)
	}
}
