package jaeger

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	configPkg "github.com/serenefiregroup/ffa_server/pkg/config"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

var (
	SpanCTX = "span-ctx"
)

func GetDefaultConfig() *config.Configuration {
	jaegerAddr := configPkg.String("jeager", "http://127.0.0.1:14268/api/traces")
	var cfg = &config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
			// 按实际情况替换你的 ip
			CollectorEndpoint: jaegerAddr,
		},
	}
	return cfg
}

func init() {
	jaegerConfig := GetDefaultConfig()
	InitJaeger("ffa", jaegerConfig)
}

/*
*
初始化
*/
func InitJaeger(service string, cfg *config.Configuration) (opentracing.Tracer, io.Closer) {
	cfg.ServiceName = service
	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("Error: connot init Jaeger: %v\n", err))
	}
	opentracing.SetGlobalTracer(tracer)
	return tracer, closer
}

func StartSpan(tracer opentracing.Tracer, name string) opentracing.Span {
	//设置顶级span
	span := tracer.StartSpan(name)
	return span
}

func WithSpan(ctx context.Context, name string) (opentracing.Span, context.Context) {
	span, ctx := opentracing.StartSpanFromContext(ctx, name)
	return span, ctx
}

func GetCarrier(span opentracing.Span) (opentracing.HTTPHeadersCarrier, error) {
	carrier := opentracing.HTTPHeadersCarrier{}
	err := span.Tracer().Inject(span.Context(), opentracing.HTTPHeaders, carrier)
	if err != nil {
		return nil, err
	}
	return carrier, nil
}

func GetParentSpan(spanName string, traceId string, header http.Header) (opentracing.Span, error) {
	carrier := opentracing.HTTPHeadersCarrier{}
	carrier.Set("uber-trace-id", traceId)

	tracer := opentracing.GlobalTracer()
	wireContext, err := tracer.Extract(
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(header),
	)

	parentSpan := opentracing.StartSpan(
		spanName,
		ext.RPCServerOption(wireContext))
	if err != nil {
		return nil, err
	}
	return parentSpan, err
}
