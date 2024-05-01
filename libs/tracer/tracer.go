package tracer

import (
	"context"
	"log"
	"os"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	oteltrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"

	"go.opentelemetry.io/otel/trace"
)

var (
	tracer        trace.Tracer
	OLTP_ENDPOINT string
)

func init() {
	_OTLP_ENDPOINT := os.Getenv("OTLP_ENDPOINT")
	if _OTLP_ENDPOINT == "" {
		log.Fatal("OTLP_ENDPOINT is not set")
	}
	OLTP_ENDPOINT = _OTLP_ENDPOINT
}

func NewConsoleExporter() (oteltrace.SpanExporter, error) {
	return stdouttrace.New(stdouttrace.WithPrettyPrint())
}

// OTLP Exporter
func NewOTLPExporter(ctx context.Context) (oteltrace.SpanExporter, error) {
	tracer, err := otlptracehttp.New(ctx,
		otlptracehttp.WithInsecure(),
		otlptracehttp.WithEndpoint(OLTP_ENDPOINT),
		otlptracehttp.WithCompression(otlptracehttp.GzipCompression),
		// otlptracehttp.WithEndpointURL(OLTP_ENDPOINT),
	)
	return tracer, err
}

func NewTraceProvider(exp oteltrace.SpanExporter) *oteltrace.TracerProvider {

	r, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			"https://opentelemetry.io/schemas/1.24.0",
			semconv.ServiceName("app-x"),
		),
	)

	// resource.Default(),
	// resource.NewWithAttributes(
	// 	semconv.SchemaURL,
	// 	semconv.ServiceName("backend-service"),
	// 	semconv.ServiceVersion("1.0.0"),
	// 	semconv.ServiceNamespace("niko-labs"),
	// 	semconv.ServiceInstanceID("backend-service-instance"),
	// ),

	if err != nil {
		panic(err)
	}

	return oteltrace.NewTracerProvider(
		oteltrace.WithSampler(oteltrace.AlwaysSample()),
		oteltrace.WithBatcher(exp),
		oteltrace.WithResource(r),
	)
}

func SaveTracer(t trace.Tracer) {
	tracer = t
}

func GetTracer() trace.Tracer {
	return tracer
}
