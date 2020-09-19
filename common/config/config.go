package config

import (
	"log"

	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/api/kv"
	"go.opentelemetry.io/otel/exporters/stdout"
	"go.opentelemetry.io/otel/exporters/trace/jaeger"

	// "go.opentelemetry.io/otel/label"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

const (
	SERVICE_GARAGE_PORT = ":7000"
	SERVICE_USER_PORT   = ":8020"
)

// Init configures an OpenTelemetry exporter and trace provider
func Init() {
	exporter, err := stdout.NewExporter(stdout.WithPrettyPrint())
	if err != nil {
		log.Fatal(err)
	}
	tp, err := sdktrace.NewProvider(
		sdktrace.WithConfig(sdktrace.Config{DefaultSampler: sdktrace.AlwaysSample()}),
		sdktrace.WithSyncer(exporter),
	)
	if err != nil {
		log.Fatal(err)
	}
	global.SetTraceProvider(tp)
}

// func initTracer() func() {
// 	// Create and install Jaeger export pipeline
// 	flush, err := jaeger.InstallNewPipeline(
// 		jaeger.WithCollectorEndpoint("http://localhost:14268/api/traces"),
// 		jaeger.WithProcess(jaeger.Process{
// 			ServiceName: "gotracergrpc",
// 			Tags: []label.KeyValue{
// 				label.String("exporter", "jaeger"),
// 				label.Float64("float", 312.23),
// 			},
// 		}),

// 		jaeger.WithSDK(&sdktrace.Config{DefaultSampler: sdktrace.AlwaysSample()}),
// 	)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return func() {
// 		flush()
// 	}
// }
func InitTraceProvider(service string) func() {
	// Create and install Jaeger export pipeline
	// _, flush, err := jaeger.NewExportPipeline(
	// 	jaeger.WithCollectorEndpoint("http://localhost:14268/api/traces"),
	// 	// jaeger.RegisterAsGlobal(),
	// 	jaeger.WithProcess(jaeger.Process{
	// 		ServiceName: service,
	// 		Tags: []kv.KeyValue{
	// 			kv.Key("exporter").String("jaeger"),
	// 		},
	// 	}),
	// 	// jaeger.RegisterAsGlobal(),
	// 	jaeger.WithSDK(&sdktrace.Config{DefaultSampler: sdktrace.AlwaysSample()}),
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// return func() {
	// 	flush()
	// }
	flush, err := jaeger.InstallNewPipeline(
		jaeger.WithCollectorEndpoint("http://localhost:14268/api/traces"),
		jaeger.WithProcess(jaeger.Process{
			ServiceName: service,
			Tags: []kv.KeyValue{
				kv.Key("exporter").String("jaeger"),
			},
		}),

		jaeger.WithSDK(&sdktrace.Config{DefaultSampler: sdktrace.AlwaysSample()}),
	)
	if err != nil {
		log.Fatal(err)
	}

	return func() {
		flush()
	}
}
