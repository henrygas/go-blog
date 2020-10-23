package tracer

import (
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"time"
)

func NewJaegerTracer(serviceName string, agentHostPort string) (opentracing.Tracer, io.Closer, error) {
	cfg := &config.Configuration{
		ServiceName: serviceName,
		Sampler: &config.SamplerConfig{	// 取样方式
			Type: "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{ // 上报途径
			LogSpans: true, 						// 是否开启日志
			BufferFlushInterval: 1 * time.Second, 	// 刷新缓冲区频率
			LocalAgentHostPort: agentHostPort,		// 上报的agent地址
		},
	}

	tracer, closer, err := cfg.NewTracer()  // 初始化Tracer对象
	if err != nil {
		return nil, nil, err
	}

	opentracing.SetGlobalTracer(tracer) // 设置全局的Tracer对象
	return tracer, closer, nil
}
