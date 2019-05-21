package metering

type NamespaceResourceMetering struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	CPU          uint64 `json:"cpu"`
	CPUUsed      uint64 `json:"cpu_used"`
	CPUUsedTime  uint64 `json:"cpu_used_time"`
	Mem          uint64 `json:"mem"`
	MemUsed      uint64 `json:"mem_used"`
	MemUsedTime  uint64 `json:"mem_used_time"`
	Disk         uint64 `json:"disk"`
	DiskUsed     uint64 `json:"disk_used"`
	DiskUsedTime uint64 `json:"disk_used_time"`
	Timestamp    uint64 `json:"timestamp"` //record time 1 hour as an interval
}
