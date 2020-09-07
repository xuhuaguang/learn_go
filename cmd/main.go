package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/expfmt"
	"net/http"
	"strings"
	"sync/atomic"
)

/*var (
	httpRequestTotal = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "http_request_in_flight",
			Help: "Current number of http requests in flight",
		},
	)

	httpRequestInFlight = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "http_request_in_flight_Gauge",
			Help: "Current number of http requests in flight Gauge",
		},
	)

	httpRequestDurationSeconds = promauto.NewHistogram(
		prometheus.HistogramOpts{
			Name: "http_request_duration_seconds",
			Help: "Histogram of lantencies for HTTP requests",
		},
	)

	http_request_total = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_request_total",
			Help: "The total number of processed http requests",
		},
		[]string{"path"},
	)
)*/

type statusCollect struct {
	reqDesc       *prometheus.CounterVec
	respSizeDesc  *prometheus.Desc
	respSizeValue int64
}

func (s *statusCollect) ReqAdd(code, method string) {
	s.reqDesc.WithLabelValues(code, method).Inc()
}

func (s *statusCollect) ReqSizeAdd(size int64) {
	atomic.AddInt64(&s.respSizeValue, size)
}

func (s *statusCollect) Describe(ch chan<- *prometheus.Desc) {
	ch <- s.respSizeDesc
	s.reqDesc.Describe(ch)
}

func (s *statusCollect) Collect(ch chan<- prometheus.Metric) {
	ch <- prometheus.MustNewConstMetric(s.respSizeDesc, prometheus.CounterValue, float64(s.respSizeValue))
	s.reqDesc.Collect(ch)
}

func NewStatusCollect() *statusCollect {
	opts := prometheus.CounterOpts{Namespace: "Test", Subsystem: "http", Name: "request_count", Help: "request count"}
	return &statusCollect{
		reqDesc:      prometheus.NewCounterVec(opts, []string{"code", "method"}),
		respSizeDesc: prometheus.NewDesc("http_resp_size_count", "http resp size count", nil, nil),
	}
}

const (
	contentTypeHeader   = "Content-Type"
	contentLengthHeader = "Content-Length"
)

func main() {
	status := NewStatusCollect()
	regist := prometheus.NewRegistry()
	regist.MustRegister(status)

	http.HandleFunc("/metric", func(w http.ResponseWriter, r *http.Request) {
		status.ReqAdd("200", strings.ToLower(r.Method))

		entry, err := regist.Gather()
		if err != nil {
			http.Error(w, "An error has occurred during metrics collection:\n\n"+err.Error(), http.StatusInternalServerError)
			return
		}

		buf := bytes.NewBuffer(nil)
		contentType := expfmt.Negotiate(r.Header)
		enc := expfmt.NewEncoder(buf, contentType)

		for _, met := range entry {
			if err := enc.Encode(met); err != nil {
				http.Error(w, "An error has occurred during metrics encoding:\n\n"+err.Error(), http.StatusInternalServerError)
				return
			}
		}

		if buf.Len() == 0 {
			http.Error(w, "No metrics encoded, last error:\n\n"+err.Error(), http.StatusInternalServerError)
			return
		}
		status.ReqSizeAdd(int64(buf.Len()))
		header := w.Header()
		header.Set(contentTypeHeader, string(contentType))
		header.Set(contentLengthHeader, fmt.Sprint(buf.Len()))
		_, _ = w.Write(buf.Bytes())
	})

	_ = http.ListenAndServe(":8000", nil)
}

/*func _() {
	http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
		now := time.Now()

		httpRequestInFlight.Inc()
		defer httpRequestInFlight.Dec()
		httpRequestTotal.Inc()

		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		httpRequestDurationSeconds.Observe(time.Since(now).Seconds())

		http_request_total.WithLabelValues("root").Inc()
	})

	http.HandleFunc("/foo", func(http.ResponseWriter, *http.Request) {
		http_request_total.WithLabelValues("foo").Inc()
	})

	http.Handle("/metrics", promhttp.Handler())
	_ = http.ListenAndServe(":8000", nil)
}*/

func _() {
	router := gin.New()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello word")
	})
	_ = router.Run(":8000")
}
