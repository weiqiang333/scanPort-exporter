package scan

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/spf13/viper"
)

type ScanResult struct {
	ScanSource string
	Results    []Result
}

type Result struct {
	Hostname string
	NodeIp   string
	NodePort int
	Status   float64
}

func NewScanPortResult() *ScanResult {
	return &ScanResult{}
}

func (s *ScanResult) getScanResult(nodeIp string, hostname string, openPorts []int, allPorts []int) {
	for _, port := range allPorts {
		status := float64(0)
		for _, openPort := range openPorts {
			if openPort == port {
				status = 1
			}
		}
		sp := Result{
			Hostname: hostname,
			NodeIp:   nodeIp,
			NodePort: port,
			Status:   status,
		}
		s.Results = append(s.Results, sp)
	}
	return
}

func (s *ScanResult) CreateScan() ScanResult {
	// timeout is ms, process is 会根据端口数量增加, debug 是否记录日志暂未使用
	timeout := viper.GetInt("timeout_ms")
	process := viper.GetInt("process")
	scanIP := NewScanIp(timeout, process, true)
	scanSource := viper.GetString("scan_source")
	if scanSource == "configfile" {
		s.getConfigFile(scanIP)
	}
	if scanSource == "prometheus" {
		s.getConfigPrometheus(scanIP)
	}
	return *s
}

func (s *ScanResult) getConfigFile(scanIP *ScanIp) {
	jobNum := len(viper.Get("scan_address_port").([]interface{}))
	for i := 0; i < jobNum; i++ {
		nodeIps := viper.GetString(fmt.Sprintf("scan_address_port.%v.node_ip", i))
		ports := viper.GetString(fmt.Sprintf("scan_address_port.%v.ports", i))
		hostname := viper.GetString(fmt.Sprintf("scan_address_port.%v.hostname", i))
		ips, err := scanIP.GetAllIp(nodeIps)
		if err != nil {
			fmt.Println(fmt.Sprintf("Failed scanIP GetAllIp %s is err: %s", nodeIps, err.Error()))
			continue
		}

		for _, ip := range ips {
			openPorts, allPorts := scanIP.GetIpOpenPort(ip, ports)
			s.getScanResult(ip, hostname, openPorts, allPorts)
		}
	}
}

type prometheusApiRes struct {
	Status    string              `json:"status"`
	Data      prmetheusApiResData `json:"data"`
	ErrorType string              `json:"errorType"`
	Error     string              `json:"error"`
	Warnings  string              `json:"warnings"`
}

type prmetheusApiResData struct {
	ResultType string         `json:"resultType"`
	Result     []resultMetric `json:"result"`
}

type resultMetric struct {
	Metric map[string]string `json:"metric"`
	Value  []interface{}     `json:"value"`
}

func (s *ScanResult) getConfigPrometheus(scanIP *ScanIp) {
	prAPI := viper.GetString("prometheus.query.api")
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get(prAPI)
	if err != nil {
		fmt.Println("Failed getConfigPrometheus http client get error: ", err.Error())
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	node := prometheusApiRes{}
	if err = json.Unmarshal(body, &node); err != nil {
		fmt.Println("Failed getConfigPrometheus json Unmarshal error:", err.Error())
		return
	}
	if node.Status != "success" {
		fmt.Println("Failed getConfigPrometheus data node.Status not is success")
		return
	}
	for _, node := range node.Data.Result {
		ip := node.Metric[viper.GetString("prometheus.query.labels.node_ip")]
		hostname := node.Metric[viper.GetString("prometheus.query.labels.hostname")]
		ports := viper.GetString("prometheus.query.labels.ports")
		openPorts, allPorts := scanIP.GetIpOpenPort(ip, ports)
		s.getScanResult(ip, hostname, openPorts, allPorts)
	}
}
