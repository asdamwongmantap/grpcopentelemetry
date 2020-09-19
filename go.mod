module tesprotogrpc

go 1.15

require (
	github.com/DataDog/sketches-go v0.0.1 // indirect
	github.com/alecthomas/units v0.0.0-20190924025748-f65c72e2690d // indirect
	github.com/benbjohnson/clock v1.0.3 // indirect
	github.com/envoyproxy/go-control-plane v0.9.4 // indirect
	github.com/go-kit/kit v0.10.0 // indirect
	github.com/gogo/protobuf v1.3.1
	github.com/golang/mock v1.4.3 // indirect
	github.com/golang/protobuf v1.4.2
	github.com/google/gofuzz v1.1.0 // indirect
	github.com/google/martian/v3 v3.0.0 // indirect
	github.com/google/pprof v0.0.0-20200708004538-1a94d8640e99 // indirect
	github.com/jpillora/backoff v1.0.0 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/julienschmidt/httprouter v1.3.0 // indirect
	github.com/labstack/echo-contrib v0.9.0
	github.com/labstack/echo/v4 v4.1.17
	github.com/mwitkow/go-conntrack v0.0.0-20190716064945-2f068394615f // indirect
	github.com/opentracing/opentracing-go v1.2.0
	github.com/pkg/errors v0.9.1 // indirect
	github.com/prometheus/client_golang v1.5.1
	github.com/prometheus/common v0.10.0 // indirect
	github.com/prometheus/procfs v0.1.3 // indirect
	github.com/sirupsen/logrus v1.6.0 // indirect
	github.com/uber/jaeger-client-go v2.25.0+incompatible
	github.com/uber/jaeger-lib v2.2.0+incompatible
	go.opencensus.io v0.22.4 // indirect
	go.opentelemetry.io/contrib v0.10.1 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc v0.10.1
	go.opentelemetry.io/otel v0.10.0
	// go.opentelemetry.io/otel v0.11.0
	go.opentelemetry.io/otel/exporter/trace/jaeger v1.0.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout v0.10.0
	go.opentelemetry.io/otel/exporters/trace/jaeger v0.10.0
	go.opentelemetry.io/otel/sdk v0.10.0
	go.uber.org/atomic v1.6.0 // indirect
	golang.org/x/exp v0.0.0-20200224162631-6cc2880d07d6 // indirect
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
	golang.org/x/net v0.0.0-20200904194848-62affa334b73 // indirect
	golang.org/x/sys v0.0.0-20200916084744-dbad9cb7cb7a // indirect
	golang.org/x/tools v0.0.0-20200804011535-6c149bb5ef0d // indirect
	google.golang.org/api v0.30.0 // indirect
	google.golang.org/appengine v1.6.6 // indirect
	google.golang.org/grpc v1.32.0
	google.golang.org/protobuf v1.25.0
	gopkg.in/yaml.v2 v2.3.0 // indirect
	honnef.co/go/tools v0.0.1-2020.1.4 // indirect

)

replace (
	go.opentelemetry.io/otel v0.11.0 => go.opentelemetry.io/otel v0.10.0
	go.opentelemetry.io/otel/api/kv v0.11.0 => go.opentelemetry.io/otel/api/kv v0.10.0
	go.opentelemetry.io/otel/plugin/httptrace v0.11.0 => go.opentelemetry.io/otel/plugin/httptrace v0.6.0
	// go.opentelemetry.io/otel/exporter/trace/jaeger v1.0.0 => go.opentelemetry.io/otel/exporter/trace/jaeger v0.2.1
)
