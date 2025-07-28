package plugins

import (
	"encoding/json"
	pkgHTTP "github.com/apache/apisix-go-plugin-runner/pkg/http"
	"strings"

	"github.com/apache/apisix-go-plugin-runner/pkg/log"
	"github.com/apache/apisix-go-plugin-runner/pkg/plugin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// 定义全局 Prometheus 指标
var (
	modelLatency = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "model_latency_seconds",
		Help:    "Model inference latency in seconds",
		Buckets: []float64{0.1, 0.5, 1, 2, 5},
	}, []string{"model_name"})

	modelTokens = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "model_tokens_total",
		Help: "Total tokens processed by model",
	}, []string{"model_name", "token_type"})
)

// 解析大模型响应的数据结构
type ModelResponse struct {
	Model   string         `json:"model"`
	Latency float64        `json:"latency"`
	Input   int            `json:"input_tokens"`
	Output  int            `json:"output_tokens"`
	Usage   map[string]int `json:"usage"`
}

// 插件主体结构
type ModelMetrics struct {
	plugin.DefaultPlugin
}

func (p *ModelMetrics) Name() string {
	return "model-metrics"
}

// 响应拦截处理
func (p *ModelMetrics) ResponseFilter(conf interface{}, w pkgHTTP.Response) {
	// 只处理 JSON 响应
	if !strings.Contains(w.Header().Get("Content-Type"), "application/json") {
		return
	}

	// 读取响应体
	originBody, err := w.ReadBody()
	if err != nil {
		log.Errorf("Failed to read response body: %v", err)
		return
	}

	// 解析大模型响应
	var resp ModelResponse
	if err := json.Unmarshal(originBody, &resp); err != nil {
		log.Errorf("Failed to parse model response: %v", err)
		return
	}

	// 记录指标到 Prometheus
	if resp.Model != "" {
		// 记录延迟指标
		if resp.Latency > 0 {
			modelLatency.WithLabelValues(resp.Model).Observe(resp.Latency)
		}

		// 记录 Token 使用情况
		if resp.Input > 0 {
			modelTokens.WithLabelValues(resp.Model, "input").Add(float64(resp.Input))
		}
		if resp.Output > 0 {
			modelTokens.WithLabelValues(resp.Model, "output").Add(float64(resp.Output))
		}

		// 如果有 usage 字段
		if usage, exists := resp.Usage["prompt_tokens"]; exists {
			modelTokens.WithLabelValues(resp.Model, "prompt").Add(float64(usage))
		}
		if usage, exists := resp.Usage["completion_tokens"]; exists {
			modelTokens.WithLabelValues(resp.Model, "completion").Add(float64(usage))
		}
	}
	goto write
write:
	_, err = w.Write(originBody)
	if err != nil {
		log.Errorf("failed to write: %s", err)
	}
}
