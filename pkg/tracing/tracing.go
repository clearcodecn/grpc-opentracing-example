package tracing

import (
	"fmt"
	opentracing "github.com/opentracing/opentracing-go"
	jaeger "github.com/uber/jaeger-client-go"
	config "github.com/uber/jaeger-client-go/config"
	"io"
)

func init() {
	//os.Setenv("JAEGER_AGENT_HOST", "192.168.1.101")
	//os.Setenv("JAEGER_AGENT_PORT", "6831")
}

// Init returns an instance of Jaeger Tracer that samples 100% of traces and logs all spans to stdout.
func Init(service string) (opentracing.Tracer, io.Closer) {
	cfg, err := config.FromEnv()

	cfg.ServiceName = service
	cfg.Sampler.Type = "const"
	cfg.Sampler.Param = 1
	cfg.Reporter.LogSpans = true

	//cfg = &config.Configuration{
	//	ServiceName: service,
	//	Sampler: &config.SamplerConfig{
	//		Type:  "const",
	//		Param: 1,
	//	},
	//	Reporter: &config.ReporterConfig{
	//		LogSpans: true,
	//	},
	//}
	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	return tracer, closer
}
