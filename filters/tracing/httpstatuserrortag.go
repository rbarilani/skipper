package tracing

import (
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/zalando/skipper/filters"
)

const (
	httpStatusErrorTagName = "tracingHTTPStatusErrorTag"
)

type httpStatusErrorTagSpec struct {
}

type httpStatusErrorTagFilter struct{}

// NewTracingHTTPStatusErrorTag creates a new filter spec
func NewTracingHTTPStatusErrorTag() filters.Spec {
	return httpStatusErrorTagSpec{}
}

func (s httpStatusErrorTagSpec) Name() string {
	return httpStatusErrorTagName
}

func (s httpStatusErrorTagSpec) CreateFilter(args []interface{}) (filters.Filter, error) {
	if len(args) > 0 {
		return nil, filters.ErrInvalidFilterParameters
	}

	return httpStatusErrorTagFilter{}, nil
}

func (f httpStatusErrorTagFilter) Request(ctx filters.FilterContext) {
}

func (f httpStatusErrorTagFilter) Response(ctx filters.FilterContext) {
	req := ctx.Request()
	res := ctx.Response()
	span := opentracing.SpanFromContext(req.Context())
	if span == nil {
		return
	}

	if res.StatusCode >= 500 {
		span.SetTag("error", true)
		span.SetTag("http.status_error", true)
	}
}
