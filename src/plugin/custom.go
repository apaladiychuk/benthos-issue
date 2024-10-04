package plugin

import (
	"context"
	"github.com/redpanda-data/benthos/v4/public/service"
	"go.opentelemetry.io/otel/attribute"

	"go.opentelemetry.io/otel/trace"
)

type Custom struct {
	mgr *service.Resources
}

func NewCustomPlugin() *Custom {
	return &Custom{}
}

func (r *Custom) Constructor(conf *service.ParsedConfig, mgr *service.Resources) (service.Processor, error) {
	r.mgr = mgr
	return r, nil
}

func (r *Custom) GetConfig() *service.ConfigSpec {
	var cs = service.NewConfigSpec().
		Summary("Custom configuration.")
	return cs
}

func (r *Custom) Process(ctx context.Context, m *service.Message) (service.MessageBatch, error) {
	//spanCtx, span := r.mgr.OtelTracer().Tracer("benthos").Start(m.Context(), "test_span", trace.WithAttributes(attribute.String("test_key", "test_value")))

	span := trace.SpanFromContext(m.Context())
	span.SetAttributes(attribute.String("otel_key_1", "value_1"))
	span.SetAttributes(attribute.String("otel_key_2", "value_2"))
	//m = m.WithContext(trace.ContextWithSpan(spanCtx, span))

	return []*service.Message{m}, nil
}

func (r *Custom) Close(ctx context.Context) error {
	return nil
}
