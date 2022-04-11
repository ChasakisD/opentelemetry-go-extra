package otelgorm

import (
	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.opentelemetry.io/otel/trace"
)

type Option func(p *otelPlugin)

// WithTracerProvider configures a tracer provider that is used to create a tracer.
func WithTracerProvider(provider trace.TracerProvider) Option {
	return func(p *otelPlugin) {
		p.provider = provider
	}
}

// WithAttributes configures attributes that are used to create a span.
func WithAttributes(attrs ...attribute.KeyValue) Option {
	return func(p *otelPlugin) {
		p.attrs = append(p.attrs, attrs...)
	}
}

// WithDBName configures a db.name attribute.
func WithDBName(name string) Option {
	return func(p *otelPlugin) {
		p.attrs = append(p.attrs, semconv.DBNameKey.String(name))
	}
}

// WithoutQueryVariables configures the db.statement attribute to exclude query variables
func WithoutQueryVariables() Option {
	return func(p *otelPlugin) {
		p.excludeQueryVars = true
	}
}

// WithoutMetrics prevents DBStats metrics from being reported.
func WithoutMetrics() Option {
	return func(p *otelPlugin) {
		p.excludeMetrics = true
	}
}

// WithQueryLimitation prevents DBStats metrics from uploading the full query if it exceeds max characters.
func WithQueryLimitation(queryMaxChars int) Option {
	return func(p *otelPlugin) {
		p.queryLimitation = true
		p.queryMaxChars = queryMaxChars
	}
}
