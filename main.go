package main

import (
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"runtime"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"scanPort-exporter/pkg/scan"
)

type Exporter struct {
	ScanPort *prometheus.GaugeVec
}

const (
	namespace = "scanport_exporter"
	version   = "v1.2"
)

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
	e.ScanPort.Reset()
	s := scan.NewScanPortResult()
	scanPortResults := s.CreateScan()
	for _, res := range scanPortResults.Results {
		e.ScanPort.WithLabelValues(res.Hostname, res.NodeIp, strconv.Itoa(res.NodePort)).Set(res.Status)
	}
	e.ScanPort.Collect(ch)
}

func init() {
	pflag.Bool("version", false, "print the version")
	pflag.String("scan_source", "configfile", "scan_source configfile/prometheus")
	pflag.String("address", ":9106", "Listen address")
	pflag.String("config_file", "config/config.yaml", "config file")
	pflag.Parse()
}

func main() {
	viper.BindPFlags(pflag.CommandLine)
	if viper.GetBool("version") {
		fmt.Println("scanPort-exporter version: ", version)
		return
	}
	viper.SetConfigType("yaml")
	viper.SetConfigFile(viper.GetString("config_file"))
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	prometheus.MustRegister(NewExporter())

	address := viper.GetString("address")
	if err = open(fmt.Sprintf("http://127.0.0.1:9106/metrics")); err != nil {
		fmt.Println(err.Error())
	}
	// 暴露
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/reload", reloadConfig)
	http.ListenAndServe(address, nil)
}

func reloadConfig(w http.ResponseWriter, _ *http.Request) {
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		fmt.Println(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	fmt.Println(fmt.Sprintf("reload config file: %s", viper.ConfigFileUsed()))
	io.WriteString(w, fmt.Sprintf("reload config file: %s", viper.ConfigFileUsed()))
}

func open(uri string) error {
	var commands = map[string]string{
		"windows": "start",
		"darwin":  "open",
		"linux":   "xdg-open",
	}
	run, ok := commands[runtime.GOOS]
	if !ok {
		return fmt.Errorf("%s platform ？？？", runtime.GOOS)
	}

	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "start ", uri)
		//cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	} else {
		cmd = exec.Command(run, uri)
	}
	return cmd.Start()
}
