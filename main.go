package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"strconv"

	"scanPort-exporter/pkg/scan"
)

type scanPortResult struct {
	Hostname string
	NodeIp   string
	NodePort int
	Status   float64
}

type Exporter struct {
	ScanPort *prometheus.GaugeVec
}

const namespace = "scanport_exporter"

func NewExporter() *Exporter {
	return &Exporter{
		ScanPort: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "scan_port",
			Help:      "端口连通性扫描, 0 关闭，1 开放",
		}, []string{"hostname", "node_ip", "node_port"}),
	}
}

// Describe implements the prometheus.Collector interface
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	e.ScanPort.Describe(ch)
}

// Collect implements the prometheus.Collector interface.
func (e *Exporter) Collect(ch chan<- prometheus.Metric) {

	scanPortResults := CreateScan()
	for _, res := range scanPortResults {
		e.ScanPort.WithLabelValues(res.Hostname, res.NodeIp, strconv.Itoa(res.NodePort)).Set(res.Status)
	}
	e.ScanPort.Collect(ch)
}

func getScanResult(nodeIp string, hostname string, openPorts []int, allPorts []int) []scanPortResult {
	s := []scanPortResult{}
	for _, port := range allPorts {
		status := float64(0)
		for _, openPort := range openPorts {
			if openPort == port {
				status = 1
			}
		}
		sp := scanPortResult{
			Hostname: hostname,
			NodeIp:   nodeIp,
			NodePort: port,
			Status:   status,
		}
		s = append(s, sp)
	}
	return s
}

func CreateScan() []scanPortResult {
	// timeout is ms, process is 会根据端口数量增加, debug 是否记录日志
	nodeIp := "127.0.0.1"
	port := "80,65432,50001-50030"

	scanIP := scan.NewScanIp(100, 10, true)
	openPorts, allPorts := scanIP.GetIpOpenPort(nodeIp, port)
	scanPortResults := getScanResult(nodeIp, "hostname", openPorts, allPorts)
	return scanPortResults
}

func main() {
	prometheus.MustRegister(NewExporter())
	// 暴露
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":9106", nil)
}
