package server_in

import (
	prometheus2 "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/prometheus"
	metric2 "go.opentelemetry.io/otel/sdk/metric"
	"log"
	"net/http"
)

var registry *prometheus2.Registry

func init() {
	initMeter()
}

// 导出初始化
func initMeter() *prometheus2.Registry {
	registry = prometheus2.NewRegistry()
	// 创建Prometheus exporter
	exporter, err := prometheus.New(prometheus.WithRegisterer(registry))
	if err != nil {
		log.Fatalf("创建Prometheus exporter失败: %v", err)
	}

	// 创建MeterProvider
	provider := metric2.NewMeterProvider(
		//metric.WithResource(res),
		metric2.WithReader(exporter),
		metric2.WithView(metric2.NewView(
			metric2.Instrument{Name: "histogram_*"},
			metric2.Stream{Aggregation: metric2.AggregationExplicitBucketHistogram{
				//Boundaries: []float64{0, 5, 10, 25, 50, 75, 100, 250, 500, 1000},
			}},
		)),
	)

	// 设置全局MeterProvider
	otel.SetMeterProvider(provider)

	return registry
}

// 指标输出
func metricsHandle() http.Handler {
	return promhttp.HandlerFor(registry, promhttp.HandlerOpts{MaxRequestsInFlight: 1})
}

/**
// 创建meter实例
example = otel.GetMeterProvider().Meter("my-meter")

// 累加计数器
c64, _ := example.Int64Counter(
	"A1",
	metric.WithDescription("描述文本"),
	metric.WithUnit("单位"),
)
// 记录步进值
c64.Add(r.Context(), 1)

// 直方图
h64, _ := example.Int64Histogram("A2")
// 直方图 记录值
h64.Record(context.Background(), count)

// 瞬时值
g64, _ := example.Int64ObservableGauge("A3",
	metric.WithDescription("描述文本"),
	metric.WithUnit("单位"),
)
// 瞬时值 注册回调
example.RegisterCallback(func(ctx context.Context, o metric.Observer) error {
	o.ObserveInt64(g64, xxx)
	return nil
}, g64)



*/
