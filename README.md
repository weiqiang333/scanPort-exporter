# scanPort-exporter

端口扫描连通性状态的 metrics API, 提供给 prometheus 的 exporter.

端口扫描 metrics 使用 golang 实现高性能高并发.


- metrics API
```
# HELP scanport_exporter_scan_port 端口连通性扫描, 0 关闭，1 开放
# TYPE scanport_exporter_scan_port gauge
scanport_exporter_scan_port{hostname="hostname",node_ip="178.157.63.155",node_port="50001"} 1
scanport_exporter_scan_port{hostname="hostname",node_ip="178.157.63.155",node_port="50002"} 1
scanport_exporter_scan_port{hostname="hostname",node_ip="178.157.63.155",node_port="50003"} 1
scanport_exporter_scan_port{hostname="hostname",node_ip="178.157.63.155",node_port="50004"} 1
scanport_exporter_scan_port{hostname="hostname",node_ip="178.157.63.155",node_port="50005"} 1
scanport_exporter_scan_port{hostname="hostname",node_ip="178.157.63.155",node_port="50006"} 1
scanport_exporter_scan_port{hostname="hostname",node_ip="178.157.63.155",node_port="50007"} 1
scanport_exporter_scan_port{hostname="hostname",node_ip="178.157.63.155",node_port="50008"} 1
scanport_exporter_scan_port{hostname="hostname",node_ip="178.157.63.155",node_port="50009"} 1
scanport_exporter_scan_port{hostname="hostname",node_ip="178.157.63.155",node_port="50010"} 1
scanport_exporter_scan_port{hostname="hostname",node_ip="178.157.63.155",node_port="50011"} 1
scanport_exporter_scan_port{hostname="hostname",node_ip="178.157.63.155",node_port="50012"} 1
scanport_exporter_scan_port{hostname="hostname",node_ip="178.157.63.155",node_port="50013"} 1
scanport_exporter_scan_port{hostname="hostname",node_ip="178.157.63.155",node_port="50014"} 1
scanport_exporter_scan_port{hostname="hostname",node_ip="178.157.63.155",node_port="50015"} 1
scanport_exporter_scan_port{hostname="hostname",node_ip="178.157.63.155",node_port="50016"} 1
scanport_exporter_scan_port{hostname="hostname",node_ip="178.157.63.155",node_port="50017"} 1
scanport_exporter_scan_port{hostname="hostname",node_ip="178.157.63.155",node_port="50018"} 1
scanport_exporter_scan_port{hostname="hostname",node_ip="178.157.63.155",node_port="50019"} 1
scanport_exporter_scan_port{hostname="hostname",node_ip="178.157.63.155",node_port="50020"} 1
scanport_exporter_scan_port{hostname="hostname",node_ip="178.157.63.155",node_port="50021"} 1
scanport_exporter_scan_port{hostname="hostname",node_ip="178.157.63.155",node_port="50022"} 1
scanport_exporter_scan_port{hostname="hostname",node_ip="178.157.63.155",node_port="50023"} 1
scanport_exporter_scan_port{hostname="hostname",node_ip="178.157.63.155",node_port="50024"} 1
scanport_exporter_scan_port{hostname="hostname",node_ip="178.157.63.155",node_port="50025"} 1
scanport_exporter_scan_port{hostname="hostname",node_ip="178.157.63.155",node_port="50026"} 1
scanport_exporter_scan_port{hostname="hostname",node_ip="178.157.63.155",node_port="50027"} 1
scanport_exporter_scan_port{hostname="hostname",node_ip="178.157.63.155",node_port="50028"} 1
scanport_exporter_scan_port{hostname="hostname",node_ip="178.157.63.155",node_port="50029"} 1
scanport_exporter_scan_port{hostname="hostname",node_ip="178.157.63.155",node_port="50030"} 1
scanport_exporter_scan_port{hostname="hostname",node_ip="178.157.63.155",node_port="65432"} 1
scanport_exporter_scan_port{hostname="hostname",node_ip="178.157.63.155",node_port="80"} 0
```